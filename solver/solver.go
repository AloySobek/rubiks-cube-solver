package solver

import (
	"github.com/AloySobek/Rubik/model"
	"math"
)

func Solve(c *model.Cube, d Database) *model.Cube {
	return (IDAStar(
		IDAStar(
			IDAStar(
				IDAStar(
					&node{c, ""},
					model.G0,
					func(n *node) int { return 0 },
					func(n *node) bool { return true },
				),
				model.G1,
				func(n *node) int { return 0 },
				func(n *node) bool { return true },
			),
			model.G2,
			func(n *node) int { return 0 },
			func(n *node) bool { return true },
		),
		model.G3,
		func(n *node) int { return 0 },
		func(n *node) bool { return true },
	)).c
}

func IDAStar(
	root *node,
	g map[string]func(*model.Cube) *model.Cube,
	h func(*node) int,
	s func(*node) bool,
) *node {
	bound := h(root)
	path := []*node{root}

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
	path *[]*node,
	cost int,
	g map[string]func(*model.Cube) *model.Cube,
	h func(*node) int,
	s func(*node) bool,
	bound int,
) int {
	n := (*path)[len(*path)-1]

	f := cost + h(n)

	if f > bound {
		return f
	}

	if s(n) {
		return 0
	}

	min := math.MaxInt

	for k, v := range g {
		*path = append(*path, &node{v(model.Create(n.c)), n.m + k + " "})

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
