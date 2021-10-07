package server

type rpcServer struct {
	opts Options
}

func newRpcServer(opts ...Option) Server {
	options := newOptions(opts...)
	return &rpcServer{
		opts: options,
	}
}

func (r *rpcServer) Version() string {
	return r.opts.Version
}

func (r *rpcServer) Start() error {
	return nil
}

func (r *rpcServer) Stop() error {
	return nil
}

func (r *rpcServer) String() string {
	return "yyds-server"
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
