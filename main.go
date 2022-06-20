package main

import (
	"bufio"
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

	fmt.Printf("Initial mix sequence: %s\n\n", sequence)

	c := cube.ApplyMoves(cube.Create(), strings.Split(sequence, " "), nil)

	fmt.Printf("Mixed cube:\n\n")
	cube.Print(c)

	start := time.Now()
	solution, c := algorithm.Solve(c)
	elapsed := time.Since(start)

	fmt.Printf("\nSolution sequence: %s\n\n", solution)
	fmt.Printf("Solution time: %d\n\n", elapsed)

	fmt.Printf("Would you like to see cube is solved move by move? (y/n)\n")

	reader := bufio.NewReader(os.Stdin)

	if text, error := reader.ReadString('\n'); error != nil {
		return error
	} else if strings.Replace(text, "\n", "", -1) == "y" {
		cube.ApplyMoves(c, strings.Split(strings.TrimSpace(solution), " "), nil)
	}

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
