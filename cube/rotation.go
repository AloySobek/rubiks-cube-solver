package cube

import "math/bits"

const (
	CLR_ZERO  = 0xFFFFFFFFFFFFFF00
	CLR_ONE   = 0xFFFFFFFFFFFF00FF
	CLR_TWO   = 0xFFFFFFFFFF00FFFF
	CLR_THREE = 0xFFFFFFFF00FFFFFF
	CLR_FOUR  = 0xFFFFFF00FFFFFFFF
	CLR_FIVE  = 0xFFFF00FFFFFFFFFF
	CLR_SIX   = 0xFF00FFFFFFFFFFFF
	CLR_SEVEN = 0x00FFFFFFFFFFFFFF

	GET_ZERO  = 0x00000000000000FF
	GET_ONE   = 0x000000000000FF00
	GET_TWO   = 0x0000000000FF0000
	GET_THREE = 0x00000000FF000000
	GET_FOUR  = 0x000000FF00000000
	GET_FIVE  = 0x0000FF0000000000
	GET_SIX   = 0x00FF000000000000
	GET_SEVEN = 0xFF00000000000000
)

func rotateBlue(cube *Cube, reverse bool) {
	if !reverse {
		cube.Blue = bits.RotateLeft64(cube.Blue, 16)

		red := cube.Red

		cube.Red = (cube.Red & CLR_ZERO & CLR_SIX & CLR_SEVEN) | ((cube.Yellow & GET_FOUR) << 16) | ((cube.Yellow & GET_FIVE) << 16) | ((cube.Yellow & GET_SIX) >> 48)
		cube.Yellow = (cube.Yellow & CLR_FOUR & CLR_FIVE & CLR_SIX) | ((cube.Orange & GET_TWO) << 16) | ((cube.Orange & GET_THREE) << 16) | ((cube.Orange & GET_FOUR) << 16)
		cube.Orange = (cube.Orange & CLR_TWO & CLR_THREE & CLR_FOUR) | ((cube.White & GET_ZERO) << 16) | ((cube.White & GET_ONE) << 16) | ((cube.White & GET_TWO) << 16)
		cube.White = (cube.White & CLR_ZERO & CLR_ONE & CLR_TWO) | ((red & GET_ZERO) << 16) | ((red & GET_SIX) >> 48) | ((red & GET_SEVEN) >> 48)
	} else {
		cube.Blue = bits.RotateLeft64(cube.Blue, -16)

		orange := cube.Orange

		cube.Orange = (cube.Orange & CLR_TWO & CLR_THREE & CLR_FOUR) | ((cube.Yellow & GET_FOUR) >> 16) | ((cube.Yellow & GET_FIVE) >> 16) | ((cube.Yellow & GET_SIX) >> 16)
		cube.Yellow = (cube.Yellow & CLR_FOUR & CLR_FIVE & CLR_SIX) | ((cube.Red & GET_ZERO) << 48) | ((cube.Red & GET_SIX) >> 16) | ((cube.Red & GET_SEVEN) >> 16)
		cube.Red = (cube.Red & CLR_ZERO & CLR_SIX & CLR_SEVEN) | ((cube.White & GET_ZERO) << 48) | ((cube.White & GET_ONE) << 48) | ((cube.White & GET_TWO) >> 16)
		cube.White = (cube.White & CLR_ZERO & CLR_ONE & CLR_TWO) | ((orange & GET_TWO) >> 16) | ((orange & GET_THREE) >> 16) | ((orange & GET_FOUR) >> 16)
	}
}

func rotateRed(cube *Cube, reverse bool) {
	if !reverse {
		cube.Red = bits.RotateLeft64(cube.Red, 16)

		green := cube.Green

		cube.Green = (cube.Green & CLR_TWO & CLR_THREE & CLR_FOUR) | (cube.Yellow & GET_TWO) | (cube.Yellow & GET_THREE) | (cube.Yellow & GET_FOUR)
		cube.Yellow = (cube.Yellow & CLR_TWO & CLR_THREE & CLR_FOUR) | (cube.Blue & GET_TWO) | (cube.Blue & GET_THREE) | (cube.Blue & GET_FOUR)
		cube.Blue = (cube.Blue & CLR_TWO & CLR_THREE & CLR_FOUR) | (cube.White & GET_TWO) | (cube.White & GET_THREE) | (cube.White & GET_FOUR)
		cube.White = (cube.White & CLR_TWO & CLR_THREE & CLR_FOUR) | (green & GET_TWO) | (green & GET_THREE) | (green & GET_FOUR)
	} else {
		cube.Red = bits.RotateLeft64(cube.Red, -16)

		blue := cube.Blue

		cube.Blue = (cube.Blue & CLR_TWO & CLR_THREE & CLR_FOUR) | (cube.Yellow & GET_TWO) | (cube.Yellow & GET_THREE) | (cube.Yellow & GET_FOUR)
		cube.Yellow = (cube.Yellow & CLR_TWO & CLR_THREE & CLR_FOUR) | (cube.Green & GET_TWO) | (cube.Green & GET_THREE) | (cube.Green & GET_FOUR)
		cube.Green = (cube.Green & CLR_TWO & CLR_THREE & CLR_FOUR) | (cube.White & GET_TWO) | (cube.White & GET_THREE) | (cube.White & GET_FOUR)
		cube.White = (cube.White & CLR_TWO & CLR_THREE & CLR_FOUR) | (blue & GET_TWO) | (blue & GET_THREE) | (blue & GET_FOUR)
	}
}

