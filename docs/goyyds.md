# v1 

# demo

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