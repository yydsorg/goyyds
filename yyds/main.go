package main

import (
	"github.com/goyyds/goyyds/v1/plugins/cmd"
	"log"
)

func main() {
	c := cmd.DefaultCmd
	c.Init(cmd.Name("aaaaa"))
	log.Println(c.Options().Name, c.Options().Command, c.Options().Version, c.Options().Description)
	c.Run()
}
