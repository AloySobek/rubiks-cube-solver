package solver

import (
	"github.com/AloySobek/Rubik/model"
)

type Database struct {
	G0 map[uint64]int
	G1 map[uint64]int
	G2 map[uint64]int
	G3 map[uint64]int

	G0Goal map[uint64]bool
	G1Goal map[uint64]bool
	G2Goal map[uint64]bool
	G3Goal map[uint64]bool
}

type node struct {
	c *model.Cube // Cube representation
	m string      // Cube representation in the form of move sequence from solved
}
