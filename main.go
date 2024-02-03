package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
	"unsafe"

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

	c := cube.ApplyMoves(cube.Create(), strings.Split(sequence, " "), nil)

	fmt.Printf("Size of struct: %d\n", unsafe.Sizeof(*c))

	fmt.Printf("Mixed cube:\n\n")

	cube.Print(c)

	start := time.Now()

	solution := algorithm.Solve(c)

	elapsed := time.Since(start)

	fmt.Printf("Hello: %s\n", solution)

	c = cube.ApplyMoves(c, strings.Split(solution, " "), nil)

	fmt.Printf("\nSolution sequence: %s\n\nSolution time: %f\n\nSolved cube:\n\n", solution, elapsed.Seconds())

	cube.Print(c)

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
		Action: app,
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
