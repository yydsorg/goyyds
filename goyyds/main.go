package main

import (
	"github.com/goyyds/goyyds/v1"
	"log"
)

func main() {
	service := goyyds.NewService()
	log.Println(service.Name())
	log.Println(service.Version())
	err := service.Run()
	if err != nil {
		log.Println(err)
	}
}
