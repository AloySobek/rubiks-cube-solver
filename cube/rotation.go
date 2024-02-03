package cube

import "math/bits"

const (
	CLR_0 = 0xFFFFFFFFFFFFFF00
	CLR_1 = 0xFFFFFFFFFFFF00FF
	CLR_2 = 0xFFFFFFFFFF00FFFF
	CLR_3 = 0xFFFFFFFF00FFFFFF
	CLR_4 = 0xFFFFFF00FFFFFFFF
	CLR_5 = 0xFFFF00FFFFFFFFFF
	CLR_6 = 0xFF00FFFFFFFFFFFF
	CLR_7 = 0x00FFFFFFFFFFFFFF

	GET_0 = 0x00000000000000FF
	GET_1 = 0x000000000000FF00
	GET_2 = 0x0000000000FF0000
	GET_3 = 0x00000000FF000000
	GET_4 = 0x000000FF00000000
	GET_5 = 0x0000FF0000000000
	GET_6 = 0x00FF000000000000
	GET_7 = 0xFF00000000000000
)

func rotateBlue(cube *Cube, reverse bool) {
	red := cube.Red

	if !reverse {
		cube.Blue = bits.RotateLeft64(cube.Blue, 16)

		cube.Red = (cube.Red & CLR_0 & CLR_6 & CLR_7) | bits.RotateLeft64((cube.Yellow&GET_4)|(cube.Yellow&GET_5)|(cube.Yellow&GET_6), 16)
		cube.Yellow = (cube.Yellow & CLR_4 & CLR_5 & CLR_6) | bits.RotateLeft64((cube.Orange&GET_2)|(cube.Orange&GET_3)|(cube.Orange&GET_4), 16)
		cube.Orange = (cube.Orange & CLR_2 & CLR_3 & CLR_4) | bits.RotateLeft64((cube.White&GET_0)|(cube.White&GET_1)|(cube.White&GET_2), 16)
		cube.White = (cube.White & CLR_0 & CLR_1 & CLR_2) | bits.RotateLeft64((red&GET_0)|(red&GET_6)|(red&GET_7), 16)
	} else {
		cube.Blue = bits.RotateLeft64(cube.Blue, -16)

		cube.Red = (cube.Red & CLR_0 & CLR_6 & CLR_7) | bits.RotateLeft64((cube.White&GET_0)|(cube.White&GET_1)|(cube.White&GET_2), -16)
		cube.White = (cube.White & CLR_0 & CLR_1 & CLR_2) | bits.RotateLeft64((cube.Orange&GET_2)|(cube.Orange&GET_3)|(cube.Orange&GET_4), -16)
		cube.Orange = (cube.Orange & CLR_2 & CLR_3 & CLR_4) | bits.RotateLeft64((cube.Yellow&GET_4)|(cube.Yellow&GET_5)|(cube.Yellow&GET_6), -16)
		cube.Yellow = (cube.Yellow & CLR_4 & CLR_5 & CLR_6) | bits.RotateLeft64((red&GET_0)|(red&GET_6)|(red&GET_7), -16)
	}
}

func rotateRed(cube *Cube, reverse bool) {
	green := cube.Green

	if !reverse {
		cube.Red = bits.RotateLeft64(cube.Red, 16)

		cube.Green = (cube.Green & CLR_0 & CLR_6 & CLR_7) | bits.RotateLeft64((cube.Yellow&GET_2)|(cube.Yellow&GET_3)|(cube.Yellow&GET_4), 32)
		cube.Yellow = (cube.Yellow & CLR_2 & CLR_3 & CLR_4) | (cube.Blue & GET_2) | (cube.Blue & GET_3) | (cube.Blue & GET_4)
		cube.Blue = (cube.Blue & CLR_2 & CLR_3 & CLR_4) | (cube.White & GET_2) | (cube.White & GET_3) | (cube.White & GET_4)
		cube.White = (cube.White & CLR_2 & CLR_3 & CLR_4) | bits.RotateLeft64((green&GET_0)|(green&GET_6)|(green&GET_7), 32)
	} else {
		cube.Red = bits.RotateLeft64(cube.Red, -16)

		cube.Green = (cube.Green & CLR_0 & CLR_6 & CLR_7) | bits.RotateLeft64((cube.White&GET_2)|(cube.White&GET_3)|(cube.White&GET_4), 32)
		cube.White = (cube.White & CLR_2 & CLR_3 & CLR_4) | (cube.Blue & GET_2) | (cube.Blue & GET_3) | (cube.Blue & GET_4)
		cube.Blue = (cube.Blue & CLR_2 & CLR_3 & CLR_4) | (cube.Yellow & GET_2) | (cube.Yellow & GET_3) | (cube.Yellow & GET_4)
		cube.Yellow = (cube.Yellow & CLR_2 & CLR_3 & CLR_4) | bits.RotateLeft64((green&GET_0)|(green&GET_6)|(green&GET_7), 32)
	}
}

