package cube

import "math/bits"

const (
	// Clear
	c0 = 0xFFFFFFFFFFFFFF00
	c1 = 0xFFFFFFFFFFFF00FF
	c2 = 0xFFFFFFFFFF00FFFF
	c3 = 0xFFFFFFFF00FFFFFF
	c4 = 0xFFFFFF00FFFFFFFF
	c5 = 0xFFFF00FFFFFFFFFF
	c6 = 0xFF00FFFFFFFFFFFF
	c7 = 0x00FFFFFFFFFFFFFF

	// Get
	g0 = 0x00000000000000FF
	g1 = 0x000000000000FF00
	g2 = 0x0000000000FF0000
	g3 = 0x00000000FF000000
	g4 = 0x000000FF00000000
	g5 = 0x0000FF0000000000
	g6 = 0x00FF000000000000
	g7 = 0xFF00000000000000
)

var G0 map[string]func(*Cube) *Cube = map[string]func(*Cube) *Cube{
	"U":  Up,
	"U2": Up2,
	"U'": UpPrime,

	"L":  Left,
	"L2": Left2,
	"L'": LeftPrime,

	"F":  Front,
	"F2": Front2,
	"F'": FrontPrime,

	"R":  Right,
	"R2": Right2,
	"R'": RightPrime,

	"B":  Back,
	"B2": Back2,
	"B'": BackPrime,

	"D":  Down,
	"D2": Down2,
	"D'": DownPrime,
}

var G1 map[string]func(*Cube) *Cube = map[string]func(*Cube) *Cube{
	"U2": Up2,

	"L":  Left,
	"L2": Left2,
	"L'": LeftPrime,

	"F":  Front,
	"F2": Front2,
	"F'": FrontPrime,

	"R":  Right,
	"R2": Right2,
	"R'": RightPrime,

	"B":  Back,
	"B2": Back2,
	"B'": BackPrime,

	"D2": Down2,
}

var G2 map[string]func(*Cube) *Cube = map[string]func(*Cube) *Cube{
	"U2": Up2,

	"L":  Left,
	"L2": Left2,
	"L'": LeftPrime,

	"F2": Front2,

	"R":  Right,
	"R2": Right2,
	"R'": RightPrime,

	"B2": Back2,

	"D2": Down2,
}

var G3 map[string]func(*Cube) *Cube = map[string]func(*Cube) *Cube{
	"U2": Up2,

	"L2": Left2,

	"F2": Front2,

	"R2": Right2,

	"B2": Back2,

	"D2": Down2,
}

func Up(c *Cube) *Cube {
	r := c.R

	c.U = bits.RotateLeft64(c.U, 16)

	c.R = (c.R & c0 & c1 & c2) | (c.B & g0) | (c.B & g1) | (c.B & g2)
	c.B = (c.B & c0 & c1 & c2) | (c.L & g0) | (c.L & g1) | (c.L & g2)
	c.L = (c.L & c0 & c1 & c2) | (c.F & g0) | (c.F & g1) | (c.F & g2)
	c.F = (c.F & c0 & c1 & c2) | (r & g0) | (r & g1) | (r & g2)

	return c
}

func Left(c *Cube) *Cube {
	f := c.F

	c.L = bits.RotateLeft64(c.L, 16)

	c.F = (c.F & c0 & c6 & c7) | (c.U & g0) | (c.U & g6) | (c.U & g7)
	c.U = (c.U & c0 & c6 & c7) | bits.RotateLeft64((c.B&g2)|(c.B&g3)|(c.B&g4), -32)
	c.B = (c.B & c2 & c3 & c4) | bits.RotateLeft64((c.D&g0)|(c.D&g6)|(c.D&g7), 32)
	c.D = (c.D & c0 & c6 & c7) | (f & g0) | (f & g6) | (f & g7)

	return c
}

