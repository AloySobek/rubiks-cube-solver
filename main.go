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

	"github.com/AloySobek/Rubik/algorithm"
	"github.com/AloySobek/Rubik/cube"
	"github.com/AloySobek/Rubik/database"
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

func app(ctx *cli.Context) error {
	sequence := ctx.Args().Get(0)

	if sequence == "" {
		sequence = GetRandomMixSequence(ctx.Int("n"))
	} else {
		sequence = regexp.MustCompile("\\s+").ReplaceAllString(sequence, " ")
	}

	fmt.Printf("Initial mix sequence: %s\n\n", sequence)

	c := cube.Create()

	cube.ApplyMoves(c, strings.Split(sequence, " "))

	fmt.Printf("Mixed cube:\n")

	cube.Print(c)

	start := time.Now()

	tables := algorithm.Tables{
		G0: algorithm.ReadG0Table(algorithm.ReadDataFromFile("G0.table")),
	}

	algorithm.Solve(c, &tables)

	elapsed := time.Since(start)

	fmt.Printf("Solution time: %f\n", elapsed.Seconds())

	return nil
}

func interactive(ctx *cli.Context) error {
	c := model.Create(nil)

	for reader := bufio.NewReader(os.Stdin); ; {
		fmt.Printf("\033[2J")

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

		if v, ok := solver.G0[input]; ok {
			v(c)
		}
	}

	return nil
}

func gen(ctx *cli.Context) error {
	i := 0

	for _, v := range database.GenerateG0() {
		if v != "" {
			fmt.Printf("%d: %s\n", i, v)

			i += 1
		}

	}

	// algorithm.WriteDataToFile(algorithm.GenerateG0Table(), "G0.table")

	// table := algorithm.ReadG0Table(algorithm.ReadDataFromFile("G0.table"))

	return nil
}

func main() {
	app := &cli.App{
		Name:  "Rubik",
		Usage: "Rubik's cube solver",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "n",
				Value: 1,
				Usage: "n random generated moves",
			},
		},
		Action: gen,
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
