# api-skeleton

- 一款适用于phper学习和使用的go，api骨架项目
- 项目目录结构
~~~
api-skeleton
├── app
│   ├── Cache
│   ├── Console
│   ├── Enum
│   ├── Ecode
│   ├── ConstDir
│   ├── Global
│   ├── Http
│   │	├── Controller
│   │	├── Middleware
│   │   └── Request
│   ├── Libraries
│   ├── Model
│   ├── Logic
│   ├── Listener
│   ├── Util
│   └── Service
├── bootstrap
├── config
├── database
├── routes
├── cmd
├── DockerFile
├── storage
├── makefile
├── env.yaml
└── test 
~~~

# grpc&protobuf
- 安装
***
Grpc
~~~
    golang grpc安装自行科学搜索
~~~
***
protobuf
~~~
    ###仅仅针对macOs 其他环境自行科学搜索参照流程
    brew install protobuf
    go install github.com/golang/protobuf/protoc-gen-go@latest
    初始化protoc-gen-go到$GOROOT/bin目录
    go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
~~~

***
- 使用
~~~
    坑点：
        - 容易出现 google/api/annotations.proto not found问题，需要google/api下载到本地$PATH/src,idea导入
        ![img.png](img.png)
    1 生成pb .go 定义proto文件生成是不需要引入google/api/annotations.proto 和response的 option
    对应proto指定 protoc --go_out=plugins=grpc:. user.proto
    2.pb.gw.go->定义生成需要引入google/api/annotations.proto 和response的 option
    protoc -I/usr/local/include -I. \
       -I$GOPATH/src \
       -I$GOPATH/src/google/api \
       --grpc-gateway_out=logtostderr=true:. \
       ./proto/user/*.proto
    3.定义自己的service或者其他之类的想命名的文件，实现service里的rpc，就可以开始愉快的进行rpc交互开发啦
~~~

# 常用组件工具
~~~
    1.生成表结构model struct 到app/Model/库名称/表名称.go文件
    go run cmd/Cli/main.go sql struct --schema=生产库名称 --table=生成表名称
    
    2.声明一个nsq消费者 -T topic -C chanel
     go run cmd/Msg/main.go -T createRankingMessage -C createRankingMessage
     
    3.声明kafka消费者
    go run cmd/Msg/main.go -MT kafka -T kafka-test-1
    
    4.声明rabbitmq消费者
    go run cmd/Msg/main.go -MT rabbitMq -T testQueue
~~~

# 实现功能
~~~
    1.cobra命令行脚本
    2.redis连接、redis-cluster集群
    3.鉴权、跨域、访问日志等中间件
    4.logrus日志切片，分割
    5.validator验证器
    6.全链路追踪
    7.mysql分库集群链接，读写分离
    8.nsq消息生产消费。nsq集群
    9.kafka消息生产者和消费者。
    10.rabbitmq消息
    11. gRPC&protobuf.远程过程调用（待实现）
~~~
## 待优化
~~~
    1.路由传递 （考虑使用闭包函数切路由）
    2.路由中间件分组
    3.nsq消息对了消费消息添加入库数据时消息50条会造成db报错 Too many connections
    这样也会造成消息消费失败；消息消费失败的重试和补偿机制待处理（lookupdauth？？）
    4.nsq消息消费幂等，消息体代码修改重置消息平滑重启ack等问题的整理处理
    5.匿名结构体嵌套 模拟实现继承
~~~