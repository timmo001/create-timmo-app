package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	fmt.Print(CTATitle)

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v", "V"},
		Usage:   "Show the version of the app",
	}

	app := &cli.App{
		Name:    "Create Timmo App",
		Usage:   "Create an application",
		Version: "0.0.0",
		Action:  run,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(cCtx *cli.Context) error {
	Prompt()
	return nil
}
