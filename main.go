package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/AloySobek/Rubik/cube"
)

func solve(c *cli.Context) error {
	var argc int = c.Args().Len()

	if argc < 1 {
		log.Fatal("Not enough arguments")
	}

	var notationMatrix map[string]func(*cube.Cube, bool, bool) = make(map[string]func(*cube.Cube, bool, bool))

	notationMatrix["F"] = cube.Front
	notationMatrix["R"] = cube.Right
	notationMatrix["L"] = cube.Left
	notationMatrix["B"] = cube.Back
	notationMatrix["U"] = cube.Up
	notationMatrix["D"] = cube.Down

	notationMatrix["F'"] = cube.RFront
	notationMatrix["R'"] = cube.RRight
	notationMatrix["L'"] = cube.RLeft
	notationMatrix["B'"] = cube.RBack
	notationMatrix["U'"] = cube.RUp
	notationMatrix["D'"] = cube.RDown

	notationMatrix["F2"] = cube.DFront
	notationMatrix["R2"] = cube.DRight
	notationMatrix["L2"] = cube.DLeft
	notationMatrix["B2"] = cube.DBack
	notationMatrix["U2"] = cube.DUp
	notationMatrix["D2"] = cube.DDown

	notationMatrix["F2'"] = cube.RDFront
	notationMatrix["R2'"] = cube.RDRight
	notationMatrix["L2'"] = cube.RDLeft
	notationMatrix["B2'"] = cube.RDBack
	notationMatrix["U2'"] = cube.RDUp
	notationMatrix["D2'"] = cube.RDDown

	cu := cube.Create()

	mixSequenceString := c.Args().Get(0)

	start := time.Now()

	for _, move := range strings.Split(mixSequenceString, " ") {
		if value, ok := notationMatrix[strings.ToUpper(move)]; ok {
			value(cu, false, false)
		} else {
			log.Fatalf("Unsupported move: %s", move)
		}
	}

	elapsed := time.Since(start)

	if c.Bool("v") {
		fmt.Printf("Mixing cube took %s\n", elapsed)
		cube.Print(cu)
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:  "Rubik",
		Usage: "Rubik's cube solver",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "v",
				Value: false,
				Usage: "Show debug output",
			},
		},
		Action: solve,
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
