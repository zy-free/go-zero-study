
## 目标功能
- 基础
    version:1.16
    移除GOPATH,依赖GOMOD管理
- api网关（聚合层）
    - [x] (web框架)
    - [x] validate map的反射机制
    - [x] 错误码 引进B站错误码并且稍加改动，grpc层次直接返回server端错误码
- 权限
    - [x] JWT
    - [x] [Casbin] (鉴权)
- 日志
    - [ ] Metrics
    - [x] Access Log
    - [x] Tracing Opentracing+TraceID
- 服务发现
    - [x] ETCD
- 服务降级  hystrix <-> google sre
    - [ ] todo
- 服务熔断
    - [ ] todo
- 服务限流
    - [ ] todo
- 服务超时
    - [x] 全链路的超时，通过context控制
- 配置中心
    - [x] Apollo
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
- core
    - [ ] timeWheel
    - [ ] 分布式定时任务中心
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
- 基础服务
	- 日志收集
	    - [Fluentbit](https://fluentbit.io/) + Elasticsearch
		    - [ ] [docker-compose](/console/docker-compose-fb-es.yml)
		    - [ ] Kubernetes
	- [ ] 监控告警
		- Prometheus
		- Grafana 
- 发布
    - [ ] 灰度
    - [ ] 蓝绿
- 部署
	- K8S
		- [ ] [helm](/deploy/k8s/helm)
	- [ ] Docker
	    - 示例[console](/console/docker-compose.yml)
- CI/CD
	- [Drone](https://drone.io/) [README](/deploy/docker/drone)
	    - [ ] Go & Node编译
	    - [ ] Docker镜像
	    - [ ] Kubernetes发布
	    - [ ] 缓存
	- [ ] Jenkins
- ...

