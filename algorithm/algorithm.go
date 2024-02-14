package algorithm

import (
	"fmt"
	"github.com/AloySobek/Rubik/cube"
)

func Solve(c *cube.Cube) {
	solution := make([]string, 0, 100)

	IDBFS(c, cube.G0, cube.G0Condition, 7, &solution)

	g1 := cube.ApplyMoves(c, solution, nil)

	fmt.Println(solution)
	cube.Print(g1)

	IDBFS(g1, cube.G1, cube.G1Condition, 10, &solution)

	g2 := cube.ApplyMoves(c, solution, nil)

	fmt.Println(solution)
	cube.Print(g2)
	// IDBFS(c, cube.G2, cube.G2Condition, 13, &solution)
	// IDBFS(c, cube.G3, cube.G3Condition, 15, &solution)
}

func IDBFS(c *cube.Cube, g map[string]func(*cube.Cube) *cube.Cube, s func(*cube.Cube) bool, maxDepth int, solution *[]string) {
	for i := 0; i < maxDepth; i += 1 {
		if DLS(c, g, s, i, solution) {
			break
		}
	}
}

func DLS(c *cube.Cube, g map[string]func(*cube.Cube) *cube.Cube, s func(*cube.Cube) bool, depth int, solution *[]string) bool {
	if depth <= 0 {
		if s(c) {
			return true
		} else {
			return false
		}
	}

	for k, v := range g {
		*solution = append(*solution, k)

		if DLS(v(cube.Copy(c)), g, s, depth-1, solution) {
			return true
		}

		*solution = (*solution)[:len(*solution)-1]
	}

	return false
}
