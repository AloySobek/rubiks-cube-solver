package render

import (
	"fmt"
	"github.com/AloySobek/Rubik/cube"
)

const (
	CU = 'R'
	CL = 'B'
	CF = 'W'
	CR = 'G'
	CB = 'Y'
	CD = 'O'
)

func Render(c *cube.Cube) {
	buffer := make([][]rune, 11)

	for i := 0; i < 11; i += 1 {
		buffer[i] = make([]rune, 15)
	}

	putSide(buffer, c.U, CU, 0, 4)
	putSide(buffer, c.L, CL, 4, 0)
	putSide(buffer, c.F, CF, 4, 4)
	putSide(buffer, c.R, CR, 4, 8)
	putSide(buffer, c.B, CB, 4, 12)
	putSide(buffer, c.D, CD, 8, 4)

	render(buffer)
}

func render(buffer [][]rune) {
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

func putSide(buffer [][]rune, side uint64, color rune, iy, ix int) {
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
	case cube.UU:
		return CU
	case cube.LL:
		return CL
	case cube.FF:
		return CF
	case cube.RR:
		return CR
	case cube.BB:
		return CB
	case cube.DD:
		return CD
	default:
		return ' '
	}
}
