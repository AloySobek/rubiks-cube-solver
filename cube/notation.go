package cube

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
