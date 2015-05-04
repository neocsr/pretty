package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/neocsr/pretty/format"
)

func main() {
	app := cli.NewApp()
	app.Name = "pretty"
	app.Usage = "prettify urls"
	app.Action = func(c *cli.Context) {
		s, _ := format.Prettify(os.Args[1])
		fmt.Println(s)
	}

	app.Run(os.Args)
}
