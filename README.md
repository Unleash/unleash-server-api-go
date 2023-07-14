# Generating swagger folder:

```bash
openapi-generator-cli generate --git-user-id Unleash --git-repo-id unleash-server-api-go -i openapi.json -o client -g go -c openapi-config.yaml

# these are structures that are not defined and throw a compilation issue. With this we can still move forward and figure it out later
# find swagger -name "*.go" | xargs sed -i 's/Array/[]interface{}/g'
# find swagger -name "*.go" | xargs sed -i -r 's/(Object|ModelMap|ExportQuerySchema)/interface{}/g'

# lint
gofmt -w client
```

After generating the code you can use `main.go` to test it works (remember updating unleash servere host and API token)

```bash
go run main.go -unleash-api http://localhost:4242 -authorization user:a3c8e3e76e7361c0bc79070accf70409d03fbaef4c2b6b90bff466e8
```
