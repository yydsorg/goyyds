package client

type rpcClient struct {
	opts Options
}

func (r rpcClient) Init(option ...Option) error {
	panic("implement me")
}

func (r rpcClient) Options() Options {
	panic("implement me")
}

func (r rpcClient) String() string {
	panic("implement me")
}

func newRpcClient(opt ...Option) Client {
	opts := newOptions(opt...)
	rc := rpcClient{opts: opts}

	c := Client(rc)
	return c
}
