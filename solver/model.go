package solver

import (
	"github.com/AloySobek/Rubik/model"
)

type Database struct {
	G0 map[uint64]string
	G1 map[uint64]string
	G2 map[uint64]string
	G3 map[uint64]string
}

type node struct {
	c *model.Cube // Cube representation
	m string      // Cube representation in the form of move sequence from solved
}
