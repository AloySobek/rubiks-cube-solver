package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/AloySobek/Rubik/cube"
	"github.com/AloySobek/Rubik/solver"
	"github.com/urfave/cli/v2"
)

func app(ctx *cli.Context) error {
	move := ctx.Args().Get(0)

	if move == "" {
		return cli.Exit("No move provided", 0)
	}
	_, s := solver.Solve(
		cube.ApplyMoves(
			cube.Create(nil),
			strings.Split(regexp.MustCompile("\\s+").ReplaceAllString(move, " "), " "),
		),
		solver.PatternDatabase(),
	)

	fmt.Printf("%s\n", strings.Trim(s, " "))

	return nil
}

func main() {
	if err := (&cli.App{
		Name:   "Rubik",
		Usage:  "Rubik's cube solver",
		Action: app,
	}).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
