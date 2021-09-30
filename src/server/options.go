package server

type Options struct {
	Name string
}

func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}
