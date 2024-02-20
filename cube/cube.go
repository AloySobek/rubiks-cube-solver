package cube

import (
	"log"
	"strings"
)

const (
	U = uint64(1)
	L = uint64(2)
	F = uint64(4)
	R = uint64(8)
	B = uint64(16)
	D = uint64(32)
)

const (
	SU = uint64(0) | U | (U << 8) | (U << 16) | (U << 24) | (U << 32) | (U << 40) | (U << 48) | (U << 56)
	SL = uint64(0) | L | (L << 8) | (L << 16) | (L << 24) | (L << 32) | (L << 40) | (L << 48) | (L << 56)
	SF = uint64(0) | F | (F << 8) | (F << 16) | (F << 24) | (F << 32) | (F << 40) | (F << 48) | (F << 56)
	SR = uint64(0) | R | (R << 8) | (R << 16) | (R << 24) | (R << 32) | (R << 40) | (R << 48) | (R << 56)
	SB = uint64(0) | B | (B << 8) | (B << 16) | (B << 24) | (B << 32) | (B << 40) | (B << 48) | (B << 56)
	SD = uint64(0) | D | (D << 8) | (D << 16) | (D << 24) | (D << 32) | (D << 40) | (D << 48) | (D << 56)
)

type Cube struct {
	U uint64
	L uint64
	F uint64
	R uint64
	B uint64
	D uint64
}

type Cube2 struct {
}

func Create() *Cube {
	return &Cube{
		U: SU,
		L: SL,
		F: SF,
		R: SR,
		B: SB,
		D: SD,
	}
}

func Copy(c *Cube) *Cube {
	return &Cube{
		U: c.U,
		L: c.L,
		F: c.F,
		R: c.R,
		B: c.B,
		D: c.D,
	}
}

func ApplyMoves(cube *Cube, sequence []string) []string {
	for _, c := range sequence {
		if move, ok := G0[strings.ToUpper(c)]; ok {
			move(cube)
		} else {
			log.Fatalf("Unsupported move: %s", c)
		}
	}

	return sequence
}

func Solved(c *Cube) bool {
	return c.U == SU && c.L == SL &&
		c.F == SF && c.R == SR &&
		c.B == SB && c.D == SD
}

func G0Condition(c *Cube) bool {
	var bad uint64 = L | R

	if ((c.F&g1)>>8)&bad > 0 ||
		((c.F&g3)>>24)&bad > 0 ||
		((c.F&g5)>>40)&bad > 0 ||
		((c.F&g7)>>56)&bad > 0 {
		return false
	}

	if ((c.B&g1)>>8)&bad > 0 ||
		((c.B&g3)>>24)&bad > 0 ||
		((c.B&g5)>>40)&bad > 0 ||
		((c.B&g7)>>56)&bad > 0 {
		return false
	}

	if ((c.U&g3)>>24)&bad > 0 ||
		((c.U&g7)>>56)&bad > 0 {
		return false
	}

	if ((c.D&g3)>>24)&bad > 0 ||
		((c.D&g7)>>56)&bad > 0 {
		return false
	}

	return true
}

func G1Condition(c *Cube) bool {
	var good uint64 = L | R
	var edges uint64 = D | U | F | B

	if ((c.U&g1)>>8)&edges == 0 ||
		((c.U&g5)>>40)&edges == 0 {
		return false
	}

	if ((c.F&g1)>>8)&edges == 0 ||
		((c.F&g5)>>40)&edges == 0 {
		return false
	}

	if ((c.B&g1)>>8)&edges == 0 ||
		((c.B&g5)>>40)&edges == 0 {
		return false
	}

	if ((c.D&g1)>>8)&edges == 0 ||
		((c.D&g5)>>40)&edges == 0 {
		return false
	}

	if (c.L&g0)&good == 0 ||
		((c.L&g2)>>16)&good == 0 ||
		((c.L&g4)>>32)&good == 0 ||
		((c.L&g6)>>48)&good == 0 {
		return false
	}

	if (c.R&g0)&good == 0 ||
		((c.R&g2)>>16)&good == 0 ||
		((c.R&g4)>>32)&good == 0 ||
		((c.R&g6)>>48)&good == 0 {
		return false
	}

	return true
}

