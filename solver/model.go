package solver

import (
	"github.com/AloySobek/Rubik/cube"
)

type Database struct {
	Tables []map[uint64]int
	Goals  []map[uint64]bool
}

type node struct {
	c *cube.Cube // Cube representation
	m string     // Cube representation in the form of move sequence from solved
}
