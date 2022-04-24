package cube

import (
	"fmt"
	"math/bits"
)

const (
	Yellow = 1
	Orange = 2
	Green  = 4
	White  = 8
	Blue   = 16
	Red    = 32
)

const (
	CLR_ZERO  = 0xFFFFFFFFFFFFFF00
	CLR_SIX   = 0xFF00FFFFFFFFFFFF
	CLR_SEVEN = 0x00FFFFFFFFFFFFFF
	GET_ZERO  = 0x00000000000000FF
	GET_SIX   = 0x00FF000000000000
	GET_SEVEN = 0xFF00000000000000
)

type Cube struct {
	Yellow uint64 // Up side
	Orange uint64 // Left side
	Green  uint64 // Back side
	White  uint64 // Down side
	Blue   uint64 // Front side
	Red    uint64 // Right side
}

func Create() *Cube {
	return &Cube{
		Yellow: uint64(0) | Yellow | (Yellow << 8) | (Yellow << 16) | (Yellow << 24) | (Yellow << 32) | (Yellow << 40) | (Yellow << 48) | (Yellow << 56),
		Orange: uint64(0) | Orange | (Orange << 8) | (Orange << 16) | (Orange << 24) | (Orange << 32) | (Orange << 40) | (Orange << 48) | (Orange << 56),
		Green:  uint64(0) | Green | (Green << 8) | (Green << 16) | (Green << 24) | (Green << 32) | (Green << 40) | (Green << 48) | (Green << 56),
		White:  uint64(0) | White | (White << 8) | (White << 16) | (White << 24) | (White << 32) | (White << 40) | (White << 48) | (White << 56),
		Blue:   uint64(0) | Blue | (Blue << 8) | (Blue << 16) | (Blue << 24) | (Blue << 32) | (Blue << 40) | (Blue << 48) | (Blue << 56),
		Red:    uint64(0) | Red | (Red << 8) | (Red << 16) | (Red << 24) | (Red << 32) | (Red << 40) | (Red << 48) | (Red << 56),
	}
}

func Front(cube *Cube, reverse, double bool) {
	rotateBlue(cube, reverse)

	if double {
		rotateBlue(cube, reverse)
	}
}

func Right(cube *Cube, reverse, double bool) {
	rotateRed(cube, reverse)

	if double {
		rotateRed(cube, reverse)
	}
}

func Left(cube *Cube, reverse, double bool) {
	rotateOrange(cube, reverse)

	if double {
		rotateOrange(cube, reverse)
	}
}

func Back(cube *Cube, reverse, double bool) {
	rotateGreen(cube, reverse)

	if double {
		rotateGreen(cube, reverse)
	}
}

func Up(cube *Cube, reverse, double bool) {
	rotateYellow(cube, reverse)

	if double {
		rotateYellow(cube, reverse)
	}
}

func Down(cube *Cube, reverse, double bool) {
	rotateWhite(cube, reverse)

	if double {
		rotateWhite(cube, reverse)
	}
}

func Print(cube *Cube) {
	fmt.Println("Blue side:")
	printSide(cube.Blue, "B")
	fmt.Printf("\n\n")

	fmt.Println("Red side:")
	printSide(cube.Red, "R")
	fmt.Printf("\n\n")

	fmt.Println("Orange side:")
	printSide(cube.Orange, "O")
	fmt.Printf("\n\n")

	fmt.Println("Green side:")
	printSide(cube.Green, "G")
	fmt.Printf("\n\n")

	fmt.Println("Yellow side:")
	printSide(cube.Yellow, "Y")
	fmt.Printf("\n\n")

	fmt.Println("White side:")
	printSide(cube.White, "W")
	fmt.Printf("\n\n")
}

func printSide(side uint64, name string) {
	for i := 0; i < 8; i += 1 {
		switch (side & (uint64(0xFF) << (8 * i))) >> (8 * i) {
		case Blue:
			fmt.Printf("B ")
		case Red:
			fmt.Printf("R ")
		case Orange:
			fmt.Printf("O ")
		case Green:
			fmt.Printf("G ")
		case Yellow:
			fmt.Printf("Y ")
		case White:
			fmt.Printf("W ")
		}

		if i == 3 {
			fmt.Printf("%s ", name)
		}
	}
}

