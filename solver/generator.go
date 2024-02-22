package solver

import (
	"github.com/AloySobek/Rubik/model"
	"strings"
)

var (
	cornerNames = [8]string{"ULB", "ULF", "URF", "URB", "DLB", "DLF", "DRF", "DRB"}
	edgeNames   = [12]string{"UL", "UF", "UR", "UB", "DL", "DF", "DR", "DB", "LB", "LF", "RF", "RB"}
	faces       = "FRUBLD"
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

func boolToInt(v bool) int {
	if v {
		return 1
	}

	return 0
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

			return 1

			// for i := 0; i < 7; i += 1 {
			// 	for j := 0; j < 3; j += 1 {
			// 		result <<= 1

			// 		t := cornerNames[c.CP[i]][(int(c.CO[i])+j)%3]

			// 		if !(t == cornerNames[i][j] || t == faces[(strings.Index(faces, string(cornerNames[i][j]))+3)%6]) {
			// 			result += 1
			// 		}
			// 	}
			// }

			// for i := 0; i < 11; i += 1 {
			// 	for j := 0; j < 2; j += 1 {
			// 		result <<= 1

			// 		t := edgeNames[c.EP[i]][(boolToInt(c.EO[i])+j)%2]

			// 		if !(t == edgeNames[i][j] || t == faces[(strings.Index(faces, string(edgeNames[i][j]))+3)%6]) {
			// 			result += 1
			// 		}
			// 	}
			// }

			// for i := 0; i < 8; i += 1 {
			// 	result <<= 1

			// 	if (int(c.CP[i]) % 4) != (i % 4) {
			// 		result += 1
			// 	}
			// }

			// result <<= 1

			// for i := 0; i < 8; i += 1 {
			// 	for j := i + 1; j < 8; j += 1 {
			// 		result ^= uint64(boolToInt(c.CP[i] > c.CP[j]))
			// 	}
			// }

			return result
		}),
		G3: generate(model.G3, func(c *model.Cube) uint64 {
			var result uint64 = 0

			return 1

			// offset := 0

			// for i := 0; i < 12; i, offset = i+1, offset+1 {
			// 	if c.EP[i] != model.Edges[i] {
			// 		result |= 1 << offset
			// 	}
			// }

			// for i := 0; i < 8; i, offset = i+1, offset+1 {
			// 	if c.CP[i] != model.Corners[i] {
			// 		result |= 1 << offset
			// 	}
			// }

			// for i := 0; i < 8; i += 1 {
			// 	for j := 0; j < 3; j += 1 {
			// 		result <<= 1

			// 		t := cornerNames[c.CP[i]][(int(c.CO[i])+j)%3]

			// 		if t == faces[(strings.Index(faces, string(cornerNames[i][j]))+3)%6] {
			// 			result += 1
			// 		}
			// 	}
			// }

			// for i := 0; i < 12; i += 1 {
			// 	for j := 0; j < 2; j += 1 {
			// 		result <<= 1

			// 		t := edgeNames[c.EP[i]][(boolToInt(c.EO[i])+j)%2]

			// 		if t == faces[(strings.Index(faces, string(edgeNames[i][j]))+3)%6] {
			// 			result += 1
			// 		}
			// 	}
			// }

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
