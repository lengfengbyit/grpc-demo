1. 生成proto文件
```
protoc --go_out=plugins=grpc:. ./proto/*proto
```
- --go_out: 设置所生成的go代码输出的目录，该指令会自动加载 ***protoc-gen-go*** 插件，已达到生成go代码的目的。 生成的文件已 .pb.go 为文件后缀， 这里的 ":"(冒号) 有分隔符的作用，后跟命令所需要的参数，这意味着把生成的GO代码输出到指向的protoc编译的当前目录。
- plugins=plugin1+plugin2: 指定要加载的子插件列表。 我们定义的proto文件是设计了RPC服务的，而默认是不会生成RPC代码的，因此需要在go_out中给出plugins参数，将其床底给protocol-gen-go，即告诉编译器，请支持RPC。

执行完这条命令后， 就会在proto文件夹下生成对应的.pd.go文件