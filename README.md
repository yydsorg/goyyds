# goyyds  [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0)  
[goyyds](https://github.com/yydsorg/goyyds)

# Overview
a simple rewrite of go-micro. learn from it.

[micro](https://github.com/asim/go-micro)

# Structure

- goyyds.go - goyyds
- yyds/ - cli  is a simple, fast, and fun package for building command line apps in Go. 
- plugins/  - plugins in goyyds
- demo/ - example
- docs/ - document for goyyds use


# Demo

- [demo](./docs/goyyds.md)
- [cli](./docs/cli.md)

```go
package main

import (
	"github.com/yydsorg/goyyds"
	"log"
)

func main()  {
	s := goyyds.NewService()
	err := s.Run()
	if err != nil {
		log.Println(err)
	}

}
```