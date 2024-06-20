## GRPC Usage in microservice and mono repo arch

This project demonstrate how to use GRPC for communication between microservices. We have a microservice which generates a list of traders that traded an asset, this microservice is exposes a RPC(remote procedural call) which can be called by other services.

### Project structure for easy naviagtion

- server.go : This file contains the trader grpc server implementation that return all the trades. Currently it returns a static list of trades. 

- client.go : This file contains the trader grpc client implementation that calls the server to get all the trades and returns the top 10 trades over http.

- trader.proto : This file contains the protobuf definition for the trader service.

- trader : This directory contains the generated go code from the trader.proto file.


### How to run the project

- Clone the repository
- Run the server.go file using the command `go run server.go` to run the traders grpc server.
- Run the client.go file using the command `go run client.go` to client that calls the server to get all the trades and returns the top 10 trades over http.

- The grpc server runs on port 9092 and the client runs on port 8080. The client can be accessed at `http://localhost:8083/toptraders`

  

