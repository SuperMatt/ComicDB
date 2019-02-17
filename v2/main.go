package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/supermatt/comicdb/v2/cli"
	"github.com/supermatt/comicdb/v2/web"
)

func main() {

	fset := flag.NewFlagSet("flags", flag.ExitOnError)

	fset.Usage = func() {
		fmt.Println("lksjdlfkjs")
		fset.PrintDefaults()
	}

	fset.Parse(os.Args)

	subcommand := fset.Args()[1]

	if subcommand == "web" {
		web.Start(fset.Args()[2:])
	} else if subcommand == "cli" {
		cli.Start(fset.Args()[2:])
	} else {
		fset.Usage()
	}
}
