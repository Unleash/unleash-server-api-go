# Client documentation

[Client docs](./client/README.md)

# OpenAPI generator used

Instalation instructions: https://openapi-generator.tech/docs/installation#bash-launcher-script

Version used:

```
openapi-generator-cli 6.6.0
  commit : 7f8b853
  built  : -999999999-01-01T00:00:00+18:00
  source : https://github.com/openapitools/openapi-generator
  docs   : https://openapi-generator.tech/
```

# Generating the client:

(optional) Update the `openapi.json` spec:

1. docker compose up
2. curl -s https://localhost:4242/docs/openapi.json | jq > openapi.json

## Generate client from opeanpi.json

```bash
./generate.sh
```

After generating the code you can use `main.go` to test it works (before running the command below start Unleash server with `docker compose up`)

```bash
go run main.go -unleash-api http://localhost:4242 -authorization *:*.unleash-insecure-admin-api-token
```
