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
  
  client send `one` request, server send back `one` response.

  in `greet` directory show example that simple unary mode un gRPC. server would append `hello` with message, then send back to client. these are includes:
  - create protobuf file for communication
  - Generate go file(message, request, response, service)
  - How hook new Service to own "Server"
  - How implement a client call for Unary RPC

  How run up example
  - start the server
  ```
  go run greet/greet_server/server.go
  ```
  - start the client
  ```
  go run greet/greet_client/client.go

2. Server streaming

  client send `one` request, server send back `many` response, possibly and infinite number.
  in `greet-server-streaming` directory show example that simple server streaming mode un gRPC. it only different with Unary mode in `proto` file definition. marked `stream` for GreetResponse
  ```proto
  service GreetService {
    rpc GreetManyTimes (GreetRequest) returns (stream GreetResponse);
  }
  ```

  How run up example
  - start the server
  ```
  go run greet-server-streaming/greet_server/server.go
  ```
  - start the client
  ```
  go run greet-server-streaming/greet_client/client.go
  ```

3. client streaming

  client send `many` requests, server send back `one` response.
  - example1: Say hello with all names

    in `greet-client-streaming` directory show example that is client send three names as `stream` to server, server response message which includes names from multip-requests.
    ```proto
    service GreetService {
      rpc LongGreet (stream GreetRequest) returns (GreetResponse);
    }
    ```
    How run up example
    - start the server
    ```
    go run greet-client-streaming/greet_server/server.go
    ```
    - start the client
    ```
    go run greet-client-streaming/greet_client/client.go
    ```

  - example2: calculating average number

    calculating average number by `many` number from client `stream`.

    ```proto
    service GreetService {
      rpc Average (stream CalculateRequest) returns (CalculateResponse) ;  
    }
    ```
    How run up example
    - start the server
    ```
    go run greet-client-streaming/greet_server/server.go
    ```
    - start the client
    ```
    go run greet-client-streaming/average_client/average_client.go
    ```

4. Bi-Directional streaming
  client send `many` requests, server send back `many` response too.
