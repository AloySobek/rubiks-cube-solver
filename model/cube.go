package model

import (
	"log"
	"strings"
)

const (
	UL = iota // 0
	UF = iota // 1 y
	UR = iota // 2
	UB = iota // 3 y

	DL = iota // 4
	DF = iota // 5 y
	DR = iota // 6
	DB = iota // 7 y

	LB = iota
	LF = iota
	RF = iota
	RB = iota
)

//  0   1   2   3   4   5   6   7
// ULB ULF URF URB DLB DLF DRF DRB

const (
	ULB = iota
	ULF = iota
	URF = iota
	URB = iota

	DLB = iota
	DLF = iota
	DRF = iota
	DRB = iota
)

var (
	Edges   = [12]uint8{UL, UF, UR, UB, DL, DF, DR, DB, LB, LF, RF, RB}
	Corners = [8]uint8{ULB, ULF, URF, URB, DLB, DLF, DRF, DRB}
)

type Cube struct {
	EO [12]bool
	EP [12]uint8
	CO [8]uint8
	CP [8]uint8
}

func Create(c *Cube) *Cube {
	if c != nil {
		return &Cube{
			EO: c.EO,
			EP: c.EP,
			CO: c.CO,
			CP: c.CP,
		}
	}

	return &Cube{
		EO: [12]bool{},
		EP: [12]uint8{UL, UF, UR, UB, DL, DF, DR, DB, LB, LF, RF, RB},
		CO: [8]uint8{},
		CP: [8]uint8{ULB, ULF, URF, URB, DLB, DLF, DRF, DRB},
	}
}

func ApplyMoves(c *Cube, sequence []string) []string {
	for _, v := range sequence {
		if move, ok := G0[strings.ToUpper(v)]; ok {
			move(c)
		} else {
			log.Fatalf("Unsupported move: %s", v)
		}
	}

	return sequence
}