func Front(c *Cube) *Cube {
	r := c.R

	c.F = bits.RotateLeft64(c.F, 16)

	c.R = (c.R & c0 & c6 & c7) | bits.RotateLeft64((c.U&g4)|(c.U&g5)|(c.U&g6), 16)
	c.U = (c.U & c4 & c5 & c6) | bits.RotateLeft64((c.L&g2)|(c.L&g3)|(c.L&g4), 16)
	c.L = (c.L & c2 & c3 & c4) | bits.RotateLeft64((c.D&g0)|(c.D&g1)|(c.D&g2), 16)
	c.D = (c.D & c0 & c1 & c2) | bits.RotateLeft64((r&g0)|(r&g6)|(r&g7), 16)

	return c
}

func Right(c *Cube) *Cube {
	b := c.B

	c.R = bits.RotateLeft64(c.R, 16)

	c.B = (c.B & c0 & c6 & c7) | bits.RotateLeft64((c.U&g2)|(c.U&g3)|(c.U&g4), 32)
	c.U = (c.U & c2 & c3 & c4) | (c.F & g2) | (c.F & g3) | (c.F & g4)
	c.F = (c.F & c2 & c3 & c4) | (c.D & g2) | (c.D & g3) | (c.D & g4)
	c.D = (c.D & c2 & c3 & c4) | bits.RotateLeft64((b&g0)|(b&g6)|(b&g7), 32)

	return c
}

func Back(c *Cube) *Cube {
	l := c.L

	c.B = bits.RotateLeft64(c.B, 16)

	c.L = (c.L & c0 & c6 & c7) | bits.RotateLeft64((c.U&g0)|(c.U&g1)|(c.U&g2), -16)
	c.U = (c.U & c0 & c1 & c2) | bits.RotateLeft64((c.R&g2)|(c.R&g3)|(c.R&g4), -16)
	c.R = (c.R & c2 & c3 & c4) | bits.RotateLeft64((c.D&g4)|(c.D&g5)|(c.D&g6), -16)
	c.D = (c.D & c4 & c5 & c6) | bits.RotateLeft64((l&g0)|(l&g6)|(l&g7), -16)

	return c
}

func Down(c *Cube) *Cube {
	r := c.R

	c.D = bits.RotateLeft64(c.D, 16)

	c.R = (c.R & c4 & c5 & c6) | (c.F & g4) | (c.F & g5) | (c.F & g6)
	c.F = (c.F & c4 & c5 & c6) | (c.L & g4) | (c.L & g5) | (c.L & g6)
	c.L = (c.L & c4 & c5 & c6) | (c.B & g4) | (c.B & g5) | (c.B & g6)
	c.B = (c.B & c4 & c5 & c6) | (r & g4) | (r & g5) | (r & g6)

	return c
}

func Up2(c *Cube) *Cube {
	Up(c)
	Up(c)

	return c
}

func Left2(c *Cube) *Cube {
	Left(c)
	Left(c)

	return c
}

func Front2(c *Cube) *Cube {
	Front(c)
	Front(c)

	return c
}

func Right2(c *Cube) *Cube {
	Right(c)
	Right(c)

	return c
}

func Back2(c *Cube) *Cube {
	Back(c)
	Back(c)

	return c
}

func Down2(c *Cube) *Cube {
	Down(c)
	Down(c)

	return c
}

func UpPrime(c *Cube) *Cube {
	Up2(c)
	Up(c)

	return c
}

func LeftPrime(c *Cube) *Cube {
	Left2(c)
	Left(c)

	return c
}

func FrontPrime(c *Cube) *Cube {
	Front2(c)
	Front(c)

	return c
}

func RightPrime(c *Cube) *Cube {
	Right2(c)
	Right(c)

	return c
}

func BackPrime(c *Cube) *Cube {
	Back2(c)
	Back(c)

	return c
}

func DownPrime(c *Cube) *Cube {
	Down2(c)
	Down(c)

	return c
}
