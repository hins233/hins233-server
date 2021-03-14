# web开发

### web 服务最核心的是什么呢？
当然是业务啦，那么如何更好的服务这些业务呢，如何更高效的开发和迭代业务发展，是trpc-go的使命。
trpc-go 提供了服务业务的辅助功能，例如项目模板创建，rpc代理接口封装，数据编码，路由，配置管理，插件装载，请求超时等。

1. 通过trpc [command] 可以生成项目服务端工程结构和客户端请求的stub代码桩。

代码桩主要描述了客户端或者rpc调用方，本地有一个实现服务端一样方法的代理对象，
可以按照pb中定义的请求参数和响应参数请求到真实服务端实例，并获取服务端响应的结果和错误信息。

服务端：服务端需要实现.proto中定义的方法，并启动一个gRPC服务器用于处理客户端请求。gRPC反序列化到达的请求，执行服务方法，序列化服务端响应并发送给客户
客户都：客户端本地有一个实现了服务端一样方法的对象，gRPC中称为桩，其他语言中更习惯称为客户端。客户调用本地桩的方法，将参数按照合适的协议封装并将请求发送给服务端，并接收服务端的响应。



讲的内容：
第一，trpc-go做了什么：主要是看他生成了那些代码:
1. 桩代码，pb.go 描述proto声明的对象，相当于描述符，别称
2. trpc.go 实现自定义的client，通过自定义的client.invoke调用rpc服务端。。。grpc/grpc-gateway 使用的是clientConn.invoke。
create

mux : 多路复用处理http请求和routing路由和报文dispatch分发
proto 定义rpc服务和传输的proto msg类型信息
pb.go

### grpc-gateway 
grpc-gateway是protoc的一个插件。它读取gRPC服务定义，并生成一个反向代理服务器，将RESTful JSON API转换为gRPC。此服务器是根据gRPC定义中的自定义选项生成的。
主要生成xxx.pb.gw.go文件，里面定义了handle method的注册，和直接通过
multiplexer

// curl -X POST 127.0.0.1:8080/rpc/path/hinstang/mchannel/receive-order/rpc
// -H "Content-Type: application/json" -X POST -d '{"user_id": "123", "coin":100, "success":1, "msg":"OK!" }'

