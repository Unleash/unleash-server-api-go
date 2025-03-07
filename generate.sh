#!/bin/bash

rm -rf client
set -e

# Download the latest OpenAPI specification
# This should be another step
# get enterprise version
# curl -s https://us.app.unleash-hosted.com/ushosted/docs/openapi.json | jq > openapi.json

# Keep only the operations we support
openapi-format openapi.json --filterFile operations.yaml --json -o modified-openapi.json

# Parse out our filter ops file, iterate over it, and if there's a match, only retain the listed properties
filters=$(jq -r 'keys[]' filter-ops.json)
for filterName in $filters; do
    jq --argfile filters filter-ops.json \
    --arg filterName "$filterName" '
    .components.schemas[$filterName].properties |=
        with_entries(select(.key as $k | $filters[$filterName] | index($k) != null))
    ' modified-openapi.json > tmp.json && mv tmp.json modified-openapi.json
done

# There's no reason to have additionalProperties equal to false, it can only cause breakages, so let's just... turn that on
jq 'walk(if type == "object" and .additionalProperties == false then .additionalProperties = true else . end)' modified-openapi.json > tmp.json && mv tmp.json modified-openapi.json

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