package algorithm

import (
	"fmt"

	"github.com/AloySobek/Rubik/cube"
	"github.com/AloySobek/Rubik/graph"
)

func solve(cube *cube.Cube) *cube.Cube {
	var root *graph.Node = graph.CreateNode(&graph.Data{
		Cube: cube,
	})

	fmt.Println(root.Data)

	return cube
}

func IDAStar() {

}
