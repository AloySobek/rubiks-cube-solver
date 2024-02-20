package database

import (
	"bytes"
	"encoding/gob"
	"github.com/AloySobek/Rubik/model"
	"github.com/AloySobek/Rubik/solver"
	"strings"
)

func BFS(root *model.Cube, g map[string]func(*model.Cube) *model.Cube, table map[uint16]string, i func(c *model.Cube) uint16) {
	queue := make([]struct {
		c *model.Cube
		m string
		l int
	}, 0)

	queue = append(queue, struct {
		c *model.Cube
		m string
		l int
	}{root, "", 0})

	for c := queue[0]; len(queue) > 0; {
		c, queue = queue[0], queue[1:]

		for k, v := range g {
			cc := struct {
				c *model.Cube
				m string
				l int
			}{v(model.Create(c.c)), c.m + k + " ", c.l + 1}

			index := i(cc.c)

			if _, ok := table[index]; !ok {
				table[index] = strings.TrimSpace(cc.m)

				queue = append(queue, cc)
			}
		}

	}
}

func GenerateG0() map[uint16]string {
	table := make(map[uint16]string, 0)

	c := model.Create(nil)

	BFS(c, solver.G0, table, func(c *model.Cube) uint16 {
		var result uint16 = 0

		for i, v := range c.EO {
			if v {
				result |= 1 << i
			}
		}

		return result
	})

	return table
}

func G0ToBytes(table map[uint16]string) *bytes.Buffer {
	buffer := bytes.NewBuffer([]byte{})

	encoder := gob.NewEncoder(buffer)

	if err := encoder.Encode(table); err != nil {
		panic(err)
	}

	return buffer
}
