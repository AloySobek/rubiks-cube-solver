package solver

import (
	"github.com/AloySobek/Rubik/model"
	"math"
)

func Solve(c *model.Cube, d Database) (*model.Cube, string) {
	s :=
		IDAStar(
			IDAStar(
				node{c, ""},
				model.G0,
				func(n node) int {
					return d.G0[G0Index(n.c)]
				}, func(n node) bool {
					if _, ok := d.G0Goal[G0Index(n.c)]; !ok {
						return false
					}

					return true
				}),
			model.G1,
			func(n node) int {
				return d.G1[G1Index(n.c)]
			}, func(n node) bool {
				if _, ok := d.G1Goal[G1Index(n.c)]; !ok {
					return false
				}

				return true
			})

	// IDDFS(s.c, model.G2, func(c *model.Cube) bool {
	// 	return true
	// })

	return s.c, s.m
}

func IDAStar(root node, g map[string]func(*model.Cube) *model.Cube, h func(node) int, i func(node) bool) node {
	bound := h(root)
	path := []node{root}

	for {
		t := search(&path, 0, g, h, i, bound)

		if t == 0 {
			return path[len(path)-1]
		}

		bound = t
	}
}

func search(
	path *[]node,
	cost int,
	g map[string]func(*model.Cube) *model.Cube,
	h func(node) int,
	i func(node) bool,
	bound int,
) int {
	n := (*path)[len(*path)-1]

	f := cost + h(n)

	if f > bound {
		return f
	}

	if i(n) {
		return 0
	}

	min := math.MaxInt

	for k, v := range g {
		nn := node{v(model.Create(n.c)), n.m + k + " "}

		if alreadyInPath(nn, *path) {
			continue
		}

		*path = append(*path, nn)

		t := search(path, cost+1, g, h, i, bound)

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

func IDDFS(c *model.Cube, g map[string]func(*model.Cube) *model.Cube, s func(*model.Cube) bool) []string {
	solution := make([]string, 0, 64)

	for i := 0; !DLS(c, g, s, i, &solution); i += 1 {
	}

	return solution
}

func DLS(c *model.Cube, g map[string]func(*model.Cube) *model.Cube, s func(*model.Cube) bool, depth int, solution *[]string) bool {
	if depth <= 0 {
		return s(c)
	}

	for k, v := range g {
		*solution = append(*solution, k)

		if DLS(v(model.Create(c)), g, s, depth-1, solution) {
			return true
		}

		*solution = (*solution)[:len(*solution)-1]
	}

	return false
}

func alreadyInPath(n node, path []node) bool {
	for _, v := range path {
		same := true

		for i := 0; i < 12; i += 1 {
			if n.c.EO[i] != v.c.EO[i] || n.c.EP[i] != v.c.EP[i] {
				same = false

				break
			}
		}

		if !same {
			continue
		}

		for i := 0; i < 8; i += 1 {
			if n.c.CO[i] != v.c.CO[i] || n.c.CP[i] != v.c.CP[i] {
				same = false

				break
			}
		}

		if same {
			return true
		}
	}

	return false
}
