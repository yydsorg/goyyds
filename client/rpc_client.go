package client

type rpcClient struct {
	opts Options
}

func (r *rpcClient) Init(option ...Option) error {

	return nil
}

func (r *rpcClient) Options() Options {
	return r.opts
}

func (r *rpcClient) Start() error {
	return nil
}

func (r *rpcClient) String() string {
	return "yyds-client"
}

func newRpcClient(opt ...Option) Client {
	opts := newOptions(opt...)
	rc := &rpcClient{opts: opts}

	return Client(rc)
}
