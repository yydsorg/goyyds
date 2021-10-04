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

func Version(v string) Option {
	return func(o *Options) {
		o.Version = v
	}
}
