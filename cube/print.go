package cube

import (
	"fmt"
)

func Print(cube *Cube) {
	buffer := make([][]rune, 11)

	for i := 0; i < 11; i += 1 {
		buffer[i] = make([]rune, 15)
	}

	putSide(buffer, cube.Yellow, 'Y', 0, 4)
	putSide(buffer, cube.Orange, 'O', 4, 0)
	putSide(buffer, cube.Blue, 'B', 4, 4)
	putSide(buffer, cube.Red, 'R', 4, 8)
	putSide(buffer, cube.Green, 'G', 4, 12)
	putSide(buffer, cube.White, 'W', 8, 4)

	render(buffer)
}

func render(buffer [][]rune) {
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
	case Blue:
		return 'B'
	case Red:
		return 'R'
	case Orange:
		return 'O'
	case Green:
		return 'G'
	case Yellow:
		return 'Y'
	case White:
		return 'W'
	default:
		return ' '
	}
}
