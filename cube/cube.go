package cube

import (
	"log"
	"math/rand"
	"strings"
)

const (
	S0 = uint64(1)
	S1 = uint64(2)
	S2 = uint64(4)
	S3 = uint64(8)
	S4 = uint64(16)
	S5 = uint64(32)
)

const (
	SolS0 = uint64(0) | S0 | (S0 << 8) | (S0 << 16) | (S0 << 24) | (S0 << 32) | (S0 << 40) | (S0 << 48) | (S0 << 56)
	SolS1 = uint64(0) | S1 | (S1 << 8) | (S1 << 16) | (S1 << 24) | (S1 << 32) | (S1 << 40) | (S1 << 48) | (S1 << 56)
	SolS2 = uint64(0) | S2 | (S2 << 8) | (S2 << 16) | (S2 << 24) | (S2 << 32) | (S2 << 40) | (S2 << 48) | (S2 << 56)
	SolS3 = uint64(0) | S3 | (S3 << 8) | (S3 << 16) | (S3 << 24) | (S3 << 32) | (S3 << 40) | (S3 << 48) | (S3 << 56)
	SolS4 = uint64(0) | S4 | (S4 << 8) | (S4 << 16) | (S4 << 24) | (S4 << 32) | (S4 << 40) | (S4 << 48) | (S4 << 56)
	SolS5 = uint64(0) | S5 | (S5 << 8) | (S5 << 16) | (S5 << 24) | (S5 << 32) | (S5 << 40) | (S5 << 48) | (S5 << 56)
)

type Cube struct {
	S0 uint64
	S1 uint64
	S2 uint64
	S3 uint64
	S4 uint64
	S5 uint64
}

var G0 map[string]func(*Cube) *Cube = map[string]func(*Cube) *Cube{
	"F":  RotateS2,
	"B":  RotateS4,
	"R":  RotateS3,
	"L":  RotateS1,
	"U":  RotateS0,
	"D":  RotateS5,
	"F'": RotateS2Thrice,
	"B'": RotateS4Thrice,
	"R'": RotateS3Thrice,
	"L'": RotateS1Thrice,
	"U'": RotateS0Thrice,
	"D'": RotateS5Thrice,
	"F2": RotateS2Twice,
	"B2": RotateS4Twice,
	"R2": RotateS3Twice,
	"L2": RotateS1Twice,
	"U2": RotateS0Twice,
	"D2": RotateS5Twice,
}

var G1 map[string]func(*Cube) *Cube = map[string]func(*Cube) *Cube{
	"F":  RotateS2,
	"B":  RotateS4,
	"R":  RotateS3,
	"L":  RotateS1,
	"F'": RotateS2Thrice,
	"B'": RotateS4Thrice,
	"R'": RotateS3Thrice,
	"L'": RotateS1Thrice,
	"F2": RotateS2Twice,
	"B2": RotateS4Twice,
	"R2": RotateS3Twice,
	"L2": RotateS1Twice,
	"U2": RotateS0Twice,
	"D2": RotateS5Twice,
}

var G2 map[string]func(*Cube) *Cube = map[string]func(*Cube) *Cube{
	"R":  RotateS3,
	"L":  RotateS1,
	"R'": RotateS3Thrice,
	"L'": RotateS1Thrice,
	"F2": RotateS2Twice,
	"B2": RotateS4Twice,
	"R2": RotateS3Twice,
	"L2": RotateS1Twice,
	"U2": RotateS0Twice,
	"D2": RotateS5Twice,
}

var G3 map[string]func(*Cube) *Cube = map[string]func(*Cube) *Cube{
	"F2": RotateS2Twice,
	"B2": RotateS4Twice,
	"R2": RotateS3Twice,
	"L2": RotateS1Twice,
	"U2": RotateS0Twice,
	"D2": RotateS5Twice,
}

func Create() *Cube {
	return &Cube{
		S0: SolS0,
		S1: SolS1,
		S2: SolS2,
		S3: SolS3,
		S4: SolS4,
		S5: SolS5,
	}
}

func Copy(c *Cube) *Cube {
	return &Cube{
		S0: c.S0,
		S1: c.S1,
		S2: c.S2,
		S3: c.S3,
		S4: c.S4,
		S5: c.S5,
	}
}

func GetRandomMixSequence(n int) (sequence string) {
	possibleMoves := strings.Split("F R L B U D F' R' L' B' U' D' F2 R2 L2 B2 U2 D2", " ")

	for i := 0; i < n; i += 1 {

		sequence += possibleMoves[rand.Intn(len(possibleMoves))]

		if i+1 != n {
			sequence += " "
		}
	}

	return
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
	return c.S0 == SolS0 && c.S1 == SolS1 &&
		c.S2 == SolS2 && c.S3 == SolS3 &&
		c.S4 == SolS4 && c.S5 == SolS5
}

func G0Condition(c *Cube) bool {
	var bad uint64 = S1 | S3

	if ((c.S2&g1)>>8)&bad > 0 ||
		((c.S2&g3)>>24)&bad > 0 ||
		((c.S2&g5)>>40)&bad > 0 ||
		((c.S2&g7)>>56)&bad > 0 {
		return false
	}

	if ((c.S4&g1)>>8)&bad > 0 ||
		((c.S4&g3)>>24)&bad > 0 ||
		((c.S4&g5)>>40)&bad > 0 ||
		((c.S4&g7)>>56)&bad > 0 {
		return false
	}

	if ((c.S0&g3)>>24)&bad > 0 ||
		((c.S0&g7)>>56)&bad > 0 {
		return false
	}

	if ((c.S5&g3)>>24)&bad > 0 ||
		((c.S5&g7)>>56)&bad > 0 {
		return false
	}

	return true
}

