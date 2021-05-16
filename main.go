package main

import (
	"flag"
	"fmt"
	"xwdr/gen-srv-tpl/gen"
	"os"
)

var usage = func() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

	flag.PrintDefaults()
}

func main() {
	service := flag.String("service", "", "服务路径, e.g. github.com/xwdr/example")
	flag.Parse()
	if len(*service) == 0 {
		usage()
		return
	}
	gen.GenerateIndependent(*service)
}
