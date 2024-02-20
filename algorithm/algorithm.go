package algorithm

import (
	"fmt"
	"github.com/AloySobek/Rubik/cube"
	"math"
	"strings"
)

type Node struct {
	c *cube.Cube
	p string
}

func heuristic(n *Node, tables *Tables) int {
	index := cube.GetEdgeOrientations(n.c)

	if index == 0 {
		return 0
	}

	return len(strings.Split(tables.G0[index], " "))
}

func Solve(c *cube.Cube, tables *Tables) {
	root := &Node{
		c,
		"",
	}

	solved := IDAStar(
		root,
		cube.G0,
		func(n *Node) int { return heuristic(n, tables) },
		func(n *Node) bool { return cube.G0Condition(n.c) },
	)

	cube.Print(solved.c)
	fmt.Printf("Path: %s : %d : %t\n", solved.p, heuristic(solved, tables), cube.G0Condition(solved.c))
}

func IDAStar(
	root *Node,
	g map[string]func(*cube.Cube) *cube.Cube, // Possible moves
	h func(*Node) int, // Heuristic function
	s func(*Node) bool, // Is solved check
) *Node {
	bound := h(root)
	path := []*Node{root}

	for {
		t := search(&path, 0, g, h, s, bound)

		if t == 0 {
			return path[len(path)-1]
		} else if t == math.MaxInt {
			return nil
		}

		bound = t
	}
}

func search(
	path *[]*Node,
	cost int,
	g map[string]func(*cube.Cube) *cube.Cube,
	h func(*Node) int,
	s func(*Node) bool,
	bound int,
) int {
	node := (*path)[len(*path)-1]

	f := cost + h(node)

	if f > bound {
		return f
	}

	if s(node) {
		return 0
	}

	min := math.MaxInt

	for k, v := range g {
		*path = append(*path, &Node{v(cube.Copy(node.c)), node.p + k + " "})

		t := search(path, cost+1, g, h, s, bound)

		if t == 0 {
			return 0
		}

		if t < min {
			min = t
		}

		*path = (*path)[:len(*path)-1]
	}

	return min
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
