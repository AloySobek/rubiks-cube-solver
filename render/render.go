package render

import (
	"fmt"
	"github.com/AloySobek/Rubik/cube"
)

const (
	U = 'R'
	L = 'B'
	F = 'W'
	R = 'G'
	B = 'Y'
	D = 'O'
)

func Render(c *cube.Cube) {
	buffer := make([][]rune, 11)

	for i := 0; i < 11; i += 1 {
		buffer[i] = make([]rune, 15)
	}

	// putSide(buffer, cube.U, 'W', 0, 4)
	// putSide(buffer, cube.L, 'G', 4, 0)
	// putSide(buffer, cube.F, 'R', 4, 4)
	// putSide(buffer, cube.R, 'B', 4, 8)}
	// putSide(buffer, cube.B, 'O', 4, 12
	// putSide(buffer, cube.D, 'Y', 8, 4)func Print(cube *Cube) {

	fmt.Println(c.CO)

	// draw(buffer)
}

func draw(buffer [][]rune) {
	fmt.Println()

	for y := 0; y < 11; y += 1 {
		for x := 0; x < 15; x += 1 {
			switch buffer[y][x] {
			case 'B':
				fmt.Printf("🟦")
			case 'R':
				fmt.Printf("🟥")
			case 'O':
				fmt.Printf("🟧")
			case 'G':
				fmt.Printf("🟩")
			case 'Y':
				fmt.Printf("🟨")
			case 'W':
				fmt.Printf("⬜")
			default:
				fmt.Printf("⬛")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func putEdges(buffer [][]rune, side uint64, color rune, iy, ix int) {
	y := iy
	x := ix
	i := 0

	for ; x != ix+3; x, i = x+1, i+1 {
		buffer[y][x] = getSideCubieColor(side, i)
	}

	x, y = x-1, y+1

	for ; y != iy+3; y, i = y+1, i+1 {
		buffer[y][x] = getSideCubieColor(side, i)
	}

	y, x = y-1, x-1

	for ; x != ix-1; x, i = x-1, i+1 {
		buffer[y][x] = getSideCubieColor(side, i)
	}

	x, y = x+1, y-1

	for ; y != iy; y, i = y-1, i+1 {
		buffer[y][x] = getSideCubieColor(side, i)
	}

	y, x = y+1, x+1

	buffer[y][x] = color
}

func getSideCubieColor(side uint64, position int) rune {
	switch (side & (uint64(0xFF) << (8 * position))) >> (8 * position) {
	case U:
		return 'W'
	case L:
		return 'G'
	case F:
		return 'R'
	case R:
		return 'B'
	case B:
		return 'O'
	case D:
		return 'Y'
	default:
		return ' '
	}
}
