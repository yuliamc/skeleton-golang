api-server:
	go run ./cmd/apiserver/main.go server

api-config:
	go run ./cmd/apiserver/main.go config

build-api-server:
	go build ./cmd/apiserver/main.go