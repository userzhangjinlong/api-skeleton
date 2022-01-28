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
    1.生成表结构struct
    go run cmd/Cli/main.go sql struct --schema=库名称 --table=表名称
~~~