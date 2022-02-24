# Npool go service app template

[![Test](https://github.com/NpoolPlatform/go-service-app-template/actions/workflows/main.yml/badge.svg?branch=master)](https://github.com/NpoolPlatform/go-service-app-template/actions/workflows/main.yml)

[目录](#目录)
- [Npool go service app template](#npool-go-service-app-template)
    - [功能](#功能)
      - [sync](#sync)
      - [async](#async)
    - [设计](#设计)
    - [命令](#命令)
    - [步骤](#步骤)
    - [最佳实践](#最佳实践)
    - [关于mysql](#关于mysql)
    - [优化](#优化)

-----------
### 功能

#### sync
- [x] 创建钱包
- [x] 查询余额

#### async
- [x] 转账
- [x] 上报币种

### 设计

Transaction:

| Field          | Type              | Unique | Optional | Nillable | Default | UpdateDefault | Immutable | StructTag                       | Validators |
| :------------- | :---------------- | :----- | :------- | :------- | :------ | :------------ | :-------- | :------------------------------ | :--------- |
| id             | string            | true   | false    | false    | true    | false         | false     | json:"id,omitempty"             | 0          |
| nonce          | uint64            | false  | false    | false    | true    | false         | false     | json:"nonce,omitempty"          | 0          |
| coin_type      | int8              | false  | false    | false    | true    | false         | false     | json:"coin_type,omitempty"      | 0          |
| transaction_id | string            | false  | false    | false    | true    | false         | false     | json:"transaction_id,omitempty" | 1          |
| from           | string            | false  | false    | false    | true    | false         | false     | json:"from,omitempty"           | 1          |
| to             | string            | false  | false    | false    | true    | false         | false     | json:"to,omitempty"             | 1          |
| value          | float64           | false  | false    | false    | true    | false         | false     | json:"value,omitempty"          | 0          |
| state          | transaction.State | false  | false    | false    | false   | false         | false     | json:"state,omitempty"          | 0          |
| create_at      | uint32            | false  | false    | false    | true    | false         | false     | json:"create_at,omitempty"      | 0          |
| update_at      | uint32            | false  | false    | false    | true    | true          | false     | json:"update_at,omitempty"      | 0          |
| delete_at      | uint32            | false  | false    | false    | true    | false         | false     | json:"delete_at,omitempty"      | 0          |

### 命令
* make init ```初始化仓库，创建go.mod```
* make verify ```验证开发环境与构建环境，检查code conduct```
* make verify-build ```编译目标```
* make test ```单元测试```
* make generate-docker-images ```生成docker镜像```
* make service-sample ```单独编译服务```
* make service-sample-image ```单独生成服务镜像```
* make deploy-to-k8s-cluster ```部署到k8s集群```

### 步骤

### 最佳实践
* 每个服务只提供单一可执行文件，有利于docker镜像打包与k8s部署管理
* 每个服务提供http调试接口，通过curl获取调试信息
* 集群内服务间direct call调用通过服务发现获取目标地址进行调用
* 集群内服务间event call调用通过rabbitmq解耦

### 关于mysql
* 创建app后，从app.Mysql()获取本地mysql client
* [文档参考](https://entgo.io/docs/sql-integration)

### 优化

对离线钱包的操作抽象接口，从服务中剥离出来
