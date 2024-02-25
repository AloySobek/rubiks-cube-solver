package solver

import (
	"github.com/AloySobek/Rubik/model"
	"strings"
)

var (
	cornerNames = [8]string{"ULB", "ULF", "URF", "URB", "DLB", "DLF", "DRF", "DRB"}
	edgeNames   = [12]string{"UL", "UF", "UR", "UB", "DL", "DF", "DR", "DB", "LB", "LF", "RF", "RB"}
)

var opposite = map[byte]byte{'U': 'D', 'L': 'R', 'F': 'B', 'R': 'L', 'B': 'F', 'D': 'U'}

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
	return Database{
		G0: bytesToMap(readDataFromFile("assets/G0.table")),
		G1: bytesToMap(readDataFromFile("assets/G1.table")),
		G2: bytesToMap(readDataFromFile("assets/G2.table")),
		G3: bytesToMap(readDataFromFile("assets/G3.table")),

		G0Goal: map[uint64]bool{G0Index(model.Create(nil)): true},
		G1Goal: map[uint64]bool{G1Index(model.Create(nil)): true},
		G2Goal: map[uint64]bool{G2Index(model.Create(nil)): true},
		G3Goal: map[uint64]bool{G3Index(model.Create(nil)): true},
	}
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

	for i := 0; i < 7; i += 1 {
		for j := 0; j < 3; j += 1 {
			result <<= 1

			t := cornerNames[c.CP[i]][(int(c.CO[i])+j)%3]

			if !(t == cornerNames[i][j] || t == opposite[cornerNames[i][j]]) {
				result += 1
			}
		}
	}

	// for i := 0; i < 11; i += 1 {
	// 	for j := 0; j < 2; j += 1 {
	// 		result <<= 1

	// 		t := edgeNames[c.EP[i]][j]

	// 		if !(t == edgeNames[i][j] || t == opposite[edgeNames[i][j]]) {
	// 			result += 1
	// 		}
	// 	}
	// }
	// for i := 0; i < 8; i += 1 {
	// 	result <<= 1
	// 	if int(c.CP[i])%4 != i%4 {
	// 		result++
	// 	}
	// }

	// result <<= 1

	// for i := 0; i < 8; i += 1 {
	// 	for j := i + 1; j < 8; j += 1 {
	// 		if c.CP[i] > c.CP[j] {
	// 			result ^= 1
	// 		} else {
	// 			result ^= 0
	// 		}
	// 	}
	// }

	return result
}

func G3Index(c *model.Cube) uint64 {
	var result uint64 = 0

	for i := 0; i < 8; i += 1 {
		for j := 0; j < 3; j += 1 {
			result <<= 1

			if cornerNames[c.CP[i]][j] == opposite[cornerNames[i][j]] {
				result += 1
			}
		}
	}

	for i := 0; i < 12; i += 1 {
		for j := 0; j < 2; j += 1 {
			result <<= 1

			if edgeNames[c.EP[i]][j] == opposite[edgeNames[i][j]] {
				result += 1
			}
		}
	}

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
