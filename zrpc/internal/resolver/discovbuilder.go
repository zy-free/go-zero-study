package resolver

import (
	"google.golang.org/grpc/attributes"
	"strings"

	"go-zero-study/core/discov"
	"google.golang.org/grpc/resolver"
)

type discovBuilder struct{}

func (d *discovBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (
	resolver.Resolver, error) {
	hosts := strings.FieldsFunc(target.Authority, func(r rune) bool {
		return r == EndpointSepChar
	})
	sub, err := discov.NewSubscriber(hosts, target.Endpoint)
	if err != nil {
		return nil, err
	}

	update := func() {
		var addrs []resolver.Address
		for _, val := range subset(sub.Values(), subsetSize) {
			serverInstance, err := discov.DecodeServiceInstance([]byte(val))
			if err != nil {
				continue
			}

			 metedates :=  []interface{}{discov.Version{}, serverInstance.Version, discov.ID{} ,serverInstance.ID}
			for k, v := range  serverInstance.Metadata {
				metedates = append(metedates, k, v)
			}

			addrs = append(addrs, resolver.Address{
				Addr: serverInstance.Endpoints[0],
				ServerName: serverInstance.Name,
				Attributes: attributes.New(metedates...),
			})
		}
		cc.UpdateState(resolver.State{
			Addresses: addrs,
		})
	}
	sub.AddListener(update)
	update()

	return &nopResolver{cc: cc}, nil
}

func (d *discovBuilder) Scheme() string {
	return DiscovScheme
}
