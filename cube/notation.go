package cube

func Up(cube *Cube, reverse bool) {
	rotateYellow(cube, reverse)
}

func Down(cube *Cube, reverse bool) {
	rotateWhite(cube, reverse)
}

func Right(cube *Cube, reverse bool) {
	rotateRed(cube, reverse)
}

func Left(cube *Cube, reverse bool) {
	rotateOrange(cube, reverse)
}

func Front(cube *Cube, reverse bool) {
	rotateBlue(cube, reverse)
}

func Back(cube *Cube, reverse bool) {
	rotateGreen(cube, reverse)
}

func RFront(cube *Cube, _ bool) {
	Front(cube, true)
}

func RRight(cube *Cube, _ bool) {
	Right(cube, true)
}

func RLeft(cube *Cube, _ bool) {
	Left(cube, true)
}

func RBack(cube *Cube, _ bool) {
	Back(cube, true)
}

func RUp(cube *Cube, _ bool) {
	Up(cube, true)
}

func RDown(cube *Cube, _ bool) {
	Down(cube, true)
}
