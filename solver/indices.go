package solver

import "github.com/AloySobek/Rubik/cube"

const (
	G0 = 0
	G1 = 1
	G2 = 2
	G3 = 3
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

func G2Index(c *cube.Cube) uint64 {
	var result uint64 = 0

	// for i := 0; i < 7; i += 1 {
	// 	for j := 0; j < 3; j += 1 {
	// 		result <<= 1

	// 		t := cornerNames[c.CP[i]][j]

	// 		if !(t == cornerNames[i][j] || t == opposite[cornerNames[i][j]]) {
	// 			result += 1
	// 		}
	// 	}
	// }

	// for i := 0; i < 12; i += 1 {
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

	// 	if (int(c.CP[i]) % 4) != (i % 4) {
	// 		result += 1
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
