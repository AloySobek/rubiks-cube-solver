package algorithm

import (
	"github.com/AloySobek/Rubik/cube"
	"strings"
)

type Node struct {
	c *cube.Cube
	p *[]string
}

func heuristic(n Node, tables *Tables) uint64 {
	var result uint64 = 0

	p := strings.Split(tables.G0[cube.GetEdgeOrientations(n.c)], " ")

	for i, v := range p {
		if i < len(*n.p) && (*n.p)[i] == v {
			continue
		}

		result += 1
	}

	return result
}

func Solve(c *cube.Cube, tables *Tables) {
	// solution := IDAStar(c, func(c *cube.Cube) uint64 { return heuristic(c, tables) })

	// solution := append(make([]string, 0, 100), cube.ApplyMoves(c, IDBFS(c, cube.G0, cube.G0Condition))...)

	//fmt.Println(solution)
	//cube.Print(c)

	// solution = append(solution, cube.ApplyMoves(c, IDBFS(c, cube.G1, cube.G1Condition))...)

	// fmt.Println(solution)
	// cube.Print(c)

	// solution = append(solution, cube.ApplyMoves(c, IDBFS(c, cube.G2, cube.G2Condition))...)

	// fmt.Println(solution)
	// cube.Print(c)

	// solution = append(solution, cube.ApplyMoves(c, IDBFS(c, cube.G3, cube.G3Condition))...)

	// fmt.Println(solution)
	// cube.Print(c)
}

func IDAStar(
	c *cube.Cube, // Cube
	g map[string]func(*cube.Cube) *cube.Cube, // Possible moves
	h func(*cube.Cube) uint64, // Heuristic function
	s func(*cube.Cube) bool, // Is solved check
) (solution []string) {
	bound := h(c)

	for search(c, g, h, s, &solution) {

	}

	return
}

func search(
	c *cube.Cube,
	g map[string]func(*cube.Cube) *cube.Cube,
	h func(*cube.Cube) uint64,
	s func(*cube.Cube) bool,
	solution *[]string,
) bool {
	return false
}

func IDDFS(c *cube.Cube, g map[string]func(*cube.Cube) *cube.Cube, s func(*cube.Cube) bool) []string {
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
