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
	if c.Args().Len() < 1 {
		log.Fatal("Not enough arguments!")
	}

	mixSequence := c.Args().Get(0)

	for _, move := range strings.Split(mixSequence, " ") {
		fmt.Println("Move:", move)
	}

	cu := cube.Create()

	cube.Print(cu)

	fmt.Printf("\n\n\n")

	start := time.Now()

	for i := 0; i < 1000000000; i++ {
		cube.Front(cu, true, true)
	}

	elapsed := time.Since(start)

	cube.Print(cu)

	fmt.Printf("Time to make move: %s\n", elapsed)

	return nil
}

func main() {
	app := &cli.App{
		Name:   "Rubik",
		Usage:  "Rubik's cube solver",
		Action: solve,
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
