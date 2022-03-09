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

- 常用组件工具
~~~
    1.生成表结构model struct 到app/Model/库名称/表名称.go文件
    go run cmd/Cli/main.go sql struct --schema=生产库名称 --table=生成表名称
~~~

- 实现功能
~~~
    1.cobra命令行脚本
    2.redis连接、redis-cluster集群
    3.鉴权、跨域、访问日志等中间件
    4.validator验证器
    5.全链路追踪
    7.mysql分库集群链接，读写分离
    8.nsq消息投递 (待实现)
~~~

- 待优化
~~~
    1.路由传递
    2.路由中间件分组
    3.nsq消息对了消费消息添加入库数据时消息50条会造成db报错 Too many connections
    这样也会造成消息消费失败；消息消费失败的重试和补偿机制待处理（lookupdauth？？）
~~~