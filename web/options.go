package web

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Options struct {
	Addr    string
	Engine  *gin.Engine
	Context context.Context
}

func newOptions(options ...Option) Options {

	opts := Options{
		Engine:  gin.Default(),
		Context: context.Background(),
	}
	for _, o := range options {
		o(&opts)
	}

	return opts
}
