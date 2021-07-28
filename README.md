### golang chat room，go语言实现的聊天室
support 100000 users online chat in a single server the same time.
单台10万用户同时在线 

### env, 环境

- Linux
- Go 1.2+, 语言1.2版本或主线源码编译
- mem 8GB, 内存8GB
- openfiles > 100000, 数设为100000以上

### download,下载
go get github.com/ablozhou/chatit

### compile 编译
```
make
andy@minta ~/chatit $ make
go install
```

### usage 用法
start server 启动服务器
```
>andy@minta ~/chatit $ chatit
>Wrong parameter,usage:
>
>chatit server [port]
>    eg: chatit server 9090
chatit client [Server_Addr]:[Server_Port]
    eg: chatit client 192.168.0.74:9090
chatit client [Server_Addr]:[Server_Port] [count]
    eg: chatit client 192.168.0.74:9090 500

```
