package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Usage = "Print hello world"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "World",
			Usage: "Who to say hello to.",
		},
		cli.BoolFlag{
			Name:  "chinese, c",
			Usage: "Use Chinese language",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		chinese := c.GlobalBool("chinese")
		if chinese == true {
			fmt.Printf("你好 %s!\n", name)
		} else {
			fmt.Printf("Hello %s!\n", name)
		}
		return nil
	}
	app.Run(os.Args)
}
