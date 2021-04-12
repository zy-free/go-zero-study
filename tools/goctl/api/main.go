package main

import (
	"fmt"
	"os"

	"go-zero-study/core/logx"
	"go-zero-study/tools/goctl/api/parser"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}

	p, err := parser.NewParser(os.Args[1])
	logx.Must(err)
	api, err := p.Parse()
	logx.Must(err)
	fmt.Println(api)
}