func G1Condition(c *Cube) bool {
	var good uint64 = S1 | S3
	var edges uint64 = S5 | S0 | S2 | S4

	if ((c.S0&g1)>>8)&edges == 0 ||
		((c.S0&g5)>>40)&edges == 0 {
		return false
	}

	if ((c.S2&g1)>>8)&edges == 0 ||
		((c.S2&g5)>>40)&edges == 0 {
		return false
	}

	if ((c.S4&g1)>>8)&edges == 0 ||
		((c.S4&g5)>>40)&edges == 0 {
		return false
	}

	if ((c.S5&g1)>>8)&edges == 0 ||
		((c.S5&g5)>>40)&edges == 0 {
		return false
	}

	if (c.S1&g0)&good == 0 ||
		((c.S1&g2)>>16)&good == 0 ||
		((c.S1&g4)>>32)&good == 0 ||
		((c.S1&g6)>>48)&good == 0 {
		return false
	}

	if (c.S3&g0)&good == 0 ||
		((c.S3&g2)>>16)&good == 0 ||
		((c.S3&g4)>>32)&good == 0 ||
		((c.S3&g6)>>48)&good == 0 {
		return false
	}

	return true
}

func G2Condition(c *Cube) bool {
	if (c.S0&g0)&(S0|S5) == 0 ||
		((c.S0&g1)>>8)&(S0|S5) == 0 ||
		((c.S0&g2)>>16)&(S0|S5) == 0 ||
		((c.S0&g3)>>24)&(S0|S5) == 0 ||
		((c.S0&g4)>>32)&(S0|S5) == 0 ||
		((c.S0&g5)>>40)&(S0|S5) == 0 ||
		((c.S0&g6)>>48)&(S0|S5) == 0 ||
		((c.S0&g7)>>56)&(S0|S5) == 0 {
		return false
	}

	if (c.S5&g0)&(S0|S5) == 0 ||
		((c.S5&g1)>>8)&(S0|S5) == 0 ||
		((c.S5&g2)>>16)&(S0|S5) == 0 ||
		((c.S5&g3)>>24)&(S0|S5) == 0 ||
		((c.S5&g4)>>32)&(S0|S5) == 0 ||
		((c.S5&g5)>>40)&(S0|S5) == 0 ||
		((c.S5&g6)>>48)&(S0|S5) == 0 ||
		((c.S5&g7)>>56)&(S0|S5) == 0 {
		return false
	}

	if (c.S4&g0)&(S4|S2) == 0 ||
		((c.S4&g1)>>8)&(S4|S2) == 0 ||
		((c.S4&g2)>>16)&(S4|S2) == 0 ||
		((c.S4&g3)>>24)&(S4|S2) == 0 ||
		((c.S4&g4)>>32)&(S4|S2) == 0 ||
		((c.S4&g5)>>40)&(S4|S2) == 0 ||
		((c.S4&g6)>>48)&(S4|S2) == 0 ||
		((c.S4&g7)>>56)&(S4|S2) == 0 {
		return false
	}

	if (c.S2&g0)&(S4|S2) == 0 ||
		((c.S2&g1)>>8)&(S4|S2) == 0 ||
		((c.S2&g2)>>16)&(S4|S2) == 0 ||
		((c.S2&g3)>>24)&(S4|S2) == 0 ||
		((c.S2&g4)>>32)&(S4|S2) == 0 ||
		((c.S2&g5)>>40)&(S4|S2) == 0 ||
		((c.S2&g6)>>48)&(S4|S2) == 0 ||
		((c.S2&g7)>>56)&(S4|S2) == 0 {
		return false
	}

	if (c.S1&g0)&(S1|S3) == 0 ||
		((c.S1&g1)>>8)&(S1|S3) == 0 ||
		((c.S1&g2)>>16)&(S1|S3) == 0 ||
		((c.S1&g3)>>24)&(S1|S3) == 0 ||
		((c.S1&g4)>>32)&(S1|S3) == 0 ||
		((c.S1&g5)>>40)&(S1|S3) == 0 ||
		((c.S1&g6)>>48)&(S1|S3) == 0 ||
		((c.S1&g7)>>56)&(S1|S3) == 0 {
		return false
	}

	if (c.S3&g0)&(S1|S3) == 0 ||
		((c.S3&g1)>>8)&(S1|S3) == 0 ||
		((c.S3&g2)>>16)&(S1|S3) == 0 ||
		((c.S3&g3)>>24)&(S1|S3) == 0 ||
		((c.S3&g4)>>32)&(S1|S3) == 0 ||
		((c.S3&g5)>>40)&(S1|S3) == 0 ||
		((c.S3&g6)>>48)&(S1|S3) == 0 ||
		((c.S3&g7)>>56)&(S1|S3) == 0 {
		return false
	}

	return true
}

func G3Condition(c *Cube) bool {
	return Solved(c)
}
