package model

const (
	UL = iota
	UF = iota
	UR = iota
	UB = iota

	LB = iota
	LF = iota
	RF = iota
	RB = iota

	DL = iota
	DF = iota
	DR = iota
	DB = iota
)

//  0   1   2   3   4   5   6   7
// ULB ULF URF URB DLB DLF DRF DRB
//

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
		EP: [12]uint8{UL, UF, UR, UB, LB, LF, RF, RB, DL, DB, DR, DF},
		CO: [8]uint8{},
		CP: [8]uint8{ULB, ULF, URF, URB, DLB, DLF, DRF, DRB},
	}
}
