package main

import (
	"fmt"
	"os"

	"github.com/gjtorikian/concetrate/concentrate"
)

func main() {
	app := concetrate.NewApp()

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
