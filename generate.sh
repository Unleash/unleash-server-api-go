#!/bin/bash

rm -rf client

# Download the latest OpenAPI specification
# This should be another step
#curl -s http://localhost:4242/docs/openapi.json | jq > openapi.json
# get enterprise version
curl -s https://app.unleash-hosted.com/hosted/docs/openapi.json | jq > openapi.json

# Delete some schemas and paths using those schemas
cat openapi.json | \
jq 'del(.components.schemas | (
    .changeRequestsSchema,
    .changeRequestSchema,
    .changeRequestFeatureSchema,
    .changeRequestEventSchema,
    .changeRequestOneOrManyCreateSchema
))' | \
jq 'del(.paths | (
."/api/admin/projects/{projectId}/change-requests",
."/api/admin/projects/{projectId}/change-requests/config",
."/api/admin/projects/{projectId}/change-requests/open",
."/api/admin/projects/{projectId}/change-requests/pending",
."/api/admin/projects/{projectId}/change-requests/pending/{featureName}",
."/api/admin/projects/{projectId}/change-requests/{changeRequestId}/changes/{changeId}",
."/api/admin/projects/{projectId}/change-requests/{id}",
."/api/admin/projects/{projectId}/change-requests/{id}/comments",
."/api/admin/projects/{projectId}/change-requests/{id}/state",
."/api/admin/projects/{projectId}/change-requests/{id}/title",
."/api/admin/projects/{projectId}/environments/{environment}/change-requests",
."/api/admin/projects/{projectId}/environments/{environment}/change-requests/config"
))' > modified-openapi.json

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