func rotateOrange(cube *Cube, reverse bool) {
	blue := cube.Blue

	if !reverse {
		cube.Orange = bits.RotateLeft64(cube.Orange, 16)

		cube.Blue = (cube.Blue & CLR_0 & CLR_6 & CLR_7) | (cube.Yellow & GET_0) | (cube.Yellow & GET_6) | (cube.Yellow & GET_7)
		cube.Yellow = (cube.Yellow & CLR_0 & CLR_6 & CLR_7) | bits.RotateLeft64((cube.Green&GET_2)|(cube.Green&GET_3)|(cube.Green&GET_4), -32)
		cube.Green = (cube.Green & CLR_2 & CLR_3 & CLR_4) | bits.RotateLeft64((cube.White&GET_0)|(cube.White&GET_6)|(cube.White&GET_7), 32)
		cube.White = (cube.White & CLR_0 & CLR_6 & CLR_7) | (blue & GET_0) | (blue & GET_6) | (blue & GET_7)
	} else {
		cube.Orange = bits.RotateLeft64(cube.Orange, -16)

		cube.Blue = (cube.Blue & CLR_0 & CLR_6 & CLR_7) | (cube.White & GET_0) | (cube.White & GET_6) | (cube.White & GET_7)
		cube.White = (cube.White & CLR_0 & CLR_6 & CLR_7) | bits.RotateLeft64((cube.Green&GET_2)|(cube.Green&GET_3)|(cube.Green&GET_4), -32)
		cube.Green = (cube.Green & CLR_2 & CLR_3 & CLR_4) | bits.RotateLeft64((cube.Yellow&GET_0)|(cube.Yellow&GET_6)|(cube.Yellow&GET_7), 32)
		cube.Yellow = (cube.Yellow & CLR_0 & CLR_6 & CLR_7) | (blue & GET_0) | (blue & GET_6) | (blue & GET_7)
	}
}

func rotateGreen(cube *Cube, reverse bool) {
	orange := cube.Orange

	if !reverse {
		cube.Green = bits.RotateLeft64(cube.Green, 16)

		cube.Orange = (cube.Orange & CLR_0 & CLR_6 & CLR_7) | bits.RotateLeft64((cube.Yellow&GET_0)|(cube.Yellow&GET_1)|(cube.Yellow&GET_2), -16)
		cube.Yellow = (cube.Yellow & CLR_0 & CLR_1 & CLR_2) | bits.RotateLeft64((cube.Red&GET_2)|(cube.Red&GET_3)|(cube.Red&GET_4), -16)
		cube.Red = (cube.Red & CLR_2 & CLR_3 & CLR_4) | bits.RotateLeft64((cube.White&GET_4)|(cube.White&GET_5)|(cube.White&GET_6), -16)
		cube.White = (cube.White & CLR_4 & CLR_5 & CLR_6) | bits.RotateLeft64((orange&GET_0)|(orange&GET_6)|(orange&GET_7), -16)
	} else {
		cube.Green = bits.RotateLeft64(cube.Green, -16)

		cube.Orange = (cube.Orange & CLR_0 & CLR_6 & CLR_7) | bits.RotateLeft64((cube.White&GET_4)|(cube.White&GET_5)|(cube.White&GET_6), 16)
		cube.White = (cube.White & CLR_4 & CLR_5 & CLR_6) | bits.RotateLeft64((cube.Red&GET_2)|(cube.Red&GET_3)|(cube.Red&GET_4), 16)
		cube.Red = (cube.Red & CLR_2 & CLR_3 & CLR_4) | bits.RotateLeft64((cube.Yellow&GET_0)|(cube.Yellow&GET_1)|(cube.Yellow&GET_2), 16)
		cube.Yellow = (cube.Yellow & CLR_0 & CLR_1 & CLR_2) | bits.RotateLeft64((orange&GET_0)|(orange&GET_6)|(orange&GET_7), 16)
	}
}

