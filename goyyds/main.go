package main

import (
	"github.com/goyyds/goyyds/v1"
	"log"
)

func main() {
	var options []goyyds.Option
	options = append(options, goyyds.Name("aaa"), goyyds.Version("0.0.1"))

	service := goyyds.NewService(options...)
	err := service.Run()
	if err != nil {
		log.Println(err)
	}
}
