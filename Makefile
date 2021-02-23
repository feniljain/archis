rest:
	fresh -c ./fresh.conf

api:
	go run api/main.go

gen:
	protoc -I proto/ --proto_path ~/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis/ proto/*.proto --go_out=plugins=grpc:pb --grpc-gateway_out=logtostderr=true:pb/
	# protoc -I proto/ --proto_path $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ proto/*.proto --go_out=plugins=grpc:pb --grpc-gateway_out=logtostderr=true:proto/


clean:
	rm -rf pb/*.go

server:
	go run cmd/server/main.go --port 8081

client:
	go run cmd/client/main.go --address 0.0.0.0:8081

test:
	go test -cover -race ./...

.PHONY: gen clean server client test
