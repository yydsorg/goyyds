package homedir

import (
	"github.com/mitchellh/go-homedir"
	"log"
)

var BasePath = ""

func init() {
	dir, err := homedir.Dir()
	if err != nil {
		log.Println(err)
	}
	BasePath = dir
}
