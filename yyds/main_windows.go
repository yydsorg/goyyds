package main

//go:generate .\scripts\init.bat

import (
	"github.com/yydsorg/yyds/cmd"
	_ "github.com/yydsorg/yyds/cmd/cli"
)

func main() {
	cmd.Run()
}
