# Go gRPC quickstart

Setup and Documentation: https://grpc.io/docs/languages/go/quickstart/

(RE)Generate gRPC transport layer Golang code

`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false proto/service.proto`

Run server
`go run main.go`

Run client 
`go run client/client.go`

Docker build
`docker build -t grpc-quickstart:latest . `

Docker run
`docker run -p 50051:50051 grpc-quickstart:latest`