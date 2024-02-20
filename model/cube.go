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

type Cube struct {
	EO [12]bool
}

func Create(c *Cube) *Cube {
	if c != nil {
		return &Cube{
			EO: c.EO,
		}
	}

	return &Cube{
		EO: [12]bool{},
	}
}
