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

```bash
./generate.sh
```

After generating the code you can use `main.go` to test it works (remember updating unleash servere host and API token)

```bash
go run main.go -unleash-api http://localhost:4242 -authorization user:a3c8e3e76e7361c0bc79070accf70409d03fbaef4c2b6b90bff466e8
```
