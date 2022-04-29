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

func RFront(cube *Cube, _, _ bool) {
	Front(cube, true, false)
}

func RRight(cube *Cube, _, _ bool) {
	Right(cube, true, false)
}

func RLeft(cube *Cube, _, _ bool) {
	Left(cube, true, false)
}

func RBack(cube *Cube, _, _ bool) {
	Back(cube, true, false)
}

func RUp(cube *Cube, _, _ bool) {
	Up(cube, true, false)
}

func RDown(cube *Cube, _, _ bool) {
	Down(cube, true, false)
}

func DFront(cube *Cube, _, _ bool) {
	Front(cube, false, true)
}

func DRight(cube *Cube, _, _ bool) {
	Right(cube, false, true)
}

func DLeft(cube *Cube, _, _ bool) {
	Left(cube, false, true)
}

func DBack(cube *Cube, _, _ bool) {
	Back(cube, false, true)
}

func DUp(cube *Cube, _, _ bool) {
	Up(cube, false, true)
}

func DDown(cube *Cube, _, _ bool) {
	Down(cube, false, true)
}

func RDFront(cube *Cube, _, _ bool) {
	Front(cube, true, true)
}

func RDRight(cube *Cube, _, _ bool) {
	Right(cube, true, true)
}

func RDLeft(cube *Cube, _, _ bool) {
	Left(cube, true, true)
}

func RDBack(cube *Cube, _, _ bool) {
	Back(cube, true, true)
}

func RDUp(cube *Cube, _, _ bool) {
	Up(cube, true, true)
}

func RDDown(cube *Cube, _, _ bool) {
	Down(cube, true, true)
}
