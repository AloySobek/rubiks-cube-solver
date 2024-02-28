package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/AloySobek/Rubik/cube"
	"github.com/AloySobek/Rubik/render"
	"github.com/AloySobek/Rubik/solver"
	"github.com/urfave/cli/v2"
)

func GetRandomMove(n int) (move string) {
	possibleMoves := strings.Split("F R L B U D F' R' L' B' U' D' F2 R2 L2 B2 U2 D2", " ")

	for i := 0; i < n; i += 1 {

		move += possibleMoves[rand.Intn(len(possibleMoves))]

		if i+1 != n {
			move += " "
		}
	}

	return
}

func solve(ctx *cli.Context) error {
	move := ctx.Args().Get(0)

	if move == "" {
		move = GetRandomMove(ctx.Int("n"))
	} else {
		move = regexp.MustCompile("\\s+").ReplaceAllString(move, " ")
	}

	fmt.Printf("Scramble move: %s\n", move)

	c := cube.Create(nil)

	cube.ApplyMoves(c, strings.Split(move, " "))

	cc := cube.Create(c)

	render.Render(c)

	d := solver.PatternDatabase()

	start := time.Now()

	_, s := solver.Solve(c, d)

	elapsed := time.Since(start)

	fmt.Printf("Solution move: %s\n", s)

	cube.ApplyMoves(cc, strings.Split(strings.Trim(s, " "), " "))

	render.Render(cc)

	fmt.Printf("Solution time: %d ms\n", elapsed.Milliseconds())

	return nil
}

func main() {
	app := &cli.App{
		Name:  "Rubik",
		Usage: "Rubik's cube solver",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "n",
				Value: 100,
				Usage: "n random generated moves",
			},
		},
		Action: solve,
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
