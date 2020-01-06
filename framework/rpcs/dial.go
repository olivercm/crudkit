package rpcs

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Dial(name string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	authority := resolveAuthority(name)
	if authority == "" {
		return nil, fmt.Errorf("cannot resolve authority: %s", name)
	}
	timeout := time.Duration(0)
	for i := 0; i < len(opts); i++ {
		if to, ok := opts[i].(*rpcTimeout); ok {
			timeout = to.d
			break
		}
	}
	opts = append(opts, grpc.WithInsecure(), grpc.WithAuthority(authority))
	if timeout != 0 {
		i := &timeoutInterceptor{timeout: timeout}
		opts = append(opts, grpc.WithUnaryInterceptor(i.invoke))
	}
	return grpc.Dial(fmt.Sprintf("conf:///%s", name), opts...)
}

func MustDial(name string, opts ...grpc.DialOption) *grpc.ClientConn {
	cc, err := Dial(name, opts...)
	if err != nil {
		log.Panicf("dial %s: %v", name, err)
	}
	return cc
}
