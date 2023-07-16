# Generating the client:

```bash
./generate.sh
```

After generating the code you can use `main.go` to test it works (remember updating unleash servere host and API token)

```bash
go run main.go -unleash-api http://localhost:4242 -authorization user:a3c8e3e76e7361c0bc79070accf70409d03fbaef4c2b6b90bff466e8
```
