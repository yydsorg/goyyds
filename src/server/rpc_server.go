package server

type rpcServer struct {
	opts Options
}

func (r *rpcServer) Version() string {
	return r.opts.Version
}

func (r *rpcServer) Start() error {
	panic("implement me")
}

func (r *rpcServer) Stop() error {
	panic("implement me")
}

func (r *rpcServer) String() string {
	panic("implement me")
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

func newRpcServer(opts ...Option) Server {
	options := newOptions(opts...)
	return &rpcServer{
		opts: options,
	}
}
