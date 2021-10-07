package web

import "github.com/gin-gonic/gin"

type Web interface {
	Init(...Option) error
	Options() Options
	Run() error
	String() string
}

// Option used by the Client
type Option func(*Options)

type web struct {
	opts Options
}

var (
	DefaultWeb = newWeb()
)

func newWeb(opt ...Option) Web {
	opts := newOptions(opt...)

	web := &web{opts: opts}
	return Web(web)
}

func (w *web) Init(opts ...Option) error {
	for _, o := range opts {
		o(&w.opts)
	}
	return nil
}

func (w *web) Options() Options {
	return w.opts
}

func (w *web) Run() error {

	err := w.opts.Engine.Run(w.opts.Addr)
	return err
}

func (w *web) String() string {
	return "yyds-web"
}

func Addr(a string) Option {
	return func(o *Options) {
		o.Addr = a
	}
}

func Handlers(g *gin.RouterGroup) Option {
	return func(o *Options) {
		o.RouterGroup = g
	}
}

func Engine() *gin.Engine {
	return DefaultWeb.Options().Engine
}
