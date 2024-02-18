package algorithm

import (
	"bytes"
	"encoding/gob"
	"strings"

	"github.com/AloySobek/Rubik/cube"
)

type Tables struct {
	G0 []string
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

	BFS(c, cube.G0, table)

	buffer := bytes.NewBuffer([]byte{})

	encoder := gob.NewEncoder(buffer)

	if err := encoder.Encode(table); err != nil {
		panic(err)
	}

	return buffer
}

func BFS(root *cube.Cube, g map[string]func(*cube.Cube) *cube.Cube, table []string) {
	queue := make([]struct {
		c *cube.Cube
		m string
		l uint64
	}, 0)

	queue = append(queue, struct {
		c *cube.Cube
		m string
		l uint64
	}{root, "", 0})

	for len(queue) > 0 {
		c := queue[0]

		queue = queue[1:]

		index := cube.GetEdgeOrientations(c.c)

		if index != 0 && table[index] == "" {
			table[index] = strings.TrimSpace(c.m)
		}

		if c.l < 7 {
			for k, v := range g {
				queue = append(queue, struct {
					c *cube.Cube
					m string
					l uint64
				}{v(cube.Copy(c.c)), c.m + k + " ", c.l + 1})
			}
		}
	}
}

func GetTables() *Tables {
	return &Tables{
		ReadG0Table(ReadDataFromFile("G0.table")),
	}
}
