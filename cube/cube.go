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
	ULB = iota // 0
	ULF = iota // 1
	URF = iota // 2
	URB = iota // 3

	DLB = iota // 4
	DLF = iota // 5
	DRF = iota // 6
	DRB = iota // 7
)

var (
	Corners      = [8]uint8{ULB, ULF, URF, URB, DLB, DLF, DRF, DRB}
	EdgeNames    = [12]string{"UL", "UF", "UR", "UB", "DL", "DF", "DR", "DB", "LB", "LF", "RF", "RB"}
	CornerNames  = [8]string{"ULB", "ULF", "URF", "URB", "DLB", "DLF", "DRF", "DRB"}
	OppositeFace = map[byte]byte{'U': 'D', 'L': 'R', 'F': 'B', 'R': 'L', 'B': 'F', 'D': 'U'}
)

const (
	UU = uint64(1)
	LL = uint64(2)
	FF = uint64(4)
	RR = uint64(8)
	BB = uint64(16)
	DD = uint64(32)
)

const (
	SU = uint64(0) | UU | (UU << 8) | (UU << 16) | (UU << 24) | (UU << 32) | (UU << 40) | (UU << 48) | (UU << 56)
	SL = uint64(0) | LL | (LL << 8) | (LL << 16) | (LL << 24) | (LL << 32) | (LL << 40) | (LL << 48) | (LL << 56)
	SF = uint64(0) | FF | (FF << 8) | (FF << 16) | (FF << 24) | (FF << 32) | (FF << 40) | (FF << 48) | (FF << 56)
	SR = uint64(0) | RR | (RR << 8) | (RR << 16) | (RR << 24) | (RR << 32) | (RR << 40) | (RR << 48) | (RR << 56)
	SB = uint64(0) | BB | (BB << 8) | (BB << 16) | (BB << 24) | (BB << 32) | (BB << 40) | (BB << 48) | (BB << 56)
	SD = uint64(0) | DD | (DD << 8) | (DD << 16) | (DD << 24) | (DD << 32) | (DD << 40) | (DD << 48) | (DD << 56)
)

type Cube struct {
	EO [12]bool
	EP [12]uint8
	CO [8]uint8
	CP [8]uint8

	U uint64
	L uint64
	F uint64
	R uint64
	B uint64
	D uint64
}

func Create(c *Cube) *Cube {
	if c != nil {
		return &Cube{
			EO: c.EO,
			EP: c.EP,
			CO: c.CO,
			CP: c.CP,
			U:  c.U,
			L:  c.L,
			F:  c.F,
			R:  c.R,
			B:  c.B,
			D:  c.D,
		}
	}

	return &Cube{
		EO: [12]bool{},
		EP: [12]uint8{UL, UF, UR, UB, DL, DF, DR, DB, LB, LF, RF, RB},
		CO: [8]uint8{},
		CP: [8]uint8{ULB, ULF, URF, URB, DLB, DLF, DRF, DRB},
		U:  SU,
		L:  SL,
		F:  SF,
		R:  SR,
		B:  SB,
		D:  SD,
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