func G2Condition(c *Cube) bool {
	if (c.U&g0)&(U|D) == 0 ||
		((c.U&g1)>>8)&(U|D) == 0 ||
		((c.U&g2)>>16)&(U|D) == 0 ||
		((c.U&g3)>>24)&(U|D) == 0 ||
		((c.U&g4)>>32)&(U|D) == 0 ||
		((c.U&g5)>>40)&(U|D) == 0 ||
		((c.U&g6)>>48)&(U|D) == 0 ||
		((c.U&g7)>>56)&(U|D) == 0 {
		return false
	}

	if (c.D&g0)&(U|D) == 0 ||
		((c.D&g1)>>8)&(U|D) == 0 ||
		((c.D&g2)>>16)&(U|D) == 0 ||
		((c.D&g3)>>24)&(U|D) == 0 ||
		((c.D&g4)>>32)&(U|D) == 0 ||
		((c.D&g5)>>40)&(U|D) == 0 ||
		((c.D&g6)>>48)&(U|D) == 0 ||
		((c.D&g7)>>56)&(U|D) == 0 {
		return false
	}

	if (c.B&g0)&(B|F) == 0 ||
		((c.B&g1)>>8)&(B|F) == 0 ||
		((c.B&g2)>>16)&(B|F) == 0 ||
		((c.B&g3)>>24)&(B|F) == 0 ||
		((c.B&g4)>>32)&(B|F) == 0 ||
		((c.B&g5)>>40)&(B|F) == 0 ||
		((c.B&g6)>>48)&(B|F) == 0 ||
		((c.B&g7)>>56)&(B|F) == 0 {
		return false
	}

	if (c.F&g0)&(B|F) == 0 ||
		((c.F&g1)>>8)&(B|F) == 0 ||
		((c.F&g2)>>16)&(B|F) == 0 ||
		((c.F&g3)>>24)&(B|F) == 0 ||
		((c.F&g4)>>32)&(B|F) == 0 ||
		((c.F&g5)>>40)&(B|F) == 0 ||
		((c.F&g6)>>48)&(B|F) == 0 ||
		((c.F&g7)>>56)&(B|F) == 0 {
		return false
	}

	if (c.L&g0)&(L|R) == 0 ||
		((c.L&g1)>>8)&(L|R) == 0 ||
		((c.L&g2)>>16)&(L|R) == 0 ||
		((c.L&g3)>>24)&(L|R) == 0 ||
		((c.L&g4)>>32)&(L|R) == 0 ||
		((c.L&g5)>>40)&(L|R) == 0 ||
		((c.L&g6)>>48)&(L|R) == 0 ||
		((c.L&g7)>>56)&(L|R) == 0 {
		return false
	}

	if (c.R&g0)&(L|R) == 0 ||
		((c.R&g1)>>8)&(L|R) == 0 ||
		((c.R&g2)>>16)&(L|R) == 0 ||
		((c.R&g3)>>24)&(L|R) == 0 ||
		((c.R&g4)>>32)&(L|R) == 0 ||
		((c.R&g5)>>40)&(L|R) == 0 ||
		((c.R&g6)>>48)&(L|R) == 0 ||
		((c.R&g7)>>56)&(L|R) == 0 {
		return false
	}

	return true
}

func G3Condition(c *Cube) bool {
	return Solved(c)
}

func GetEdgeOrientations(c *Cube) uint16 {
	var result uint16 = 0

	if ((c.F&g1)>>8)&(L|R) > 0 {
		result |= 1
	}

	if ((c.F&g3)>>24)&(L|R) > 0 {
		result |= 2
	}

	if ((c.F&g5)>>40)&(L|R) > 0 {
		result |= 4
	}

	if ((c.F&g7)>>56)&(L|R) > 0 {
		result |= 8
	}

	if ((c.U&g3)>>24)&(L|R) > 0 {
		result |= 16
	}

	if ((c.U&g7)>>56)&(L|R) > 0 {
		result |= 32
	}

	if ((c.D&g3)>>24)&(L|R) > 0 {
		result |= 64
	}

	if ((c.D&g7)>>56)&(L|R) > 0 {
		result |= 128
	}

	if ((c.B&g1)>>8)&(L|R) > 0 {
		result |= 256
	}

	if ((c.B&g3)>>24)&(L|R) > 0 {
		result |= 512
	}

	if ((c.B&g5)>>40)&(L|R) > 0 {
		result |= 1024
	}

	if ((c.B&g7)>>56)&(L|R) > 0 {
		result |= 2048
	}

	return result
}

func GetCornerOrientationAndFourEdges(c *Cube) uint16 {
	var result uint16 = 0

	if (c.L&g0)&(L|R) == 0 {
		result |= 1
	}

	if ((c.L&g2)>>16)&(L|R) == 0 {
		result |= 2
	}

	if ((c.L&g4)>>32)&(L|R) == 0 {
		result |= 4
	}

	if ((c.L&g6)>>48)&(L|R) == 0 {
		result |= 8
	}

	if (c.R&g0)&(L|R) == 0 {
		result |= 16
	}

	if ((c.R&g2)>>16)&(L|R) == 0 {
		result |= 32
	}

	if ((c.R&g4)>>32)&(L|R) == 0 {
		result |= 64
	}

	if ((c.R&g6)>>48)&(L|R) == 0 {
		result |= 128
	}

	if ((c.L&g1)>>8)&(D|U|F|B) > 0 {
		result |= 256
	}

	if ((c.L&g5)>>40)&(D|U|F|B) > 0 {
		result |= 512
	}

	if ((c.R&g1)>>8)&(D|U|F|B) > 0 {
		result |= 1024
	}

	if ((c.R&g5)>>40)&(D|U|F|B) > 0 {
		result |= 2048
	}

	return result
}
