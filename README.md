# envoy_example

A working example of an External Auth filter wired into Envoy.   This example uses an Envoy image from 3 years ago. The current Envoy build is a bit different and I have not gotten that to work yet.

URL of example from which you built this project:
https://gerswayne.medium.com/simple-dockerized-grpc-application-with-envoy-ext-authz-example-6c200e0a2d34


To fire up the server/main.go service sitting behind Envoy's External Authorization Mechanism, from the project root:
```docker-compose up```

To fire up a gRPC Client that issues a single RPC request to Envy external authorization as a proxy to the server/main.go, from the project root:
```go run ./client/main.go```

To fire up a shell script that issues 1000 requests to the preceding, from the project root:
```go run ./client/main.go```
You might get failures, I think because of timeout settings in the Envoy lashup.