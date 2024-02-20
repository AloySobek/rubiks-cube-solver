package algorithm

import (
	"bytes"
	"encoding/gob"
	"strings"

	"github.com/AloySobek/Rubik/cube"
)

type Tables struct {
	G0 []string
	G1 []string
}

func ReadG0Table(data *bytes.Buffer) []string {
	const size = 4096 + 1

	table := make([]string, size)

	for i := range table {
		table[i] = ""
	}

	decoder := gob.NewDecoder(data)

	decoder.Decode(&table)

	return table
}

func GenerateG0Table() *bytes.Buffer {
	const size = 4096 + 1

	table := make([]string, size)

	for i := range table {
		table[i] = ""
	}

	c := cube.Create()

	BFS(c, cube.G0, table, cube.GetEdgeOrientations)

	buffer := bytes.NewBuffer([]byte{})

	encoder := gob.NewEncoder(buffer)

	if err := encoder.Encode(table); err != nil {
		panic(err)
	}

	return buffer
}

func ReadG1Table(data *bytes.Buffer) []string {
	const size = 4096 + 1

	table := make([]string, size)

	for i := range table {
		table[i] = ""
	}

	decoder := gob.NewDecoder(data)

	decoder.Decode(&table)

	return table
}

func GenerateG1Table() *bytes.Buffer {
	const size = 4096 + 1

	table := make([]string, size)

	for i := range table {
		table[i] = ""
	}

	c := cube.Create()

	BFS(c, cube.G1, table, cube.GetCornerOrientationAndFourEdges)

	buffer := bytes.NewBuffer([]byte{})

	encoder := gob.NewEncoder(buffer)

	if err := encoder.Encode(table); err != nil {
		panic(err)
	}

	return buffer
}

func BFS(root *cube.Cube, g map[string]func(*cube.Cube) *cube.Cube, table []string, i func(c *cube.Cube) uint16) {
	queue := make([]struct {
		c *cube.Cube
		m string
		l int
	}, 0)

	queue = append(queue, struct {
		c *cube.Cube
		m string
		l int
	}{root, "", 0})

	for c := queue[0]; len(queue) > 0; {
		c, queue = queue[0], queue[1:]

		for k, v := range g {
			cc := struct {
				c *cube.Cube
				m string
				l int
			}{v(cube.Copy(c.c)), c.m + k + " ", c.l + 1}

			index := i(cc.c)

			if table[index] == "" {
				table[index] = strings.TrimSpace(cc.m)

				queue = append(queue, cc)
			}
		}

	}
}

func GetTables() *Tables {
	return &Tables{
		ReadG0Table(ReadDataFromFile("G0.table")),
		ReadG1Table(ReadDataFromFile("G1.table")),
	}
}
