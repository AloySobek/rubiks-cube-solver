package algorithm

import (
	"bytes"
	"encoding/gob"
	"io"
	"os"
	"strings"

	"github.com/AloySobek/Rubik/cube"
)

func WriteDataToFile(filepath string, data *bytes.Buffer) {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	if _, err := file.Write(data.Bytes()); err != nil {
		panic(err)
	}
}

func ReadDataFromFile(filepath string) *bytes.Buffer {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	buffer := bytes.NewBuffer([]byte{})

	io.Copy(buffer, file)

	return buffer
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

		if index != 0 && (table[index] == "") {
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

// var table [2048]string
//
