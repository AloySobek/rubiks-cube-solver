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
				cube.RotateS0(c)
			}
		case "U2":
			{
				cube.RotateS0Twice(c)
			}
		case "U'":
			{
				cube.RotateS0Thrice(c)
			}

		case "D":
			{
				cube.RotateS5(c)
			}
		case "D2":
			{
				cube.RotateS5Twice(c)
			}
		case "D'":
			{
				cube.RotateS5Thrice(c)
			}
		case "R":
			{
				cube.RotateS3(c)
			}
		case "R2":
			{
				cube.RotateS3Twice(c)
			}
		case "R'":
			{
				cube.RotateS3Thrice(c)
			}
		case "L":
			{
				cube.RotateS1(c)
			}
		case "L2":
			{
				cube.RotateS1Twice(c)
			}
		case "L'":
			{
				cube.RotateS1Thrice(c)
			}
		case "F":
			{
				cube.RotateS2(c)
			}
		case "F2":
			{
				cube.RotateS2Twice(c)
			}
		case "F'":
			{
				cube.RotateS2Thrice(c)
			}
		case "B":
			{
				cube.RotateS4(c)
			}
		case "B2":
			{
				cube.RotateS4Twice(c)
			}
		case "B'":
			{
				cube.RotateS4Thrice(c)
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
	algorithm.GenerateG0Table("./G0.txt")

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

// Interactive mode
// c = cube.Create()

// for reader := bufio.NewReader(os.Stdin); ; {
// 	fmt.Printf("\033[0;0H")

// 	cube.Print(c)

// 	fmt.Print("Enter rotation: ")

// 	input, err := reader.ReadString('\n')

// 	if err != nil {
// 		fmt.Printf("Error occurred while reading input: %s", err)
// 	}

// 	input = input[:len(input)-1]

// 	if input == "quit" {
// 		break
// 	}

// 	switch input {
// 	case "U":
// 		{
// 			cube.Up(c, false, false)
// 		}
// 	case "U'":
// 		{
// 			cube.Up(c, true, false)
// 		}
// 	case "U2":
// 		{
// 			cube.Up(c, false, true)
// 		}
// 	case "D":
// 		{
// 			cube.Down(c, false, false)
// 		}
// 	case "D'":
// 		{
// 			cube.Down(c, true, false)
// 		}
// 	case "D2":
// 		{
// 			cube.Down(c, false, true)
// 		}
// 	case "R":
// 		{
// 			cube.Right(c, false, false)
// 		}
// 	case "R'":
// 		{
// 			cube.Right(c, true, false)
// 		}
// 	case "R2":
// 		{
// 			cube.Right(c, false, true)
// 		}
// 	case "L":
// 		{
// 			cube.Left(c, false, false)
// 		}
// 	case "L'":
// 		{
// 			cube.Left(c, true, false)
// 		}
// 	case "L2":
// 		{
// 			cube.Left(c, false, true)
// 		}
// 	case "F":
// 		{
// 			cube.Front(c, false, false)
// 		}
// 	case "F'":
// 		{
// 			cube.Front(c, true, false)
// 		}
// 	case "F2":
// 		{
// 			cube.Front(c, false, true)
// 		}
// 	case "B":
// 		{
// 			cube.Back(c, false, false)
// 		}
// 	case "B'":
// 		{
// 			cube.Back(c, true, false)
// 		}
// 	case "B2":
// 		{
// 			cube.Back(c, false, true)
// 		}
// 	default:
// 		{
// 			fmt.Printf("Unknown command\n")
// 		}
// 	}

// }
