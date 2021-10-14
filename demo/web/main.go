package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yydsorg/demo/web/handler"
	"github.com/yydsorg/goyyds"
	"github.com/yydsorg/goyyds/web"
	"log"
)

func main() {
	//fs, _ := nfs.NewFS(env.GetPublicPath())
	//log.Println(fs)

	a := goyyds.NewYYDS(goyyds.Name("yyds-web"))
	err := a.Web().Init(web.Addr(":9999"), web.Handlers(group(&web.Engine().RouterGroup)))
	if err != nil {
		log.Println(err)
	}
	log.Println(a.Web().String())
	if err = a.Run(); err != nil {
		log.Println(err)
	}

}

func group(e *gin.RouterGroup) *gin.RouterGroup {
	e.POST("/info", handler.GetInfo)
	r := e.Group("/aaa")
	r.GET("/info", handler.GetInfo)
	return e
}
