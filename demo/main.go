package main

import (
	"github.com/goyyds/goyyds/v1"
	"log"
)

func main() {
	s := goyyds.NewService()
	err := s.Run()
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
