
## 目标功能
- 基础
    version:1.15
    移除GOPATH,依赖GOMOD管理
- （聚合层）
    - [x] (web框架)
    - [x] validate map的反射机制
    - [x] 错误码 引进B站错误码并且稍加改动，grpc层次直接返回server端错误码
    - [x] metadata元数据

- demo
    - [ ] admin+gorm(share-db架构)
    - [x] service+sql
    - [ ] job+Beanstalkd
    - [ ] 数据库delete_time 为null的时间处理，自定义时间
    - [ ] 多租户 (流量染色)
    - [ ] errgroup并发请求
    - [ ] thread包使用
    - [x] excel返回
    - [x] 失血模型
    - [ ] 获取天气的等api的demo，client以及分层设计
    
- 日志
    - [ ] Metrics
    - [x] Access Log
    - [x] Tracing Opentracing+TraceID
- 监控
    - [ ] promethues监控+模板
- 服务发现
    - [x] ETCD
- 服务熔断
    - [x] grpc客户端层 break拦截器=>google sre
    - [x] http层 breakerHandler=>google sre
- 服务限流
    - [x] 当CPU>90%的时候开始拒绝请求
    - [x] shedding 拦截器，包括grpc服务端层，http层
- 服务超时
    - [x] 全链路的超时，通过context控制
- 负载均衡
    - [x] p2c算法
- 领域驱动
	- [x] 整洁架构
- cache
	- [x] redis
    - [x] 击穿  内存中lock然后共享数据
    - [x] 雪崩  过期时间在基础值上+了个随机值，防止大量失效
    - [x] 穿透  缓存一条内容为*,60s过期的数据，避免高并发访问数据库
    - [x] where转化主键

- stat采集
    - [x] cpu
    - [x] memory 
    
        `通过golang的runtime.MemStats实现`
        
        `其中不通的操作系统通过go build的条件编译实现`
         
    - [x] 接口或者rpc的total,pass,drop
    
        `SheddingHandler http的middleware里注入来统计`
        
        `rpc同理，通过server的UnarySheddingInterceptor来注入`
    - [x] metrics
    
        `api通过bindRoutes时new`
        
        `rpcServer通过NewRpcServer时new，其中metrics的name通过server.SetName(c.Name)来改变 ` 

