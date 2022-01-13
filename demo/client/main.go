package main

import (
	"github.com/yydsorg/goyyds"
	"log"
)

func main() {
	service := goyyds.NewYYDS(goyyds.Name("c"), goyyds.Genre("client"))

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
