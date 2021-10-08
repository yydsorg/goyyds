package nfs

import (
	"io/ioutil"
	"log"
	"os"
)

func CopyNew() error {
	if err := ioutil.WriteFile("a.txt", []byte("hello world"), os.ModePerm); err != nil {
		log.Println(err)
		return err
	}

	if err := os.Symlink("a.txt", "b.txt"); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
