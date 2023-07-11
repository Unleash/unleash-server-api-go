To generate the source code run the following command

```bash
bin/openapi-generator-cli generate -i https://us.app.unleash-hosted.com/ushosted/docs/openapi.json -g go -o src/
```

To clean up just `rm -rf ./src` as all the source code is auto-generated
