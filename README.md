# gRPC hands on
These are examples for my learning in gRPC

# pre-requirement
1. install `protobuf`
```
brew install protobuf
```
2. install `gRPC`
```
go get google.golang.org/grpc
```
3. install `protoc-gen-go`(because this go package is based on protoc)
```
go get -u github.com/golang/protobuf/protoc-gen-go
```

# Content
1. Unary
  
  in `greet` directory show example that simple unary mode un gRPC. server would append `hello` with message, then send back to client. these are includes:
  - create protobuf file for communication
  - Generate go file(message, request, response, service)
  - How hook new Service to own "Server"
  - How implement a client call for Unary RPC

2. Server streaming
3. client streaming
4. Bi-Directional streaming
