package cube

import "math/bits"

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
