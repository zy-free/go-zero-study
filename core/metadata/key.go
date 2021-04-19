package metadata

// metadata common key
const (

	// Network
	RemoteIP   = "remote_ip"
	RemotePort = "remote_port"

	// Router
	Cluster = "cluster"
	Color   = "color"

	// Trace
	Trace  = "trace"

	// Mid 外网账户用户id
	Mid = "mid" // NOTE: ！！！业务可重新修改key名！！！

)

var outgoingKey = map[string]struct{}{
	Color:       {},
	RemoteIP:    {},
	RemotePort:  {},
	Mid:{},
}

var incomingKey = map[string]struct{}{
	Mid:{},
}

// IsOutgoingKey represent this key should propagate by rpc.
func IsOutgoingKey(key string) bool {
	_, ok := outgoingKey[key]
	return ok
}

// IsIncomingKey represent this key should extract from rpc metadata.
func IsIncomingKey(key string) (ok bool) {
	_, ok = outgoingKey[key]
	if ok {
		return
	}
	_, ok = incomingKey[key]
	return
}