func rotateOrange(cube *Cube, reverse bool) {
	if !reverse {
		cube.Orange = bits.RotateLeft64(cube.Orange, 16)

		blue := cube.Blue

		cube.Blue = (cube.Blue & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Yellow & GET_ZERO) | (cube.Yellow & GET_SIX) | (cube.Yellow & GET_SEVEN)
		cube.Yellow = (cube.Yellow & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Green & GET_ZERO) | (cube.Green & GET_SIX) | (cube.Green & GET_SEVEN)
		cube.Green = (cube.Green & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.White & GET_ZERO) | (cube.White & GET_SIX) | (cube.White & GET_SEVEN)
		cube.White = (cube.White & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (blue & GET_ZERO) | (blue & GET_SIX) | (blue & GET_SEVEN)
	} else {
		cube.Orange = bits.RotateLeft64(cube.Orange, -16)

		green := cube.Green

		cube.Green = (cube.Green & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Yellow & GET_ZERO) | (cube.Yellow & GET_SIX) | (cube.Yellow & GET_SEVEN)
		cube.Yellow = (cube.Yellow & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.Blue & GET_ZERO) | (cube.Blue & GET_SIX) | (cube.Blue & GET_SEVEN)
		cube.Blue = (cube.Blue & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (cube.White & GET_ZERO) | (cube.White & GET_SIX) | (cube.White & GET_SEVEN)
		cube.White = (cube.White & CLR_ZERO & CLR_SIX & CLR_SEVEN) | (green & GET_ZERO) | (green & GET_SIX) | (green & GET_SEVEN)
	}
}

func rotateGreen(cube *Cube, reverse bool) {
	if !reverse {
		cube.Green = bits.RotateLeft64(cube.Green, 16)

		red := cube.Red

		cube.Red = (cube.Red & CLR_TWO & CLR_THREE & CLR_FOUR) | ((cube.White & GET_FOUR) >> 16) | ((cube.White & GET_FIVE) >> 16) | ((cube.White & GET_SIX) >> 16)
		cube.White = (cube.White & CLR_FOUR & CLR_FIVE & CLR_SIX) | ((cube.Orange & GET_ZERO) << 48) | ((cube.Orange & GET_SIX) >> 16) | ((cube.Orange & GET_SEVEN) >> 16)
		cube.Orange = (cube.Orange & CLR_ZERO & CLR_SIX & CLR_SEVEN) | ((cube.Yellow & GET_ZERO) << 48) | ((cube.Yellow & GET_ONE) << 48) | ((cube.Yellow & GET_TWO) >> 16)
		cube.Yellow = (cube.Yellow & CLR_ZERO & CLR_ONE & CLR_TWO) | ((red & GET_TWO) >> 16) | ((red & GET_THREE) >> 16) | ((red & GET_FOUR) >> 16)
	} else {
		cube.Green = bits.RotateLeft64(cube.Green, -16)

		orange := cube.Orange

		cube.Orange = (cube.Orange & CLR_ZERO & CLR_SIX & CLR_SEVEN) | ((cube.White & GET_FOUR) << 16) | ((cube.White & GET_FIVE) << 16) | ((cube.White & GET_SIX) >> 48)
		cube.White = (cube.White & CLR_FOUR & CLR_FIVE & CLR_SIX) | ((cube.Red & GET_TWO) << 16) | ((cube.Red & GET_THREE) << 16) | ((cube.Red & GET_FOUR) << 16)
		cube.Red = (cube.Red & CLR_TWO & CLR_THREE & CLR_FOUR) | ((cube.Yellow & GET_ZERO) << 16) | ((cube.Yellow & GET_ONE) << 16) | ((cube.Yellow & GET_TWO) << 16)
		cube.Yellow = (cube.Yellow & CLR_ZERO & CLR_ONE & CLR_TWO) | ((orange & GET_ZERO) << 16) | ((orange & GET_SIX) >> 48) | ((orange & GET_SEVEN) >> 48)
	}
}

func rotateYellow(cube *Cube, reverse bool) {
	if !reverse {
		cube.Yellow = bits.RotateLeft64(cube.Yellow, 16)

		red := cube.Red

		cube.Red = (cube.Red & CLR_ZERO & CLR_ONE & CLR_TWO) | ((cube.Green & GET_FOUR) >> 32) | ((cube.Green & GET_FIVE) >> 32) | ((cube.Green & GET_SIX) >> 32)
		cube.Green = (cube.Green & CLR_FOUR & CLR_FIVE & CLR_SIX) | ((cube.Orange & GET_ZERO) << 32) | ((cube.Orange & GET_ONE) << 32) | ((cube.Orange & GET_TWO) << 32)
		cube.Orange = (cube.Orange & CLR_ZERO & CLR_ONE & CLR_TWO) | (cube.Blue & GET_ZERO) | (cube.Blue & GET_ONE) | (cube.Blue & GET_TWO)
		cube.Blue = (cube.Blue & CLR_ZERO & CLR_ONE & CLR_TWO) | (red & GET_ZERO) | (red & GET_ONE) | (red & GET_TWO)
	} else {
		cube.Yellow = bits.RotateLeft64(cube.Yellow, -16)

		orange := cube.Orange

		cube.Orange = (cube.Orange & CLR_ZERO & CLR_ONE & CLR_TWO) | ((cube.Green & GET_FOUR) >> 32) | ((cube.Green & GET_FIVE) >> 32) | ((cube.Green & GET_SIX) >> 32)
		cube.Green = (cube.Green & CLR_FOUR & CLR_FIVE & CLR_SIX) | ((cube.Red & GET_ZERO) << 32) | ((cube.Red & GET_ONE) << 32) | ((cube.Red & GET_TWO) << 32)
		cube.Red = (cube.Red & CLR_ZERO & CLR_ONE & CLR_TWO) | (cube.Blue & GET_ZERO) | (cube.Blue & GET_ONE) | (cube.Blue & GET_TWO)
		cube.Blue = (cube.Blue & CLR_ZERO & CLR_ONE & CLR_TWO) | (orange & GET_ZERO) | (orange & GET_ONE) | (orange & GET_TWO)
	}
}

func rotateWhite(cube *Cube, reverse bool) {
	if !reverse {
		cube.White = bits.RotateLeft64(cube.White, 16)

		red := cube.Red

		cube.Red = (cube.Red & CLR_FOUR & CLR_FIVE & CLR_SIX) | (cube.Blue & GET_FOUR) | (cube.Blue & GET_FIVE) | (cube.Blue & GET_SIX)
		cube.Blue = (cube.Blue & CLR_FOUR & CLR_FIVE & CLR_SIX) | (cube.Orange & GET_FOUR) | (cube.Orange & GET_FIVE) | (cube.Orange & GET_SIX)
		cube.Orange = (cube.Orange & CLR_FOUR & CLR_FIVE & CLR_SIX) | ((cube.Green & GET_ZERO) << 32) | ((cube.Green & GET_ONE) << 32) | ((cube.Green & GET_TWO) << 32)
		cube.Green = (cube.Green & CLR_ZERO & CLR_ONE & CLR_TWO) | ((red & GET_FOUR) >> 32) | ((red & GET_FIVE) >> 32) | ((red & GET_SIX) >> 32)
	} else {
		cube.White = bits.RotateLeft64(cube.White, -16)

		orange := cube.Orange

		cube.Orange = (cube.Orange & CLR_FOUR & CLR_FIVE & CLR_SIX) | (cube.Blue & GET_FOUR) | (cube.Blue & GET_FIVE) | (cube.Blue & GET_SIX)
		cube.Blue = (cube.Blue & CLR_FOUR & CLR_FIVE & CLR_SIX) | (cube.Red & GET_FOUR) | (cube.Red & GET_FIVE) | (cube.Red & GET_SIX)
		cube.Red = (cube.Red & CLR_FOUR & CLR_FIVE & CLR_SIX) | ((cube.Green & GET_ZERO) << 32) | ((cube.Green & GET_ONE) << 32) | ((cube.Green & GET_TWO) << 32)
		cube.Green = (cube.Green & CLR_ZERO & CLR_ONE & CLR_TWO) | ((orange & GET_FOUR) >> 32) | ((orange & GET_FIVE) >> 32) | ((orange & GET_SIX) >> 32)
	}
}
