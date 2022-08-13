# Graduate Project

## 题目

对当下自己项目中的业务，进行一个微服务改造，需要考虑如下技术点：

- 微服务架构（BFF、Service、Admin、Job、Task 分模块）

  具体体现在[服务分层](#服务分层)部分

- API 设计（包括 API 定义、错误码规范、Error 的使用）

  具体体现在[接口与错误定义](#接口与错误定义)部分

- gRPC 的使用

  具体体现在[业务代码结构](#业务代码结构)部分

- Go 项目工程化（项目结构、DI、代码分层、ORM 框架）

  具体体现在[业务代码结构](#业务代码结构)部分

- 并发的使用（errgroup 的并行链路请求）

  具体体现在[statistics-service](#statistics-service)部分

- 微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）
    <br>

- 缓存的使用优化（一致性处理、Pipeline 优化）

  具体体现在[statistics-service](#statistics-service)的查询用户统计部分

## 项目设定

实现一个登入应用，包含以下内容
* 一般用户
  * 注册与登录
* 管理端
  * 查看指定用户注册和登入统计




## 框架使用

微服务框架

* kratos version v2.4.1

数据库部分

| 数据库    | ORM                 |
| --------- | ------------------- |
| mysql 5.7 | entgo.io/ent v0.9.1 |

中间件部分

| 中间件       | SDK                       |
| ------------ | ------------------------- |
| redis 5.0.7  | go-redis/redis/v8 |



## 项目结构

### 服务分层

针对[项目设定](#项目设定)，对项目进行以下分层

BFF

| 名称              | 描述                                        | 代码位置                                                     |
| ----------------- | ------------------------------------------- | ------------------------------------------------------------ |
| bff-login | 普通用户BFF接口，提供http与grpc协议的接口   | 业务逻辑: /app/login/interface<br>接口定义: /api/login/interface/v1 |



内部微服务

| 名称             | 描述                                                         | 代码位置                                                     |
| ---------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| user-service     | 用户服务，管理用户信息，提供grpc协议的接口                   | 业务逻辑: /app/user/service<br/>接口定义: /api/user/v1       |
| statistics-service | 统计用户信息服务，提供grpc协议的接口 | 业务逻辑: /app/statistics/service<br/>接口定义: /api/statistics/service/v1 |



#### 接口与错误定义

使用proto buffer来定义接口与接口返回的错误，放置于`/api/<服务名称>/<服务类型>/<接口版本>`路径下，分别使用`*.proto`来定义接口信息和`*_error.proto`来定义错误信息

在定义好接口与错误信息后，使用`kratos proto client <protobuf文件路径>`来对proto buffer进行go文件转译

实际目录接口举例：

```bash
api/login/interface/
└── v1
    ├── login.pb.go
    ├── login_error.proto
    ├── login_error.proto
    ├── login_grpc.pb.go
    └── login.proto
```



#### 业务代码结构


* cmd

  项目启动文件main.go

  构建依赖关系的wire.go，使用[google/wire](https://github.com/google/wire)项目来实现项目启动时的依赖注入的生命周期管理

* configs

  采取yaml文件定义配置信息，如服务启动端口，数据库连接信息等

* internal/biz

  定于以下内容：

  1. 领域对象
  2. 服务用例
  3. 与外部数据源进行交互的组件的接口

* internal/conf

  使用proto buffer定义需要读取的配置信息的格式

* internal/data

  基于internal/biz定义的与外部数据源进行交互的组件的接口，实现具体的数据交互的逻辑

* internal/service

  基于1. internal/service中定义的服务用例， 2.[api](#接口与错误定义)中protobuf编译后的go文件接口，实现以下内容：

  * 处理与传递外部请求所发过来数据
  * 调用biz层中的服务用例来响应请求，并返回请求结果

* internal/server

  * 定义http与grpc服务对象的创建方法


#### statistics-service

主要实现以下服务接口(http)

* 查询指定用户指定日期登入记录

  这里采取了并行请求的模式，每次请求的内容为单个用户的多个日期的登入记录，使用errgroup来统一管理错误信息，实现如下

```go
func (uc *ClockinUsecase) GetTime(ctx context.Context, user []int64, day []int64) ([]*UserTime, error) {
    time := make([]*UserTime, len(user))
	g, ctx := errgroup.WithContext(ctx)
	for k, v := range user {
		index := k
		userId := v
		g.Go(func() error {
			reply, err := uc.repo.GetUserTime(ctx, userId, day)
			if err != nil {
				return err
			}
			time[index] = &UserTime{
				User:     userId,
				Worktime: reply,
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return time, nil
}
```






