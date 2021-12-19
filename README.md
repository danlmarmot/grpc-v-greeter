# gRPC Demo Application

Based on https://github.com/rokane/grpc-demo

This is an example application used to demonstrate an approach to structuring
gRPC servers such that they can handle multiple API versions.

Post is here: [Building APIs with gRPC and Go](https://medium.com/@ryan.okane8/building-apis-with-grpc-and-go-9a6d369d7ce)

The gRPC server itself contains two implementations of the `Greeter` service.
Both services contain the same two endpoints (as explained below), but with
different message types in an attempt to illustrate a `breaking change` leading
to a new version (`v2`).

## Generate protobufs

Manually:

```bash
protoc -I proto --proto_path=proto/v1 --go_out=protogen/v1/greeter --go-grpc_out=protogen/v1/greeter proto/v1/greeter_api.proto
protoc -I proto --proto_path=proto/v2 --go_out=protogen/v2/greeter --go-grpc_out=protogen/v2/greeter proto/v2/greeter_api.proto
```
or use the makefile with `make generate`

## Build and Run

Build executable.

```bash
> go build ./cmd/greeter/...
> ./greeter
```

... or use the makefile with `make build`

## Verify 

Call the GRPC server in another terminal.

```bash
# API v1
grpcurl -plaintext -v -d '{"name": "ryan"}' localhost:8080 greeter.v1.Greeter/SayHello
grpcurl -plaintext -v -d '{"name": "ryan"}' localhost:8080 greeter.v1.Greeter/SayGoodbye

# API v2
grpcurl -plaintext -v -d '{"first_name": "ryan", "last_name": "okane"}' localhost:8080 greeter.v2.Greeter/SayHello
grpcurl -plaintext -v -d '{"first_name": "ryan", "last_name": "okane"}' localhost:8080 greeter.v2.Greeter/SayGoodbye
```
