package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func solve(c *cli.Context) error {
	mixSequence := c.Args().Get(0)

	for _, move := range strings.Split(mixSequence, " ") {
		fmt.Println("Move:", move)
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:   "Rubik",
		Usage:  "Rubik's cube solver",
		Action: solve,
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
