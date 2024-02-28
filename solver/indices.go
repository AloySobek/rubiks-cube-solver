package solver

import "github.com/AloySobek/Rubik/cube"

const (
	// Clear
	c0 = 0xFFFFFFFFFFFFFF00
	c1 = 0xFFFFFFFFFFFF00FF
	c2 = 0xFFFFFFFFFF00FFFF
	c3 = 0xFFFFFFFF00FFFFFF
	c4 = 0xFFFFFF00FFFFFFFF
	c5 = 0xFFFF00FFFFFFFFFF
	c6 = 0xFF00FFFFFFFFFFFF
	c7 = 0x00FFFFFFFFFFFFFF

	// Get
	g0 = 0x00000000000000FF
	g1 = 0x000000000000FF00
	g2 = 0x0000000000FF0000
	g3 = 0x00000000FF000000
	g4 = 0x000000FF00000000
	g5 = 0x0000FF0000000000
	g6 = 0x00FF000000000000
	g7 = 0xFF00000000000000
)

var huh = []uint64{g0, g1, g2, g3, g4, g5, g6, g7}

const (
	U = uint64(1)
	L = uint64(2)
	F = uint64(4)
	R = uint64(8)
	B = uint64(16)
	D = uint64(32)
)

func G0Index(c *cube.Cube) uint64 {
	var result uint64 = 0

	for i, v := range c.EO {
		if v {
			result |= 1 << i
		}
	}

	return result
}

func G1Index(c *cube.Cube) uint64 {
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

func containsAll(a []int, b []int) bool {
	set := make(map[int]bool)
	for _, v := range b {
		set[v] = true
	}
	for _, v := range a {
		if !set[v] {
			return false
		}
	}
	return true
}

func G2Index(c *cube.Cube) uint64 {
	var result uint64 = 0

	for i := 0; i < 8; i += 1 {
		result <<= 1

		if ((c.U&huh[i])>>(8*i))&(U|D) == 0 {
			result += 1
		}
	}

	for i := 0; i < 8; i += 1 {
		result <<= 1

		if ((c.L&huh[i])>>(8*i))&(L|R) == 0 {
			result += 1
		}
	}

	for i := 0; i < 8; i += 1 {
		result <<= 1

		if ((c.F&huh[i])>>(8*i))&(F|B) == 0 {
			result += 1
		}
	}

	for i := 0; i < 8; i += 1 {
		result <<= 1

		if ((c.R&huh[i])>>(8*i))&(L|R) == 0 {
			result += 1
		}
	}

	for i := 0; i < 8; i += 1 {
		result <<= 1

		if ((c.B&huh[i])>>(8*i))&(F|B) == 0 {
			result += 1
		}
	}

	for i := 0; i < 8; i += 1 {
		result <<= 1

		if ((c.D&huh[i])>>(8*i))&(U|D) == 0 {
			result += 1
		}
	}

	// for i := 0; i < 8; i += 1 {
	// 	for j := 0; j < 3; j += 1 {
	// 		result <<= 1

	// 		// if c.CP[i] == cube.DRF {
	// 		// 	fmt.Printf("Huh: %c : %c\n",
	// 		// 		cube.CornerNames[c.CP[i]][j],
	// 		// 		cube.CornerNames[i][j],
	// 		// 	)
	// 		// }

	// 		t := cube.CornerNames[c.CP[i]][j]

	// 		if t != cube.CornerNames[i][j] && t != cube.OppositeFace[cube.CornerNames[i][j]] {
	// 			result += 1
	// 		}
	// 	}
	// }

	// for i := 0; i < 12; i += 1 {
	// 	for j := 0; j < 2; j += 1 {
	// 		result <<= 1

	// 		t := cube.EdgeNames[c.EP[i]][j]

	// 		if !(t == cube.EdgeNames[i][j] || t == cube.OppositeFace[cube.EdgeNames[i][j]]) {
	// 			result += 1
	// 		}
	// 	}
	// }

	for i := 0; i < 8; i += 1 {
		result <<= 1

		if c.CP[i] < 4 {
			result += 1
		}
	}

	for i := 0; i < 8; i += 1 {
		result <<= 1

		if c.CP[i] >= 4 {
			result += 1
		}
	}

	result <<= 1

	for i := 0; i < 8; i += 1 {
		for j := i + 1; j < 8; j += 1 {
			if c.CP[i] > c.CP[j] {
				result ^= 1
			} else {
				result ^= 0
			}
		}
	}

	return result
}

func G3Index(c *cube.Cube) uint64 {
	var result uint64 = 0

	for i := 0; i < 8; i += 1 {
		for j := 0; j < 3; j += 1 {
			result <<= 1

			if cube.CornerNames[c.CP[i]][j] == cube.OppositeFace[cube.CornerNames[i][j]] {
				result += 1
			}
		}
	}

	for i := 0; i < 12; i += 1 {
		for j := 0; j < 2; j += 1 {
			result <<= 1

			if cube.EdgeNames[c.EP[i]][j] == cube.OppositeFace[cube.EdgeNames[i][j]] {
				result += 1
			}
		}
	}

	return result
}

var indices = []func(c *cube.Cube) uint64{G0Index, G1Index, G2Index, G3Index}
