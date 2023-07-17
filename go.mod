module github.com/Unleash/unleash-server-api-go

go 1.18

require (
	github.com/antihax/optional v1.0.0
	golang.org/x/oauth2 v0.10.0
)

require (
	github.com/Unleash/unleash-server-api-go/client v0.0.0-00010101000000-000000000000 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.12.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/Unleash/unleash-server-api-go/client => ./client
