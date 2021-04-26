package discov

import (
	"encoding/json"
	"go-zero-study/core/discov/internal"
	"go-zero-study/core/lang"
	"go-zero-study/core/logx"
	"go-zero-study/core/proc"
	"go-zero-study/core/syncx"
	"go-zero-study/core/threading"
	"go.etcd.io/etcd/clientv3"
)

type (
	PublisherOption func(client *Publisher)

	Publisher struct {
		endpoints  []string
		key        string
		fullKey    string
		id         int64
		value      string
		lease      clientv3.LeaseID
		quit       *syncx.DoneChan
		pauseChan  chan lang.PlaceholderType
		resumeChan chan lang.PlaceholderType
	}

	RpcServiceInstance struct {
		ID        string            `json:"id"`
		Name      string            `json:"name"`
		Version   string            `json:"version"`
		Metadata  map[string]string `json:"metadata"`
		Endpoints []string          `json:"endpoints"`
	}

	Version struct {}

	ID struct {}
)


func NewPublisher(endpoints []string, key string, value *RpcServiceInstance, opts ...PublisherOption) *Publisher {
	publisher := &Publisher{
		endpoints:  endpoints,
		key:        key,
		value:      EncodeServiceInstance(value),
		quit:       syncx.NewDoneChan(),
		pauseChan:  make(chan lang.PlaceholderType),
		resumeChan: make(chan lang.PlaceholderType),
	}

	for _, opt := range opts {
		opt(publisher)
	}

	return publisher
}

func (p *Publisher) KeepAlive() error {
	cli, err := internal.GetRegistry().GetConn(p.endpoints)
	if err != nil {
		return err
	}

	p.lease, err = p.register(cli)
	if err != nil {
		return err
	}

	proc.AddWrapUpListener(func() {
		p.Stop()
	})

	return p.keepAliveAsync(cli)
}

func (p *Publisher) Pause() {
	p.pauseChan <- lang.Placeholder
}

func (p *Publisher) Resume() {
	p.resumeChan <- lang.Placeholder
}

func (p *Publisher) Stop() {
	p.quit.Close()
}

func (p *Publisher) keepAliveAsync(cli internal.EtcdClient) error {
	ch, err := cli.KeepAlive(cli.Ctx(), p.lease)
	if err != nil {
		return err
	}

	threading.GoSafe(func() {
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					p.revoke(cli)
					if err := p.KeepAlive(); err != nil {
						logx.Errorf("KeepAlive: %s", err.Error())
					}
					return
				}
			case <-p.pauseChan:
				logx.Infof("paused etcd renew, key: %s, value: %s", p.key, p.value)
				p.revoke(cli)
				select {
				case <-p.resumeChan:
					if err := p.KeepAlive(); err != nil {
						logx.Errorf("KeepAlive: %s", err.Error())
					}
					return
				case <-p.quit.Done():
					return
				}
			case <-p.quit.Done():
				p.revoke(cli)
				return
			}
		}
	})

	return nil
}

func (p *Publisher) register(client internal.EtcdClient) (clientv3.LeaseID, error) {
	resp, err := client.Grant(client.Ctx(), TimeToLive)
	if err != nil {
		return clientv3.NoLease, err
	}

	lease := resp.ID
	if p.id > 0 {
		p.fullKey = makeEtcdKey(p.key, p.id)
	} else {
		p.fullKey = makeEtcdKey(p.key, int64(lease))
	}
	_, err = client.Put(client.Ctx(), p.fullKey, p.value, clientv3.WithLease(lease))

	return lease, err
}

func (p *Publisher) revoke(cli internal.EtcdClient) {
	if _, err := cli.Revoke(cli.Ctx(), p.lease); err != nil {
		logx.Error(err)
	}
}

func WithId(id int64) PublisherOption {
	return func(publisher *Publisher) {
		publisher.id = id
	}
}

// 编码服务实例信息
func EncodeServiceInstance(v *RpcServiceInstance) string {
	value, _ := json.Marshal(v)
	return string(value)
}

// 解码服务实例信息
func DecodeServiceInstance(bs []byte) (*RpcServiceInstance, error) {
	var serviceInstance RpcServiceInstance
	err := json.Unmarshal(bs, &serviceInstance)
	if err != nil {
		return nil, err
	}

	return &serviceInstance, nil
}
