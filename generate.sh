#!/bin/bash

set -e
rm -rf client

# Download the OpenAPI specification.
# Prefer local Unleash (useful for pinned Docker images), fallback to hosted.
LOCAL_OPENAPI_URL=${LOCAL_OPENAPI_URL:-http://localhost:4242/docs/openapi.json}
HOSTED_OPENAPI_URL=${HOSTED_OPENAPI_URL:-https://us.app.unleash-hosted.com/ushosted/docs/openapi.json}

if curl -fsS --max-time 5 "$LOCAL_OPENAPI_URL" | jq >openapi.json; then
  echo "Using local OpenAPI spec from: $LOCAL_OPENAPI_URL"
else
  echo "Local OpenAPI spec unavailable, falling back to: $HOSTED_OPENAPI_URL"
  curl -fsS "$HOSTED_OPENAPI_URL" | jq >openapi.json
fi

# Normalize source-specific metadata to keep diffs stable between local and hosted specs.
OPENAPI_SERVER_URL=${OPENAPI_SERVER_URL:-https://unleash.example.com}
jq --arg serverUrl "$OPENAPI_SERVER_URL" '.servers = [{url: $serverUrl}]' openapi.json >tmp.json && mv tmp.json openapi.json

# Keep only the operations we support
openapi-format openapi.json --filterFile operations.yaml --json -o modified-openapi.json

# Parse out our filter ops file, iterate over it, and if there's a match, only retain the listed properties
filters=$(jq -r 'keys[]' filter-ops.json)
for filterName in $filters; do
  jq --slurpfile filters filter-ops.json \
    --arg filterName "$filterName" '
    .components.schemas[$filterName].properties |=
        with_entries(select(.key as $k | $filters[0][$filterName] | index($k) != null))
    ' modified-openapi.json >tmp.json && mv tmp.json modified-openapi.json
done

# There's no reason to have additionalProperties equal to false, it can only cause breakages, so let's just... turn that on
jq 'walk(if type == "object" and .additionalProperties == false then .additionalProperties = true else . end)' modified-openapi.json >tmp.json && mv tmp.json modified-openapi.json

# Remove Enterprise SVG images from descriptions while keeping the Enterprise feature text
jq 'walk(if type == "string" then gsub("!\\[Unleash Enterprise\\]\\([^)]*Enterprise\\.svg\\) ?"; "") else . end)' modified-openapi.json >tmp.json && mv tmp.json modified-openapi.json

openapi-generator-cli generate \
  --git-user-id Unleash \
  --git-repo-id unleash-server-api-go \
  --openapi-normalizer REFACTOR_ALLOF_WITH_PROPERTIES_ONLY=true \
  --openapi-normalizer RESOLVE_INLINE_ENUMS=true \
  --openapi-normalizer REF_AS_PARENT_IN_ALLOF=true \
  --additional-properties packageName=client,disallowAdditionalPropertiesIfNotPresent=false \
  --additional-properties packageName=client,withGoMod=false \
  -i modified-openapi.json \
  -o client \
  -g go

# Remove the generated tests. They require some manual work to get working and they're very low value
rm -rf client/test

# Format the generated code
go mod tidy
gofmt -w client
