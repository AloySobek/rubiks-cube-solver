package solver

import (
	"github.com/AloySobek/Rubik/cube"
	"math"
)

func Solve(c *cube.Cube, d Database) (*cube.Cube, string) {
	s := node{c, ""}

	for i := 0; i < 4; i += 1 {
		s = idaStar(s, cube.GS[i], func(n node) int {
			return d.Tables[i][indices[i](n.c)]
		}, func(n node) bool {
			if _, ok := d.Goals[i][indices[i](n.c)]; ok {
				return true
			}

			return false
		})
	}

	return s.c, s.m
}

func idaStar(root node, g map[string]func(*cube.Cube) *cube.Cube, h func(node) int, i func(node) bool) node {
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
	g map[string]func(*cube.Cube) *cube.Cube,
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
		nn := node{v(cube.Create(n.c)), n.m + k + " "}

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
