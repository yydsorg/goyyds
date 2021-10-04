package main

import (
	"github.com/goyyds/goyyds/v1"
	"log"
)

func main() {

	service := goyyds.NewService()
	service.Init(func(o *goyyds.Options) {
		o.BeforeStart = beforeStart
	}, func(o *goyyds.Options) {
		o.BeforeStop = beforeStop
	}, func(o *goyyds.Options) {
		o.AfterStart = afterStart
	}, func(o *goyyds.Options) {
		o.AfterStop = afterStop
	})
	log.Println(service.Name())
	log.Println(service.Version())
	err := service.Run()
	if err != nil {
		log.Println(err)
	}
}

var (
	beforeStart = []func() error{func() error {
		log.Println("beforeStart")
		return nil
	}}
	afterStart = []func() error{func() error {
		log.Println("afterStart")
		return nil
	}}
	beforeStop = []func() error{func() error {
		log.Println("beforeStop")
		return nil
	}}
	afterStop = []func() error{func() error {
		log.Println("afterStop")
		return nil
	}}
)
