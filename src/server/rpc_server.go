package server

type rpcServer struct {
	opts Options
}

func (r *rpcServer) Init(opts ...Option) error {
	for _, opt := range opts {
		opt(&r.opts)
	}

	return nil
}

func (r *rpcServer) Options() Options {

	return r.opts
}

func newRpcServer() Server {
	return &rpcServer{}
}
