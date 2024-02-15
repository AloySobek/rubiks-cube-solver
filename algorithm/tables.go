package algorithm

import (
	// "bytes"
	// "encoding/gob"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/AloySobek/Rubik/cube"
)

func GenerateG0Table(filepath string) {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	table := make([]string, 4097)

	c := cube.Create()

	BFS(c, cube.G0, table, 0, "")

	sort.Strings(table)

	for i, v := range table {
		if v != "" {
			fmt.Printf("%d: %s\n", i, v)
		}
	}
}

func BFS(c *cube.Cube, g map[string]func(*cube.Cube) *cube.Cube, table []string, depth int, solution string) {
	if depth > 6 {
		return
	}

	i, cs := 0, make([]struct {
		cube *cube.Cube
		move string
	}, len(g))

	for k, v := range g {
		cs[i] = struct {
			cube *cube.Cube
			move string
		}{cube.Copy(c), k}

		v(cs[i].cube)

		index := cube.GetEdgeOrientations(cs[i].cube)

		if index != 0 && (table[index] == "" || len(strings.Split(table[index], " ")) > len(strings.Split(solution+k+" ", " "))) {
			table[index] = solution + k + " "
		}

		i += 1
	}

	for _, v := range cs {
		BFS(v.cube, g, table, depth+1, solution+v.move+" ")
	}
}

// var table [2048]string
//
// 	buffer := new(bytes.Buffer)
//
// 	encoder := gob.NewEncoder(buffer)
//
// 	if err = encoder.Encode(table); err != nil {
// 		panic(err)
// 	}
//
// 	if _, err = file.Write(buffer.Bytes()); err != nil {
// 		panic(err)
// 	}
