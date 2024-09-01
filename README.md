This project is just a challenge for me, full stack gRPC

1. First
```shell
# go version: 1.23.0
cd BackEnd
go mod tidy
```
```shell
# 这个要单独运行,  文件地址问题的原因
protoc --go_out=BackEnd\ --go-grpc_out=BackEnd\ BackEnd\proto\forum.proto
protoc --go_out=BackEnd\ --go-grpc_out=BackEnd\ BackEnd\proto\chat.proto
protoc --go_out=BackEnd\ --go-grpc_out=BackEnd\ BackEnd\proto\auth.proto

# 当需要cpp版本功能重载时请执行
cd BackEndCpp
protoc --cpp_out=. --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` proto/forum.proto

```

2. Second change `config/__app.yml` name to `config/app.yml`, and set the database in this file.

3. Third
```shell
go run main.go
```

```shell
cd FrontEnd
npm install
npm install -g protoc-gen-grpc-web
cd ..
protoc -I ./BackEnd/proto ./BackEnd/proto/forum.proto --js_out=import_style=commonjs:./FrontEnd/src/proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./FrontEnd/src/proto
```

现在是笔记时间:
1. 首先我在main.go中通过viper这个包读取yml文件的内容, 然后更新到配置中, 接下来进行数据库迁移的操作,可以确保每一次启动服务的时候数据库和我的结构体是对应的.
2. 然后我通过`proto/forum.proto`文件生成了`proto/forum.pb.go`和`proto/forum_grpc.pb.go`文件, 这两个文件是我在golang中使用grpc的时候需要的,
通过这两个文件我可以在golang中使用grpc的服务. 其中的内容是`ForumService`的接口.
3. 接下来在`service/forumService.go`中实现了`ForumService`的接口, 这个接口是我在`proto/forum.proto`文件中定义的, 通过这个接口我可以实现grpc的服务.

## License

This project is licensed under the Creative Commons Attribution-NonCommercial 4.0 International License. See the [LICENSE](LICENSE) file for more details.
