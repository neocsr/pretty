package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/codegangsta/cli"
	"github.com/neocsr/pretty/format"
)

func main() {
	n := len(os.Args)
	var url string

	if n < 2 {
		url, _ = clipboard.ReadAll()
	} else {
		url = os.Args[1]
	}

	app := cli.NewApp()
	app.Name = "pretty"
	app.Usage = "prettify urls"
	app.Action = func(c *cli.Context) {
		prettified, _ := format.Prettify(url)
		fmt.Println(prettified)
	}

	app.Run(os.Args)
}
