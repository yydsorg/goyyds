package server

type Options struct {
	Name    string
	Version string
}

func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

func newOptions(opt ...Option) Options {
	opts := Options{
		Name:    DefaultName,
		Version: DefaultVersion,
	}
	for _, o := range opt {
		o(&opts)
	}

	return opts
}