func rotateYellow(cube *Cube, reverse bool) {
	red := cube.Red

	if !reverse {
		cube.Yellow = bits.RotateLeft64(cube.Yellow, 16)

		cube.Red = (cube.Red & CLR_0 & CLR_1 & CLR_2) | (cube.Green & GET_0) | (cube.Green & GET_1) | (cube.Green & GET_2)
		cube.Green = (cube.Green & CLR_0 & CLR_1 & CLR_2) | (cube.Orange & GET_0) | (cube.Orange & GET_1) | (cube.Orange & GET_2)
		cube.Orange = (cube.Orange & CLR_0 & CLR_1 & CLR_2) | (cube.Blue & GET_0) | (cube.Blue & GET_1) | (cube.Blue & GET_2)
		cube.Blue = (cube.Blue & CLR_0 & CLR_1 & CLR_2) | (red & GET_0) | (red & GET_1) | (red & GET_2)
	} else {
		cube.Yellow = bits.RotateLeft64(cube.Yellow, -16)

		cube.Red = (cube.Red & CLR_0 & CLR_1 & CLR_2) | (cube.Blue & GET_0) | (cube.Blue & GET_1) | (cube.Blue & GET_2)
		cube.Blue = (cube.Blue & CLR_0 & CLR_1 & CLR_2) | (cube.Orange & GET_0) | (cube.Orange & GET_1) | (cube.Orange & GET_2)
		cube.Orange = (cube.Orange & CLR_0 & CLR_1 & CLR_2) | (cube.Green & GET_0) | (cube.Green & GET_1) | (cube.Green & GET_2)
		cube.Green = (cube.Green & CLR_0 & CLR_1 & CLR_2) | (red & GET_0) | (red & GET_1) | (red & GET_2)
	}
}

func rotateWhite(cube *Cube, reverse bool) {
	red := cube.Red

	if !reverse {
		cube.White = bits.RotateLeft64(cube.White, 16)

		cube.Red = (cube.Red & CLR_4 & CLR_5 & CLR_6) | (cube.Blue & GET_4) | (cube.Blue & GET_5) | (cube.Blue & GET_6)
		cube.Blue = (cube.Blue & CLR_4 & CLR_5 & CLR_6) | (cube.Orange & GET_4) | (cube.Orange & GET_5) | (cube.Orange & GET_6)
		cube.Orange = (cube.Orange & CLR_4 & CLR_5 & CLR_6) | (cube.Green & GET_4) | (cube.Green & GET_5) | (cube.Green & GET_6)
		cube.Green = (cube.Green & CLR_4 & CLR_5 & CLR_6) | (red & GET_4) | (red & GET_5) | (red & GET_6)

	} else {
		cube.White = bits.RotateLeft64(cube.White, -16)

		cube.Red = (cube.Red & CLR_4 & CLR_5 & CLR_6) | (cube.Green & GET_4) | (cube.Green & GET_5) | (cube.Green & GET_6)
		cube.Green = (cube.Green & CLR_4 & CLR_5 & CLR_6) | (cube.Orange & GET_4) | (cube.Orange & GET_5) | (cube.Orange & GET_6)
		cube.Orange = (cube.Orange & CLR_4 & CLR_5 & CLR_6) | (cube.Blue & GET_4) | (cube.Blue & GET_5) | (cube.Blue & GET_6)
		cube.Blue = (cube.Blue & CLR_4 & CLR_5 & CLR_6) | (red & GET_4) | (red & GET_5) | (red & GET_6)
	}
}
