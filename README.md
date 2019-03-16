# gRPC + REST Gateway Play

_Experimentations around grpc-gateway_

**master branch** A modified version of the grpc-gateway-example to have both HTTP and gRPC on the same port, without TLS. Server implementation is largely inspired from [Zenithar/go-password](https://github.com/Zenithar/go-password).

**gorilla-mux** How to use gorilla mux as http handler for routes stading aside of grpc-gateway.

## Installation & usage

To try it all out do this:

```bash
$ go get -u github.com/Stoakes/grpc-gateway-example
$ grpc-gateway-example serve
# In another terminal
$ grpc-gateway-example echo "my first rpc echo"
$ curl -X POST -k http://localhost:10000/v1/echo -H "Content-Type: text/plain" -d '{"value": "foo"}'
```

## Features

- A basic echo service returning value send as parameter. Available on port 10000, in both HTTP & GRPC, without TLS.
- A swagger description of the service at `localhost:10000/swagger.json`
- A swagger UI to interact with the REST API at `localhost:10000/swagger-ui/`

## Mentions from the initial README

Blog post: https://coreos.com/blog/gRPC-protobufs-swagger.html

Huge thanks to the hard work people have put into the [Go gRPC bindings][gogrpc] and [gRPC to JSON Gateway][grpcgateway]

[gogrpc]: https://github.com/grpc/grpc-go
[grpcgateway]: https://github.com/grpc-ecosystem/grpc-gateway
