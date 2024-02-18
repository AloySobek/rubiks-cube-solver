package main

import (
	"bufio"
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
		sequence = cube.GetRandomMixSequence(ctx.Int("n"))
	} else {
		sequence = regexp.MustCompile("\\s+").ReplaceAllString(sequence, " ")
	}

	fmt.Printf("Initial mix sequence: %s\n\n", sequence)

	c := cube.Create()

	cube.ApplyMoves(c, strings.Split(sequence, " "))

	fmt.Printf("Mixed cube:\n")

	cube.Print(c)

	start := time.Now()

	algorithm.Solve(c)

	elapsed := time.Since(start)

	fmt.Printf("Solution time: %f\n", elapsed.Seconds())

	return nil
}

func interactive(ctx *cli.Context) error {
	c := cube.Create()

	for reader := bufio.NewReader(os.Stdin); ; {
		fmt.Printf("\033[0;0H")

		cube.Print(c)

		fmt.Print("Enter rotation: ")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("Error occurred while reading input: %s", err)
		}

		input = input[:len(input)-1]

		if input == "quit" {
			break
		}

		switch input {
		case "U":
			{
				cube.Up(c)
			}
		case "U2":
			{
				cube.Up2(c)
			}
		case "U'":
			{
				cube.UpPrime(c)
			}

		case "D":
			{
				cube.Down(c)
			}
		case "D2":
			{
				cube.Down2(c)
			}
		case "D'":
			{
				cube.DownPrime(c)
			}
		case "R":
			{
				cube.Right(c)
			}
		case "R2":
			{
				cube.Right2(c)
			}
		case "R'":
			{
				cube.RightPrime(c)
			}
		case "L":
			{
				cube.Left(c)
			}
		case "L2":
			{
				cube.Left2(c)
			}
		case "L'":
			{
				cube.LeftPrime(c)
			}
		case "F":
			{
				cube.Front(c)
			}
		case "F2":
			{
				cube.Front2(c)
			}
		case "F'":
			{
				cube.FrontPrime(c)
			}
		case "B":
			{
				cube.Back(c)
			}
		case "B2":
			{
				cube.Back2(c)
			}
		case "B'":
			{
				cube.BackPrime(c)
			}
		default:
			{
				fmt.Printf("Unknown command\n")
			}
		}

	}

	return nil
}

func gen(ctx *cli.Context) error {
	algorithm.WriteDataToFile("G0.table", algorithm.GenerateG0Table())

	table := algorithm.ReadG0Table(algorithm.ReadDataFromFile("G0.table"))

	for i, v := range table {
		fmt.Printf("%d: %s\n", i, v)
	}

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
