package rpcs

import (
	"google.golang.org/grpc/resolver"
	"sync/atomic"
)

var config struct {
	services atomic.Value
}

func resolveFromConf(service string) string {
	return config.services.Load().(map[string]string)[service]
}

func resolveAuthority(service string) string {
	return resolveFromConf(service)
}

type resolverImpl struct {
	target resolver.Target
	cc     resolver.ClientConn
}

func (r *resolverImpl) ResolveNow(resolver.ResolveNowOption) {
	addr := resolveFromConf(r.target.Endpoint)
	if addr != "" {
		r.cc.UpdateState(resolver.State{
			Addresses: []resolver.Address{{Addr: addr}},
		})
	}
}

func (r *resolverImpl) Close() {
}

type builder struct {
}

func (*builder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOption) (resolver.Resolver, error) {
	r := &resolverImpl{
		target: target,
		cc:     cc,
	}
	r.ResolveNow(resolver.ResolveNowOption{})
	return r, nil
}

func (*builder) Scheme() string {
	return "conf"
}

func init() {
	resolver.Register(&builder{})
}
