package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/AloySobek/Rubik/algorithm"
	"github.com/AloySobek/Rubik/cube"
	"github.com/AloySobek/Rubik/graph"
)

func app(ctx *cli.Context) error {
	sequence := strings.TrimSpace(ctx.String("s"))

	if sequence == "random" || sequence == "" {
		sequence = cube.GetRandomMixSequence()
	} else {
		sequence = regexp.MustCompile("\\s+").ReplaceAllString(sequence, " ")
	}

	fmt.Println("Initial mix sequence:", sequence)
	fmt.Println()

	c := cube.Mix(cube.Create(), strings.Split(sequence, " "))

	fmt.Println("Mixed cube before solving:")
	fmt.Println()
	cube.Print(c)
	fmt.Println()

	start := time.Now()
	solution, c := algorithm.Solve(c)
	elapsed := time.Since(start)

	fmt.Println("Solved cube:")
	fmt.Println()
	cube.Print(c)
	fmt.Println()

	fmt.Println("Solution sequence:", solution)
	fmt.Println()
	fmt.Println("Algorithm time:", elapsed)

	return nil
}

func main() {
	// app := &cli.App{
	// 	Name:  "Rubik",
	// 	Usage: "Rubik's cube solver",
	// 	Flags: []cli.Flag{
	// 		&cli.StringFlag{
	// 			Name:  "s",
	// 			Value: "random",
	// 			Usage: "Space separated sequence of moves to mix the cube",
	// 		},
	// 	},
	// 	Action: app,
	// }

	// err := app.Run(os.Args)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
	// 	panic(err)
	// }
	// defer sdl.Quit()

	// window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
	// 	800, 600, sdl.WINDOW_SHOWN)
	// if err != nil {
	// 	panic(err)
	// }
	// defer window.Destroy()

	// surface, err := window.GetSurface()
	// if err != nil {
	// 	panic(err)
	// }
	// surface.FillRect(nil, 0)

	// rect := sdl.Rect{0, 0, 200, 200}
	// surface.FillRect(&rect, 0xffff0000)
	// window.UpdateSurface()

	// running := true
	// for running {
	// 	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	// 		switch event.(type) {
	// 		case *sdl.QuitEvent:
	// 			println("Quit")
	// 			running = false
	// 			break
	// 		}
	// 	}
	// }

	origin := graph.PremadeGraphOne()

	goal := algorithm.Dijkstra(origin)

	if goal == nil {
		fmt.Println("Fuck! Didn't found goal node")
	} else {
		fmt.Println(goal.Data.Path)
	}

	for goal.Data.Label != graph.ORIGIN {
		fmt.Println(goal.Data.Distance)

		goal = goal.Data.Path
	}

	// graph.Print(goal)
}
