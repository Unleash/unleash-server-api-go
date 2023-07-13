# Generating swagger folder:

```bash
java -jar swagger-codegen-cli-3.0.20.jar generate -i https://us.app.unleash-hosted.com/ushosted/docs/openapi.json -l go -o swagger

# these are structures that are not defined and throw a compilation issue. With this we can still move forward and figure it out later
find swagger -name "*.go" | xargs sed -i 's/Array/[]interface{}/g'
find swagger -name "*.go" | xargs sed -i -r 's/(Object|ModelMap|ExportQuerySchema)/interface{}/g'

# lint
gofmt -w swagger
```

After generating the code you can use `main.go` to test it works (remember updating unleash servere host and API token)

```bash
go run main.go -unleash-api http://localhost:4242 -authorization user:a3c8e3e76e7361c0bc79070accf70409d03fbaef4c2b6b90bff466e8
```
