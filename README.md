# gRPC + REST Gateway Play

A modified version of the grpc-gateway-example to have both HTTP and gRPC on the same port, without TLS.

Server implementation is largely inspired from [Zenithar/go-password](https://github.com/Zenithar/go-password).

## README

Blog post: https://coreos.com/blog/gRPC-protobufs-swagger.html

To try it all out do this:

```
$ go get -u github.com/Stoakes/grpc-gateway-example
$ grpc-gateway-example serve
$ grpc-gateway-example echo "my first rpc echo"
$ curl -X POST -k http://localhost:10000/v1/echo -H "Content-Type: text/plain" -d '{"value": "foo"}'
{"value":"my REST echo"}
```


Huge thanks to the hard work people have put into the [Go gRPC bindings][gogrpc] and [gRPC to JSON Gateway][grpcgateway]

[gogrpc]: https://github.com/grpc/grpc-go
[grpcgateway]: https://github.com/grpc-ecosystem/grpc-gateway
