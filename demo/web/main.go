package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yydsorg/demo/web/handler"
	"github.com/yydsorg/goyyds"
	"github.com/yydsorg/goyyds/web"
	"log"
)

func main() {

	w := goyyds.NewService(goyyds.Name("yyds-web")).Web()
	err := w.Init(web.Addr(":9999"), web.Handlers(group(&web.Engine().RouterGroup)))
	if err != nil {
		log.Println(err)
	}
	log.Println(w.String())
	err = w.Run()
	if err != nil {
		log.Println(err)
	}

}

func group(e *gin.RouterGroup) *gin.RouterGroup {
	e.GET("/get", handler.GetInfo)
	r := e.Group("/aaa")
	r.GET("/info", handler.GetInfo)
	return e
}
