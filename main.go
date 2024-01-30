package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/AloySobek/Rubik/algorithm"
	"github.com/AloySobek/Rubik/cube"
	"github.com/urfave/cli/v2"
)

func app(ctx *cli.Context) error {
	sequence := ctx.Args().Get(0)

	if sequence == "" {
		sequence = cube.GetRandomMixSequence()
	} else {
		sequence = regexp.MustCompile("\\s+").ReplaceAllString(sequence, " ")
	}

	fmt.Printf("Initial mix sequence: %s\n\n", sequence)

	c := cube.ApplyMoves(cube.Create(), strings.Split(sequence, " "), nil)

	fmt.Printf("Mixed cube:\n\n")
	cube.Print(c)

	start := time.Now()

	solution, c := algorithm.Solve(c)

	elapsed := time.Since(start)

	c = cube.ApplyMoves(c, solution, nil)

	fmt.Printf("\nSolution sequence: %s\n\nSolution time: %f\n\nSolved cube:\n\n", solution, elapsed.Seconds())

	cube.Print(c)

	return nil
}

func main() {
	app := &cli.App{
		Name:   "Rubik",
		Usage:  "Rubik's cube solver",
		Action: app,
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
