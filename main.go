package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/AloySobek/Rubik/algorithm"
	"github.com/AloySobek/Rubik/cube"
)

func app(ctx *cli.Context) error {
	sequence := strings.TrimSpace(ctx.String("s"))

	if sequence == "random" || sequence == "" {
		sequence = cube.GetRandomMixSequence()
	} else {
		sequence = regexp.MustCompile("\\s+").ReplaceAllString(sequence, " ")
	}

	fmt.Println("Initial mix sequence:", sequence)
	fmt.Println()

	c := cube.Mix(cube.Create(), strings.Split(sequence, " "))

	fmt.Println("Mixed cube before solving:")
	fmt.Println()
	cube.Print(c)
	fmt.Println()

	start := time.Now()
	solution, c := algorithm.Solve(c)
	elapsed := time.Since(start)

	fmt.Println("Solved cube:")
	fmt.Println()
	cube.Print(c)
	fmt.Println()

	fmt.Println("Solution sequence:", solution)
	fmt.Println()
	fmt.Println("Algorithm time:", elapsed)

	return nil
}

func main() {
	app := &cli.App{
		Name:  "Rubik",
		Usage: "Rubik's cube solver",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "s",
				Value: "random",
				Usage: "Space separated sequence of moves to mix the cube",
			},
		},
		Action: app,
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
