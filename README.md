# Generating swagger folder:

```bash
java -jar swagger-codegen-cli-3.0.20.jar generate -i openapi.json -l go -o swagger

find swagger -name "*.go" | xargs sed -i 's/Array/[]interface{}/g'
find swagger -name "*.go" | xargs sed -i 's/ModelMap/interface{}/g'
find swagger -name "*.go" | xargs sed -i 's/Object/interface{}/g'

# then run
gofmt -w swagger
```

After generating the code you can use `main.go` to test it works (remember updating unleash servere host and API token)

```bash
go run main.go --debug-http
```
