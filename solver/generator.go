package solver

import (
	"github.com/AloySobek/Rubik/model"
	"strings"
)

func bfs(root *model.Cube, g map[string]func(*model.Cube) *model.Cube, table map[uint64]string, i func(c *model.Cube) uint64) {
	queue := make([]node, 0)

	queue = append(queue, node{root, ""})

	for c := queue[0]; len(queue) > 0; {
		c, queue = queue[0], queue[1:]

		for k, v := range g {
			cc := node{v(model.Create(c.c)), c.m + k + " "}

			index := i(cc.c)

			if _, ok := table[index]; !ok {
				table[index] = strings.TrimSpace(cc.m)

				queue = append(queue, cc)
			}
		}

	}
}

func generate(g map[string]func(*model.Cube) *model.Cube, f func(c *model.Cube) uint64) map[uint64]string {
	table := make(map[uint64]string, 0)

	c := model.Create(nil)

	bfs(c, g, table, f)

	return table
}

func DatabaseFromFile() Database {
	return Database{
		G0: bytesToMap(readDataFromFile("assets/G0.table")),
		G1: bytesToMap(readDataFromFile("assets/G1.table")),
		G2: bytesToMap(readDataFromFile("assets/G2.table")),
		G3: bytesToMap(readDataFromFile("assets/G3.table")),
	}
}

func NewDatabase() Database {
	return Database{
		G0: generate(model.G0, func(c *model.Cube) uint64 {
			var result uint64 = 0

			for i, v := range c.EO {
				if v {
					result |= 1 << i
				}
			}

			return result
		}),
		G1: generate(model.G1, func(c *model.Cube) uint64 {
			var result uint64 = 0

			i := 0

			for _, v := range c.CO {
				result |= uint64(v) << (2 * i)

				i += 1
			}

			for j := 0; j < 12; j += 1 {
				if c.EP[j] < 8 {
					result |= 1 << ((i * 2) + (2 * j))
				}
			}

			return result
		}),
		G2: generate(model.G2, func(c *model.Cube) uint64 {
			var result uint64 = 0

			return result
		}),
		G3: generate(model.G3, func(c *model.Cube) uint64 {
			var result uint64 = 0

			return result
		}),
	}
}

func Save(d Database) {
	writeDataToFile(mapToBytes(d.G0), "assets/G0.table")
	writeDataToFile(mapToBytes(d.G1), "assets/G1.table")
	writeDataToFile(mapToBytes(d.G2), "assets/G2.table")
	writeDataToFile(mapToBytes(d.G3), "assets/G3.table")
}
