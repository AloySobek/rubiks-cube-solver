package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/AloySobek/Rubik/cube"
	"github.com/AloySobek/Rubik/solver"
	"github.com/urfave/cli/v2"
)

func getRandomMove(n int) (move string) {
	possibleMoves := strings.Split("F R L B U D F' R' L' B' U' D' F2 R2 L2 B2 U2 D2", " ")

	for i := 0; i < n; i += 1 {

		move += possibleMoves[rand.Intn(len(possibleMoves))]

		if i+1 != n {
			move += " "
		}
	}

	return
}

func visualize(scrumbled, solved *cube.Cube, solution []string, elapsed time.Duration) {
	fmt.Printf(
		"The cube has been solved in %d ms with %d rotations (Half Turn Metric)!\n",
		elapsed.Milliseconds(),
		len(solution),
	)

	fmt.Printf("\nCube initial state:\n")
	cube.Print(scrumbled)
	fmt.Printf("Cube solved state:\n")
	cube.Print(solved)
	fmt.Printf("Solution move: %s\n", solution)
	fmt.Printf("\nPress 'return' to see each rotation in action...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	c := cube.Create(scrumbled)

	fmt.Printf("\033[2J")

	for i := 0; i < len(solution); i += 1 {
		if solution[i] == "" {
			continue
		}

		cube.G0[solution[i]](c)

		fmt.Printf("\033[H\033[1m%s\033[0m\n", solution[i])

		cube.Print(c)

		time.Sleep(time.Second)
	}
}

func app(ctx *cli.Context) error {
	move := ctx.Args().Get(0)

	if move == "" {
		move = getRandomMove(ctx.Int("r"))
	}

	scrumbled := cube.ApplyMoves(cube.Create(nil), strings.Split(move, " "))

	database := solver.PatternDatabase()

	start := time.Now()

	solved, solution := solver.Solve(cube.Create(scrumbled), database, ctx.Bool("v"))

	elapsed := time.Since(start)

	if ctx.Bool("d") {
		visualize(scrumbled, solved, strings.Split(solution, " "), elapsed)
	} else if !ctx.Bool("v") {
		fmt.Printf("%s\n", strings.Trim(solution, " "))
	}

	return nil
}

func main() {
	if err := (&cli.App{
		Name:   "Rubik",
		Usage:  "Rubik's cube solver",
		Action: app,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "d",
				Value: false,
				Usage: "Colored 2D cube visualization",
			},
			&cli.BoolFlag{
				Name:  "v",
				Value: false,
				Usage: "Detailed description of algorithm's steps",
			},
			&cli.IntFlag{
				Name:  "r",
				Value: 30,
				Usage: "Generate move randomly with specified number of rotations",
			},
		},
	}).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
