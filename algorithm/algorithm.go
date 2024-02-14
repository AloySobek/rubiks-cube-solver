package algorithm

import (
	"fmt"
	"github.com/AloySobek/Rubik/cube"
)

func Solve(c *cube.Cube) {
	solution := append(make([]string, 0, 100), cube.ApplyMoves(c, IDBFS(c, cube.G0, cube.G0Condition))...)

	fmt.Println(solution)
	cube.Print(c)

	solution = append(solution, cube.ApplyMoves(c, IDBFS(c, cube.G1, cube.G1Condition))...)

	fmt.Println(solution)
	cube.Print(c)

	solution = append(solution, cube.ApplyMoves(c, IDBFS(c, cube.G2, cube.G2Condition))...)

	fmt.Println(solution)
	cube.Print(c)

	solution = append(solution, cube.ApplyMoves(c, IDBFS(c, cube.G3, cube.G3Condition))...)

	fmt.Println(solution)
	cube.Print(c)
}

func IDBFS(c *cube.Cube, g map[string]func(*cube.Cube) *cube.Cube, s func(*cube.Cube) bool) []string {
	solution := make([]string, 0, 64)

	for i := 0; !DLS(c, g, s, i, &solution); i += 1 {
	}

	return solution
}

func DLS(c *cube.Cube, g map[string]func(*cube.Cube) *cube.Cube, s func(*cube.Cube) bool, depth int, solution *[]string) bool {
	if depth <= 0 {
		return s(c)
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
