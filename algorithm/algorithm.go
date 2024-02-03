package algorithm

import (
	"github.com/AloySobek/Rubik/cube"
	"strings"
)

type Try struct {
	cube *cube.Cube
	move string
}

// Stupidiest  algo ever, just playing
func solve(c *cube.Cube, solution string, lvl int) string {
	if lvl > 5 {
		return ""
	}

	i, cs := 0, make([]Try, len(cube.NotationMatrix))

	for k, v := range cube.NotationMatrix {
		cs[i] = Try{cube.Copy(c), k}

		v(cs[i].cube, false, false)

		if cube.IsSolved(cs[i].cube) {
			return solution + k + " "
		}

		i += 1
	}

	result := ""

	for _, v := range cs {
		res := solve(v.cube, solution+v.move+" ", lvl+1)

		if res != "" {
			result = res

			break
		}
	}

	return result
}

func Solve(c *cube.Cube) string {
	return strings.TrimRight(solve(c, "", 0), " ")
}
