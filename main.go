package main

import (
	"bufio"
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

func GetRandomMixSequence(n int) (sequence string) {
	possibleMoves := strings.Split("F R L B U D F' R' L' B' U' D' F2 R2 L2 B2 U2 D2", " ")

	for i := 0; i < n; i += 1 {

		sequence += possibleMoves[rand.Intn(len(possibleMoves))]

		if i+1 != n {
			sequence += " "
		}
	}

	return
}

func solve(ctx *cli.Context) error {
	sequence := ctx.Args().Get(0)

	if sequence == "" {
		sequence = GetRandomMixSequence(ctx.Int("n"))
	} else {
		sequence = regexp.MustCompile("\\s+").ReplaceAllString(sequence, " ")
	}

	fmt.Printf("Initial mix sequence: %s\n\n", sequence)

	c := cube.Create(nil)

	cube.ApplyMoves(c, strings.Split(sequence, " "))

	d := solver.PatternDatabase()

	fmt.Printf("%d : %d : %d : %d\n", len(d.Tables[0]), len(d.Tables[1]), len(d.Tables[2]), len(d.Tables[3]))

	start := time.Now()

	_, s := solver.Solve(c, d)

	elapsed := time.Since(start)

	fmt.Println(s)

	fmt.Printf("Solution time: %d ms\n", elapsed.Milliseconds())

	return nil
}

func interactive(ctx *cli.Context) error {
	c := cube.Create(nil)

	for reader := bufio.NewReader(os.Stdin); ; {
		render.Render(c)

		fmt.Print("Enter rotation: ")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("Error occurred while reading input: %s", err)
		}

		input = input[:len(input)-1]

		if input == "quit" {
			break
		}

		if v, ok := cube.G0[input]; ok {
			v(c)
		}
	}

	return nil
}

func application(ctx *cli.Context) error {
	if ctx.Bool("i") {
		return interactive(ctx)
	}

	return solve(ctx)
}

func main() {
	app := &cli.App{
		Name:  "Rubik",
		Usage: "Rubik's cube solver",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "i",
				Value: false,
				Usage: "Interactive mode",
			},
			&cli.IntFlag{
				Name:  "n",
				Value: 100,
				Usage: "n random generated moves",
			},
		},
		Action: application,
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
