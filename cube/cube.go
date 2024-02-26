package cube

import (
	"log"
	"strings"
)

const (
	UL = iota
	UF = iota
	UR = iota
	UB = iota

	DL = iota
	DF = iota
	DR = iota
	DB = iota

	LB = iota
	LF = iota
	RF = iota
	RB = iota
)

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
	EdgeNames    = [12]string{"UL", "UF", "UR", "UB", "DL", "DF", "DR", "DB", "LB", "LF", "RF", "RB"}
	CornerNames  = [8]string{"ULB", "ULF", "URF", "URB", "DLB", "DLF", "DRF", "DRB"}
	OppositeFace = map[byte]byte{'U': 'D', 'L': 'R', 'F': 'B', 'R': 'L', 'B': 'F', 'D': 'U'}
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