func rotateBlue(cube *Cube, reverse bool) {
	if !reverse {
		bits.RotateLeft64(cube.Blue, -16)

		red := cube.Red

		cube.Red = (cube.Red & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Yellow & GET_ZERO) | (cube.Yellow & GET_SIX) | (cube.Yellow & GET_SEVEN)
		cube.Yellow = (cube.Yellow & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Orange & GET_ZERO) | (cube.Orange & GET_SIX) | (cube.Orange & GET_SEVEN)
		cube.Orange = (cube.Orange & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.White & GET_ZERO) | (cube.White & GET_SIX) | (cube.White & GET_SEVEN)
		cube.White = (cube.White & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (red & GET_ZERO) | (red & GET_SIX) | (red & GET_SEVEN)
	} else {
		bits.RotateLeft64(cube.Blue, 16)

		orange := cube.Orange

		cube.Orange = (cube.Orange & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Yellow & GET_ZERO) | (cube.Yellow & GET_SIX) | (cube.Yellow & GET_SEVEN)
		cube.Yellow = (cube.Yellow & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Red & GET_ZERO) | (cube.Red & GET_SIX) | (cube.Red & GET_SEVEN)
		cube.Red = (cube.Red & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.White & GET_ZERO) | (cube.White & GET_SIX) | (cube.White & GET_SEVEN)
		cube.White = (cube.White & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (orange & GET_ZERO) | (orange & GET_SIX) | (orange & GET_SEVEN)
	}
}

func rotateRed(cube *Cube, reverse bool) {
	if !reverse {
		bits.RotateLeft64(cube.Red, -16)

		green := cube.Green

		cube.Green = (cube.Green & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Yellow & GET_ZERO) | (cube.Yellow & GET_SIX) | (cube.Yellow & GET_SEVEN)
		cube.Yellow = (cube.Yellow & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Blue & GET_ZERO) | (cube.Blue & GET_SIX) | (cube.Blue & GET_SEVEN)
		cube.Blue = (cube.Blue & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.White & GET_ZERO) | (cube.White & GET_SIX) | (cube.White & GET_SEVEN)
		cube.White = (cube.White & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (green & GET_ZERO) | (green & GET_SIX) | (green & GET_SEVEN)
	} else {
		bits.RotateLeft64(cube.Red, 16)

		blue := cube.Blue

		cube.Blue = (cube.Blue & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Yellow & GET_ZERO) | (cube.Yellow & GET_SIX) | (cube.Yellow & GET_SEVEN)
		cube.Yellow = (cube.Yellow & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Green & GET_ZERO) | (cube.Green & GET_SIX) | (cube.Green & GET_SEVEN)
		cube.Green = (cube.Green & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.White & GET_ZERO) | (cube.White & GET_SIX) | (cube.White & GET_SEVEN)
		cube.White = (cube.White & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (blue & GET_ZERO) | (blue & GET_SIX) | (blue & GET_SEVEN)
	}
}

func rotateOrange(cube *Cube, reverse bool) {
	if !reverse {
		bits.RotateLeft64(cube.Orange, -16)

		blue := cube.Blue

		cube.Blue = (cube.Blue & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Yellow & GET_ZERO) | (cube.Yellow & GET_SIX) | (cube.Yellow & GET_SEVEN)
		cube.Yellow = (cube.Yellow & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Green & GET_ZERO) | (cube.Green & GET_SIX) | (cube.Green & GET_SEVEN)
		cube.Green = (cube.Green & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.White & GET_ZERO) | (cube.White & GET_SIX) | (cube.White & GET_SEVEN)
		cube.White = (cube.White & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (blue & GET_ZERO) | (blue & GET_SIX) | (blue & GET_SEVEN)
	} else {
		bits.RotateLeft64(cube.Orange, 16)

		green := cube.Green

		cube.Green = (cube.Green & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Yellow & GET_ZERO) | (cube.Yellow & GET_SIX) | (cube.Yellow & GET_SEVEN)
		cube.Yellow = (cube.Yellow & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Blue & GET_ZERO) | (cube.Blue & GET_SIX) | (cube.Blue & GET_SEVEN)
		cube.Blue = (cube.Blue & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.White & GET_ZERO) | (cube.White & GET_SIX) | (cube.White & GET_SEVEN)
		cube.White = (cube.White & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (green & GET_ZERO) | (green & GET_SIX) | (green & GET_SEVEN)
	}
}

func rotateGreen(cube *Cube, reverse bool) {
	if !reverse {
		bits.RotateLeft64(cube.Green, -16)

		orange := cube.Orange

		cube.Orange = (cube.Orange & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Yellow & GET_ZERO) | (cube.Yellow & GET_SIX) | (cube.Yellow & GET_SEVEN)
		cube.Yellow = (cube.Yellow & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Red & GET_ZERO) | (cube.Red & GET_SIX) | (cube.Red & GET_SEVEN)
		cube.Red = (cube.Red & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.White & GET_ZERO) | (cube.White & GET_SIX) | (cube.White & GET_SEVEN)
		cube.White = (cube.White & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (orange & GET_ZERO) | (orange & GET_SIX) | (orange & GET_SEVEN)
	} else {
		bits.RotateLeft64(cube.Green, 16)

		red := cube.Red

		cube.Red = (cube.Red & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Yellow & GET_ZERO) | (cube.Yellow & GET_SIX) | (cube.Yellow & GET_SEVEN)
		cube.Yellow = (cube.Yellow & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Orange & GET_ZERO) | (cube.Orange & GET_SIX) | (cube.Orange & GET_SEVEN)
		cube.Orange = (cube.Orange & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.White & GET_ZERO) | (cube.White & GET_SIX) | (cube.White & GET_SEVEN)
		cube.White = (cube.White & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (red & GET_ZERO) | (red & GET_SIX) | (red & GET_SEVEN)
	}
}

func rotateYellow(cube *Cube, reverse bool) {
	if !reverse {
		bits.RotateLeft64(cube.Yellow, -16)

		red := cube.Red

		cube.Red = (cube.Red & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Green & GET_ZERO) | (cube.Green & GET_SIX) | (cube.Green & GET_SEVEN)
		cube.Green = (cube.Green & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Orange & GET_ZERO) | (cube.Orange & GET_SIX) | (cube.Orange & GET_SEVEN)
		cube.Orange = (cube.Orange & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Blue & GET_ZERO) | (cube.Blue & GET_SIX) | (cube.Blue & GET_SEVEN)
		cube.Blue = (cube.Blue & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (red & GET_ZERO) | (red & GET_SIX) | (red & GET_SEVEN)
	} else {
		bits.RotateLeft64(cube.Yellow, 16)

		orange := cube.Orange

		cube.Orange = (cube.Orange & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Green & GET_ZERO) | (cube.Green & GET_SIX) | (cube.Green & GET_SEVEN)
		cube.Green = (cube.Green & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Red & GET_ZERO) | (cube.Red & GET_SIX) | (cube.Red & GET_SEVEN)
		cube.Red = (cube.Red & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Blue & GET_ZERO) | (cube.Blue & GET_SIX) | (cube.Blue & GET_SEVEN)
		cube.Blue = (cube.Blue & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (orange & GET_ZERO) | (orange & GET_SIX) | (orange & GET_SEVEN)
	}
}

func rotateWhite(cube *Cube, reverse bool) {
	if !reverse {
		bits.RotateLeft64(cube.White, -16)

		orange := cube.Orange

		cube.Orange = (cube.Orange & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Green & GET_ZERO) | (cube.Green & GET_SIX) | (cube.Green & GET_SEVEN)
		cube.Green = (cube.Green & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Red & GET_ZERO) | (cube.Red & GET_SIX) | (cube.Red & GET_SEVEN)
		cube.Red = (cube.Red & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Blue & GET_ZERO) | (cube.Blue & GET_SIX) | (cube.Blue & GET_SEVEN)
		cube.Blue = (cube.Blue & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (orange & GET_ZERO) | (orange & GET_SIX) | (orange & GET_SEVEN)
	} else {
		bits.RotateLeft64(cube.White, 16)

		red := cube.Red

		cube.Red = (cube.Red & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Green & GET_ZERO) | (cube.Green & GET_SIX) | (cube.Green & GET_SEVEN)
		cube.Green = (cube.Green & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Orange & GET_ZERO) | (cube.Orange & GET_SIX) | (cube.Orange & GET_SEVEN)
		cube.Orange = (cube.Orange & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Blue & GET_ZERO) | (cube.Blue & GET_SIX) | (cube.Blue & GET_SEVEN)
		cube.Blue = (cube.Blue & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (red & GET_ZERO) | (red & GET_SIX) | (red & GET_SEVEN)
	}
}
