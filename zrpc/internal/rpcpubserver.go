package internal

import "go-zero-study/core/discov"

func NewRpcPubServer(etcdEndpoints []string, etcdKey string, serviceInstance *discov.RpcServiceInstance, opts ...ServerOption) (Server, error) {
	registerEtcd := func() error {
		pubClient := discov.NewPublisher(etcdEndpoints, etcdKey, serviceInstance)
		return pubClient.KeepAlive()
	}
	server := keepAliveServer{
		registerEtcd: registerEtcd,
		Server:       NewRpcServer(serviceInstance.Endpoints[0], opts...),
	}

	return server, nil
}

type keepAliveServer struct {
	registerEtcd func() error
	Server
}

func (ags keepAliveServer) Start(fn RegisterFn) error {
	if err := ags.registerEtcd(); err != nil {
		return err
	}

	return ags.Server.Start(fn)
}
