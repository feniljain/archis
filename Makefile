rest:
	fresh -c ./fresh.conf

api:
	go run api/main.go

gen:
	protoc -I proto/ proto/*.proto --go_out=plugins=grpc:pb

clean:
	rm -rf pb/*.go

server:
	go run cmd/server/main.go --port 8080

client:
	go run cmd/client/main.go --address 0.0.0.0:8080

test:
	go test -cover -race ./...

.PHONY: gen clean server client test