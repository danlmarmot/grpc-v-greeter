build: generate
	go build ./cmd/greeter/...

generate:
	mkdir -p protogen/v1/greeter; protoc -I proto --proto_path=proto/v1 --go_out=protogen/v1/greeter --go-grpc_out=protogen/v1/greeter proto/v1/greeter_api.proto
	mkdir -p protogen/v2/greeter; protoc -I proto --proto_path=proto/v2 --go_out=protogen/v2/greeter --go-grpc_out=protogen/v2/greeter proto/v2/greeter_api.proto