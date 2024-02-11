package algorithm

import (
	"fmt"
	"github.com/AloySobek/Rubik/cube"
	"strings"
)

type Try struct {
	cube *cube.Cube
	move string
}

func solve(c *cube.Cube, solution string, lvl, maxDepth int) string {
	if lvl > maxDepth {
		return ""
	}

	i, cs := 0, make([]Try, len(cube.NotationMatrix))

	for k, v := range cube.NotationMatrix {
		cs[i] = Try{cube.Copy(c), k}

		v(cs[i].cube, false)

		if cube.IsSolved(cs[i].cube) {
			return solution + k + " "
		}

		i += 1
	}

	result := ""

	for _, v := range cs {
		res := solve(v.cube, solution+v.move+" ", lvl+1, maxDepth)

		if res != "" {
			result = res

			return res
		}
	}

	return result
}

func G0ToG1(c *cube.Cube, solution string, lvl int) string {
	if lvl > 7 {
		return ""
	}

	i, cs := 0, make([]Try, len(cube.NotationMatrix))

	for k, v := range cube.NotationMatrix {
		cs[i] = Try{cube.Copy(c), k}

		v(cs[i].cube, false)

		if cube.IsGoodEdges(cs[i].cube) {
			return solution + k + " "
		}

		i += 1
	}

	result := ""

	for _, v := range cs {
		res := G0ToG1(v.cube, solution+v.move+" ", lvl+1)

		if res != "" {
			result = res

			return res
		}
	}

	return result

}

func G1ToG2(c *cube.Cube, solution string, lvl int) string {
	if lvl > 10 {
		return ""
	}

	i, cs := 0, make([]Try, len(cube.NotationMatrix1))

	for k, v := range cube.NotationMatrix1 {
		cs[i] = Try{cube.Copy(c), k}

		v(cs[i].cube, false)

		if cube.IsGoodCorners(cs[i].cube) {
			return solution + k + " "
		}

		i += 1
	}

	result := ""

	for _, v := range cs {
		res := G1ToG2(v.cube, solution+v.move+" ", lvl+1)

		if res != "" {
			result = res

			return res
		}
	}

	return result

}

func G2ToG3(c *cube.Cube, solution string, lvl int) string {
	if lvl > 13 {
		return ""
	}

	i, cs := 0, make([]Try, len(cube.NotationMatrix2))

	for k, v := range cube.NotationMatrix2 {
		cs[i] = Try{cube.Copy(c), k}

		v(cs[i].cube, false)

		if cube.IsGoodSides(cs[i].cube) {
			return solution + k + " "
		}

		i += 1
	}

	result := ""

	for _, v := range cs {
		res := G2ToG3(v.cube, solution+v.move+" ", lvl+1)

		if res != "" {
			result = res

			return res
		}
	}

	return result

}

func G3ToG4(c *cube.Cube, solution string, lvl int) string {
	if lvl > 10 {
		return ""
	}

	i, cs := 0, make([]Try, len(cube.NotationMatrix3))

	for k, v := range cube.NotationMatrix3 {
		cs[i] = Try{cube.Copy(c), k}

		v(cs[i].cube, false)

		if cube.IsSolved(cs[i].cube) {
			return solution + k + " "
		}

		i += 1
	}

	result := ""

	for _, v := range cs {
		res := G3ToG4(v.cube, solution+v.move+" ", lvl+1)

		if res != "" {
			result = res

			return res
		}
	}

	return result

}

func Solve(c *cube.Cube, maxDepth int) string {
	result := strings.TrimRight(G0ToG1(c, "", 0), " ")

	c = cube.ApplyMoves(c, strings.Split(result, " "), nil)

	fmt.Println()
	cube.Print(c)

	result = strings.TrimRight(G1ToG2(c, "", 0), " ")

	c = cube.ApplyMoves(c, strings.Split(result, " "), nil)

	fmt.Println()
	cube.Print(c)

	result = strings.TrimRight(G2ToG3(c, "", 0), " ")

	c = cube.ApplyMoves(c, strings.Split(result, " "), nil)

	fmt.Println()
	cube.Print(c)

	result = strings.TrimRight(G3ToG4(c, "", 0), " ")

	return result
}
