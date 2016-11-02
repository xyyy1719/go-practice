package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to.")
var chinese bool

func init() {
	flag.BoolVar(&chinese, "chinese", false, "Use Chinese language.")
	flag.BoolVar(&chinese, "c", false, "Use Chinese language.")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if chinese == true {
		fmt.Printf("你好 %s!\n", *name)
	} else {
		fmt.Printf("Hello %s", *name)
	}
}
