package rpcs

import (
	"context"
	"google.golang.org/grpc"
	"time"
)

type rpcTimeout struct {
	d time.Duration
	grpc.EmptyDialOption
	grpc.EmptyCallOption
}

func dummyCancel() {
}

type timeoutInterceptor struct {
	timeout time.Duration
}

func (i *timeoutInterceptor) invoke(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	timeout := i.timeout
	for i := 0; i < len(opts); i++ {
		if to, ok := opts[i].(*rpcTimeout); ok {
			timeout = to.d
			break
		}
	}
	cancel := dummyCancel
	if timeout != 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
	}
	err := invoker(ctx, method, req, reply, cc, opts...)
	cancel()
	return err
}
