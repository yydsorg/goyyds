# goyyds  [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0)  
[goyyds](https://github.com/goyyds/goyyds)

a simple rewrite of go-micro. learn from it.

[micro](https://github.com/asim/go-micro)

# structure

demo/
docs/
goyyds.go


# demo

[demo v1](./docs/v1.md)

```go
package main

import (
	"github.com/goyyds/goyyds/v1"
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