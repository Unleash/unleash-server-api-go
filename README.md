# Client documentation

[Client docs](./client/README.md)

# Requirements

- Go > 1.18

# Requirements

1. OpenAPI generator: https://openapi-generator.tech/docs/installation#bash-launcher-script

Version used:

```
openapi-generator-cli 6.6.0
  commit : 7f8b853
  built  : -999999999-01-01T00:00:00+18:00
  source : https://github.com/openapitools/openapi-generator
  docs   : https://openapi-generator.tech/
```

2. openapi format: https://github.com/thim81/openapi-format

# Generating the client:

(optional) Update the `openapi.json` spec:

1. docker compose up
2. curl -s http://localhost:4242/docs/openapi.json | jq > openapi.json

## Generate client from opeanpi.json

```bash
./generate.sh
```

After generating the code you can use `main.go` to test it works (before running the command below start Unleash server with `docker compose up`)

```bash
go run main.go -unleash-api http://localhost:4242 -authorization *:*.unleash-insecure-admin-api-token
```

# Testing

Testing relies on a running environment. You can use docker-compose.yml file to spin up an environment with a clean state and predictable admin API token.

**Note**: some tests rely on an enterprise version of Unleash. To run those tests locally you need to set the environment variable `UNLEASH_ENTERPRISE=true`. To run docker with an enterprise image, login to GH docker registry following [preparation steps](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry#authenticating-with-a-personal-access-token-classic) and then set the environment variable `UNLEASH_DOCKER_IMAGE` pointing to the enterprise docker image.

A basic test is automatically generated by the tool (find them under `client/test`), but they don't have a default configuration and doesn't have payloads required for some of the endpoints. This is the reason why the tools marks all of them as Skip. Still they're a good starting point.

We recommend copying specific tests from `client/test` into `test` and adapting them accordingly. We've done that already for api_users_test and some of its methods. Also, recommended to use `apiClient := testClient()` to instantiate the client.

## Running the tests

1. (optional but recommended) `docker compose rm --force`
1. `docker compose up`
1. `go test ./test/... -count=1`
