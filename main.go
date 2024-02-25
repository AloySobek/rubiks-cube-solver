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

	"github.com/AloySobek/Rubik/model"
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

	c := model.Create(nil)

	model.ApplyMoves(c, strings.Split(sequence, " "))

	d := solver.DatabaseFromFile()

	start := time.Now()

	_, s := solver.Solve(c, d)

	elapsed := time.Since(start)

	fmt.Println(s)

	fmt.Printf("Solution time: %d ms\n", elapsed.Milliseconds())

	return nil
}

func interactive(ctx *cli.Context) error {
	c := model.Create(nil)

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

		if v, ok := model.G0[input]; ok {
			v(c)
		}
	}

	return nil
}

func dbGen(ctx *cli.Context) error {
	fmt.Println("Generating pattern databases...")

	if _, err := os.Stat("./assets"); os.IsNotExist(err) {
		os.Mkdir("./assets", 0777)
	}

	start := time.Now()

	solver.Save(solver.NewDatabase())

	elapsed := time.Since(start)

	fmt.Printf("Pattern databases has been successfully generated, elapsed time(in seconds): %f\n", elapsed.Seconds())

	d := solver.DatabaseFromFile()

	fmt.Printf("%d : %d : %d : %d\n", len(d.G0), len(d.G1), len(d.G2), len(d.G3))

	return nil
}

func application(ctx *cli.Context) error {
	if ctx.Bool("db") {
		return dbGen(ctx)
	} else if ctx.Bool("i") {
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
				Name:  "db",
				Value: false,
				Usage: "Generate pattern databases only",
			},
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
