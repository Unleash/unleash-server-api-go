#!/bin/bash

rm -rf client
set -e

# Download the latest OpenAPI specification
# This should be another step
# get enterprise version
curl -s https://us.app.unleash-hosted.com/ushosted/docs/openapi.json | jq > openapi.json

# Keep only the operations we support
openapi-format openapi.json --filterFile operations.yaml --json -o modified-openapi.json

openapi-generator-cli generate \
    --git-user-id Unleash \
    --git-repo-id unleash-server-api-go \
    --openapi-normalizer REFACTOR_ALLOF_WITH_PROPERTIES_ONLY=true \
    --openapi-normalizer RESOLVE_INLINE_ENUMS=true \
    --openapi-normalizer REF_AS_PARENT_IN_ALLOF=true \
    --additional-properties packageName=client \
    -i modified-openapi.json \
    -o client \
    -g go

rm client/go.mod client/go.sum
# After removing client/go.mod and client/go.sum, the generated code will not compile when importing the package
# so we make some modifications to make it compile:
sed -i 's/openapiclient "github.com\/Unleash\/unleash-server-api-go"/"github.com\/Unleash\/unleash-server-api-go\/client"/g' client/test/*
sed -i 's/openapiclient\./client./g' client/test/*
go mod tidy

# Remove the generated tests. They require some manual work to get working such as creating test parameters
#find client -name *_test.go -exec sed -i '/remove to run test/d' {} \;

# Format the generated code
gofmt -w client