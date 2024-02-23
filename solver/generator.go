package solver

import (
	"github.com/AloySobek/Rubik/model"
	"strings"
)

func bfs(root *model.Cube, g map[string]func(*model.Cube) *model.Cube, table map[uint64]int, i func(c *model.Cube) uint64) {
	queue := make([]node, 0)

	queue = append(queue, node{root, ""})

	for c := queue[0]; len(queue) > 0; {
		c, queue = queue[0], queue[1:]

		for k, v := range g {
			cc := node{v(model.Create(c.c)), c.m + k + " "}

			index := i(cc.c)

			if _, ok := table[index]; !ok {
				table[index] = len(strings.Split(strings.TrimSpace(cc.m), " "))

				queue = append(queue, cc)
			}
		}

	}
}

func generate(g map[string]func(*model.Cube) *model.Cube, f func(c *model.Cube) uint64) map[uint64]int {
	table := make(map[uint64]int, 0)

	c := model.Create(nil)

	bfs(c, g, table, f)

	return table
}

func DatabaseFromFile() Database {
	d := Database{
		G0: bytesToMap(readDataFromFile("assets/G0.table")),
		G1: bytesToMap(readDataFromFile("assets/G1.table")),
		G2: bytesToMap(readDataFromFile("assets/G2.table")),
		G3: bytesToMap(readDataFromFile("assets/G3.table")),

		G0Goal: map[uint64]bool{},
		G1Goal: map[uint64]bool{},
		G2Goal: map[uint64]bool{},
		G3Goal: map[uint64]bool{},
	}

	solved := model.Create(nil)

	d.G0Goal[G0Index(solved)] = true

	for i := 0; i < 4; i += 1 {
		d.G1Goal[G1Index(solved)] = true

		solved = model.L(solved)
		solved = model.L(solved)
		solved = model.L(solved)
		solved = model.R(solved)
	}

	return d
}

func G0Index(c *model.Cube) uint64 {
	var result uint64 = 0

	for i, v := range c.EO {
		if v {
			result |= 1 << i
		}
	}

	return result
}

func G1Index(c *model.Cube) uint64 {
	var result uint64 = 0

	i := 0

	for _, v := range c.CO {
		result |= uint64(v) << (2 * i)

		i += 1
	}

	for j := 0; j < 12; j += 1 {
		if c.EP[j] < 8 && (c.EP[j]%2) == 1 {
			result |= 1 << ((i * 2) + j)
		}
	}

	return result
}

func G2Index(c *model.Cube) uint64 {
	var result uint64 = 0

	return result
}

func G3Index(c *model.Cube) uint64 {
	var result uint64 = 0

	return result
}

func NewDatabase() Database {
	return Database{
		G0: generate(model.G0, G0Index),
		G1: generate(model.G1, G1Index),
		G2: generate(model.G2, G2Index),
		G3: generate(model.G3, G3Index),
	}
}

func Save(d Database) {
	writeDataToFile(mapToBytes(d.G0), "assets/G0.table")
	writeDataToFile(mapToBytes(d.G1), "assets/G1.table")
	writeDataToFile(mapToBytes(d.G2), "assets/G2.table")
	writeDataToFile(mapToBytes(d.G3), "assets/G3.table")
}
