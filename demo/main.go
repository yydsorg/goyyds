package main

import (
	"github.com/goyyds/goyyds"
	"log"
)

func main() {

	service := goyyds.NewService(goyyds.Name("yyds"))
	service.Init(goyyds.BeforeStart(beforeStart), goyyds.AfterStart(afterStart), goyyds.BeforeStop(beforeStop), goyyds.AfterStop(afterStop))
	// name
	log.Println(service.Name())
	// version
	log.Println(service.Version())

	err := service.Run()
	if err != nil {
		log.Println(err)
	}
}

var (
	beforeStart = func() error {
		log.Println("beforeStart")
		return nil
	}
	afterStart = func() error {
		log.Println("afterStart")
		return nil
	}
	beforeStop = func() error {
		log.Println("beforeStop")
		return nil
	}
	afterStop = func() error {
		log.Println("afterStop")
		return nil
	}
)
