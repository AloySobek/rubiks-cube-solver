package solver

import (
	"fmt"
	"github.com/AloySobek/Rubik/cube"
	"strings"
)

func bfs(root *cube.Cube, g map[string]func(*cube.Cube) *cube.Cube, table map[uint64]int, i func(c *cube.Cube) uint64) {
	queue := make([]node, 0)

	queue = append(queue, node{root, ""})

	for c := queue[0]; len(queue) > 0; {
		c, queue = queue[0], queue[1:]

		for k, v := range g {
			cc := node{v(cube.Create(c.c)), c.m + k + " "}

			index := i(cc.c)

			if _, ok := table[index]; !ok {
				table[index] = len(strings.Split(strings.TrimSpace(cc.m), " "))

				queue = append(queue, cc)
			}
		}

	}
}

func generate(g map[string]func(*cube.Cube) *cube.Cube, f func(c *cube.Cube) uint64) map[uint64]int {
	table := make(map[uint64]int, 0)

	c := cube.Create(nil)

	bfs(c, g, table, f)

	return table
}

func PatternDatabase() Database {
	d := Database{
		tables: make([]map[uint64]int, 4),
		goals:  make([]map[uint64]bool, 4),
	}

	for i := 0; i < 4; i += 1 {
		data, err := readDataFromFile(fmt.Sprint("G", i, ".tab"))

		if err == nil {
			d.tables[i], err = bytesToMap(data)
		} else {
			d.tables[i] = generate(cube.GS[i], indices[i])

			data, err = mapToBytes(d.tables[i])

			if err == nil {
				writeDataToFile(data, fmt.Sprint("G", i, ".tab"))
			}
		}

		d.goals[i] = map[uint64]bool{indices[i](cube.Create(nil)): true}
	}

	return d
}
