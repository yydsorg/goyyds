package nfs

import (
	"fmt"
	"github.com/go-git/go-billy/v5/memfs"
	nfs2 "github.com/willscott/go-nfs"
	nfshelper "github.com/willscott/go-nfs/helpers"
	"log"
	"net"
)

func nfs() {
	listener, err := net.Listen("tcp", ":0")
	panicOnErr(err, "starting TCP listener")
	fmt.Printf("Server running at %s\n", listener.Addr())
	mem := memfs.New()
	f, err := mem.Create("hello.txt")
	panicOnErr(err, "creating file")
	_, err = f.Write([]byte("hello world"))
	panicOnErr(err, "writing data")
	f.Close()
	handler := nfshelper.NewNullAuthHandler(mem)
	cacheHelper := nfshelper.NewCachingHandler(handler, 1)
	panicOnErr(nfs2.Serve(listener, cacheHelper), "serving nfs")
}

func panicOnErr(err error, desc ...interface{}) {
	if err == nil {
		return
	}
	log.Println(desc...)
	log.Panicln(err)
}
