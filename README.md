1. 安装protobuf
```
brew install protobuf
```
2. 安装protoc-gen-go
```
go get -u github.com/golang/protobuf/protoc-gen-go
```
3. 生成proto文件
```
protoc --go_out=plugins=grpc:. ./proto/*proto
```
- --go_out: 设置所生成的go代码输出的目录，该指令会自动加载 ***protoc-gen-go*** 插件，已达到生成go代码的目的。 生成的文件已 .pb.go 为文件后缀， 这里的 ":"(冒号) 有分隔符的作用，后跟命令所需要的参数，这意味着把生成的GO代码输出到指向的protoc编译的当前目录。
- plugins=plugin1+plugin2: 指定要加载的子插件列表。 我们定义的proto文件是设计了RPC服务的，而默认是不会生成RPC代码的，因此需要在go_out中给出plugins参数，将其床底给protocol-gen-go，即告诉编译器，请支持RPC。

执行完这条命令后， 就会在proto文件夹下生成对应的.pd.go文件

4. 基本数据类型
![基本数据类型](https://gitee.com/fym321/picgo/raw/master/imgs/20201201141052.png)

5. gRPC和RESTful API 对比
![gRPC和RESTful API对比](https://gitee.com/fym321/picgo/raw/master/imgs/20201201141424.png)

6. gRPC调用方式
- Unary RPC(一元RPC): 客户端一次调用，服务端一次响应
- Server-side streaming RPC(服务端流式RPC): 客户端一次调用，可以持续接收多次服务端响应
- Client-side streaming RPC(客户端流式RPC): 客户端多次调用，服务端一次响应
- Bidirectional streaming  RPC(双向流式RPC): 客户端以流式方式发送请求，服务端以流式方式响应请求