package main

import (
	"fmt"
	"os"

	"github.com/gjtorikian/concentrate/concentrate"
)

func main() {
	app := concentrate.NewApp()

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
