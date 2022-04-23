package cube

import (
	"fmt"
	"math/bits"
)

const (
	Yellow = 0
	Orange = 1
	Green  = 2
	White  = 3
	Blue   = 4
	Red    = 5
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
	if !reverse {
		rotateBlue(cube)

		if double {
			rotateBlue(cube)
		}
	} else {
		if !double {
			bits.RotateLeft64(cube.Blue, 16)
		} else {
			bits.RotateLeft64(cube.Blue, 32)
		}
	}

	red := cube.Red

	cube.Red = (cube.Red & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Yellow & GET_ZERO) | (cube.Yellow & GET_SIX) | (cube.Yellow & GET_SEVEN)
	cube.Yellow = (cube.Yellow & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Orange & GET_ZERO) | (cube.Orange & GET_SIX) | (cube.Orange & GET_SEVEN)
	cube.Orange = (cube.Orange & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.White & GET_ZERO) | (cube.White & GET_SIX) | (cube.White & GET_SEVEN)
	cube.White = (cube.White & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (red & GET_ZERO) | (red & GET_SIX) | (red & GET_SEVEN)
}

func Right(cube *Cube, reverse, double bool) {
	if !reverse {
		if !double {
			bits.RotateLeft64(cube.Red, -16)
		} else {
			bits.RotateLeft64(cube.Red, -32)
		}
	} else {
		if !double {
			bits.RotateLeft64(cube.Red, 16)
		} else {
			bits.RotateLeft64(cube.Red, 32)
		}
	}

	green := cube.Green

	cube.Green = (cube.Green & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Yellow & GET_ZERO) | (cube.Yellow & GET_SIX) | (cube.Yellow & GET_SEVEN)
	cube.Yellow = (cube.Yellow & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Blue & GET_ZERO) | (cube.Blue & GET_SIX) | (cube.Blue & GET_SEVEN)
	cube.Blue = (cube.Blue & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.White & GET_ZERO) | (cube.White & GET_SIX) | (cube.White & GET_SEVEN)
	cube.White = (cube.White & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (green & GET_ZERO) | (green & GET_SIX) | (green & GET_SEVEN)
}

func Left(cube *Cube, reverse, double bool) {
	if !reverse {
		if !double {
			bits.RotateLeft64(cube.Orange, -16)
		} else {
			bits.RotateLeft64(cube.Orange, -32)
		}
	} else {
		if !double {
			bits.RotateLeft64(cube.Orange, 16)
		} else {
			bits.RotateLeft64(cube.Orange, 32)
		}
	}

	blue := cube.Blue

	cube.Blue = (cube.Blue & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Yellow & GET_ZERO) | (cube.Yellow & GET_SIX) | (cube.Yellow & GET_SEVEN)
	cube.Yellow = (cube.Yellow & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Green & GET_ZERO) | (cube.Green & GET_SIX) | (cube.Green & GET_SEVEN)
	cube.Green = (cube.Green & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.White & GET_ZERO) | (cube.White & GET_SIX) | (cube.White & GET_SEVEN)
	cube.White = (cube.White & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (blue & GET_ZERO) | (blue & GET_SIX) | (blue & GET_SEVEN)
}

func Up(cube *Cube, reverse, double bool) {
	if !reverse {
		if !double {
			bits.RotateLeft64(cube.Yellow, -16)
		} else {
			bits.RotateLeft64(cube.Yellow, -32)
		}
	} else {
		if !double {
			bits.RotateLeft64(cube.Yellow, 16)
		} else {
			bits.RotateLeft64(cube.Yellow, 32)
		}
	}

	red := cube.Red

	cube.Red = (cube.Red & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Green & GET_ZERO) | (cube.Green & GET_SIX) | (cube.Green & GET_SEVEN)
	cube.Green = (cube.Green & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Orange & GET_ZERO) | (cube.Orange & GET_SIX) | (cube.Orange & GET_SEVEN)
	cube.Orange = (cube.Orange & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Blue & GET_ZERO) | (cube.Blue & GET_SIX) | (cube.Blue & GET_SEVEN)
	cube.Blue = (cube.Blue & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (red & GET_ZERO) | (red & GET_SIX) | (red & GET_SEVEN)
}

func Down(cube *Cube, reverse, double bool) {
	if !reverse {
		if !double {
			bits.RotateLeft64(cube.White, -16)
		} else {
			bits.RotateLeft64(cube.White, -32)
		}
	} else {
		if !double {
			bits.RotateLeft64(cube.White, 16)
		} else {
			bits.RotateLeft64(cube.White, 32)
		}
	}

	orange := cube.Orange

	cube.Orange = (cube.Orange & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Green & GET_ZERO) | (cube.Green & GET_SIX) | (cube.Green & GET_SEVEN)
	cube.Green = (cube.Green & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Red & GET_ZERO) | (cube.Red & GET_SIX) | (cube.Red & GET_SEVEN)
	cube.Red = (cube.Red & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Blue & GET_ZERO) | (cube.Blue & GET_SIX) | (cube.Blue & GET_SEVEN)
	cube.Blue = (cube.Blue & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (orange & GET_ZERO) | (orange & GET_SIX) | (orange & GET_SEVEN)
}

func Back(cube *Cube, reverse, double bool) {
	if !reverse {
		if !double {
			bits.RotateLeft64(cube.Green, -16)
		} else {
			bits.RotateLeft64(cube.Green, -32)
		}
	} else {
		if !double {
			bits.RotateLeft64(cube.Green, 16)
		} else {
			bits.RotateLeft64(cube.Green, 32)
		}
	}

	orange := cube.Orange

	cube.Orange = (cube.Orange & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Yellow & GET_ZERO) | (cube.Yellow & GET_SIX) | (cube.Yellow & GET_SEVEN)
	cube.Yellow = (cube.Yellow & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Red & GET_ZERO) | (cube.Red & GET_SIX) | (cube.Red & GET_SEVEN)
	cube.Red = (cube.Red & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.White & GET_ZERO) | (cube.White & GET_SIX) | (cube.White & GET_SEVEN)
	cube.White = (cube.White & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (orange & GET_ZERO) | (orange & GET_SIX) | (orange & GET_SEVEN)
}

func Print(cube *Cube) {
	fmt.Println(cube)
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
