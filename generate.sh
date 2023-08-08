#!/bin/bash

rm -rf client

# Download the latest OpenAPI specification
# This should be another step
#curl -s http://localhost:4242/docs/openapi.json | jq > openapi.json
# get enterprise version
curl -s https://app.unleash-hosted.com/hosted/docs/openapi.json | jq > openapi.json

openapi-generator-cli generate \
    --git-user-id Unleash \
    --git-repo-id unleash-server-api-go \
    --openapi-normalizer REMOVE_ANYOF_ONEOF_AND_KEEP_PROPERTIES_ONLY=true \
    --openapi-normalizer SIMPLIFY_ONEOF_ANYOF=true \
    --openapi-normalizer SIMPLIFY_ANYOF_STRING_AND_ENUM_STRING=true \
    --openapi-normalizer REFACTOR_ALLOF_WITH_PROPERTIES_ONLY=true \
    --additional-properties packageName=client \
    --legacy-discriminator-behavior \
    --import-mappings time.Time=time \
    -i openapi.json \
    -o client \
    -g go

rm client/go.mod client/go.sum

go mod tidy

# Remove the generated tests. They require some manual work to get working such as creating test parameters
#find client -name *_test.go -exec sed -i '/remove to run test/d' {} \;

# Format the generated code
gofmt -w client