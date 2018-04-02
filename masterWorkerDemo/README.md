# master-worker demo

An examples of using grpc-gateway and go third-party packages. Provides remote target management.

## master 

grpc gateway server: 创建http server，并注册到gw.go中。同时作为pb.go的client端转发请求给worker。

## worker

target worker: 继承pb.go的interface（Add, List, Delete等interface），并注册为pb.go的server端，接受内部请求后处理具体的功能，并返回消息。

## run

- [master](/cmd/master): go run main.go
- [worker](/cmd/worker): go run main.go
