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

func BFS2(root *model.Cube, g map[string]func(*model.Cube) *model.Cube, table map[uint32]string, i func(c *model.Cube) uint32) {
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

func GenerateG1() map[uint32]string {
	table := make(map[uint32]string, 0)

	c := model.Create(nil)

	BFS2(c, solver.G0, table, func(c *model.Cube) uint32 {
		var result uint32 = 0

		i := 0

		for _, v := range c.CO {
			result |= uint32(v) << (2 * i)

			i += 1
		}

		for _, v := range c.EP {
			if v < 8 {
				result |= 1 << (2 * i)

				i += 1
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

func G1ToBytes(table map[uint32]string) *bytes.Buffer {
	buffer := bytes.NewBuffer([]byte{})

	encoder := gob.NewEncoder(buffer)

	if err := encoder.Encode(table); err != nil {
		panic(err)
	}

	return buffer
}
