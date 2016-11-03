package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Name    string `short:"n" long:"name" default:"World" description:"A name to say hello to."`
	Chinese bool   `short:"c" long:"chinese" description:"Use Chinese language"`
}

func main() {
	flags.Parse(&opts)
	if opts.Chinese == true {
		fmt.Printf("你好 %s!\n", opts.Name)
	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}
