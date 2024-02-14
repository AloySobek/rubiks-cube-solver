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

func RotateS0(c *Cube) *Cube {
	s3 := c.S3

	c.S0 = bits.RotateLeft64(c.S0, 16)

	c.S3 = (c.S3 & c0 & c1 & c2) | (c.S4 & g0) | (c.S4 & g1) | (c.S4 & g2)
	c.S4 = (c.S4 & c0 & c1 & c2) | (c.S1 & g0) | (c.S1 & g1) | (c.S1 & g2)
	c.S1 = (c.S1 & c0 & c1 & c2) | (c.S2 & g0) | (c.S2 & g1) | (c.S2 & g2)
	c.S2 = (c.S2 & c0 & c1 & c2) | (s3 & g0) | (s3 & g1) | (s3 & g2)

	// Reverse
	// c.S0 = bits.RotateLeft64(c.S0, -16)

	// c.S3 = (c.S3 & c0 & c1 & c2) | (c.S2 & g0) | (c.S2 & g1) | (c.S2 & g2)
	// c.S2 = (c.S2 & c0 & c1 & c2) | (c.S1 & g0) | (c.S1 & g1) | (c.S1 & g2)
	// c.S1 = (c.S1 & c0 & c1 & c2) | (c.S4 & g0) | (c.S4 & g1) | (c.S4 & g2)
	// c.S4 = (c.S4 & c0 & c1 & c2) | (s3 & g0) | (s3 & g1) | (s3 & g2)

	return c
}

func RotateS1(c *Cube) *Cube {
	s2 := c.S2

	c.S1 = bits.RotateLeft64(c.S1, 16)

	c.S2 = (c.S2 & c0 & c6 & c7) | (c.S0 & g0) | (c.S0 & g6) | (c.S0 & g7)
	c.S0 = (c.S0 & c0 & c6 & c7) | bits.RotateLeft64((c.S4&g2)|(c.S4&g3)|(c.S4&g4), -32)
	c.S4 = (c.S4 & c2 & c3 & c4) | bits.RotateLeft64((c.S5&g0)|(c.S5&g6)|(c.S5&g7), 32)
	c.S5 = (c.S5 & c0 & c6 & c7) | (s2 & g0) | (s2 & g6) | (s2 & g7)

	// Reverse
	// c.S1 = bits.RotateLeft64(c.S1, -16)

	// c.S2 = (c.S2 & c0 & c6 & c7) | (c.S5 & g0) | (c.S5 & g6) | (c.S5 & g7)
	// c.S5 = (c.S5 & c0 & c6 & c7) | bits.RotateLeft64((c.S4&g2)|(c.S4&g3)|(c.S4&g4), -32)
	// c.S4 = (c.S4 & c2 & c3 & c4) | bits.RotateLeft64((c.S0&g0)|(c.S0&g6)|(c.S0&g7), 32)
	// c.S0 = (c.S0 & c0 & c6 & c7) | (s2 & g0) | (s2 & g6) | (s2 & g7)

	return c
}

func RotateS2(c *Cube) *Cube {
	s3 := c.S3

	c.S2 = bits.RotateLeft64(c.S2, 16)

	c.S3 = (c.S3 & c0 & c6 & c7) | bits.RotateLeft64((c.S0&g4)|(c.S0&g5)|(c.S0&g6), 16)
	c.S0 = (c.S0 & c4 & c5 & c6) | bits.RotateLeft64((c.S1&g2)|(c.S1&g3)|(c.S1&g4), 16)
	c.S1 = (c.S1 & c2 & c3 & c4) | bits.RotateLeft64((c.S5&g0)|(c.S5&g1)|(c.S5&g2), 16)
	c.S5 = (c.S5 & c0 & c1 & c2) | bits.RotateLeft64((s3&g0)|(s3&g6)|(s3&g7), 16)

	// Reverse
	// c.S2 = bits.RotateLeft64(c.S2, -16)

	// c.S3 = (c.S3 & c0 & c6 & c7) | bits.RotateLeft64((c.S5&g0)|(c.S5&g1)|(c.S5&g2), -16)
	// c.S5 = (c.S5 & c0 & c1 & c2) | bits.RotateLeft64((c.S1&g2)|(c.S1&g3)|(c.S1&g4), -16)
	// c.S1 = (c.S1 & c2 & c3 & c4) | bits.RotateLeft64((c.S0&g4)|(c.S0&g5)|(c.S0&g6), -16)
	// c.S0 = (c.S0 & c4 & c5 & c6) | bits.RotateLeft64((s3&g0)|(s3&g6)|(s3&g7), -16)

	return c
}

func RotateS3(c *Cube) *Cube {
	s4 := c.S4

	c.S3 = bits.RotateLeft64(c.S3, 16)

	c.S4 = (c.S4 & c0 & c6 & c7) | bits.RotateLeft64((c.S0&g2)|(c.S0&g3)|(c.S0&g4), 32)
	c.S0 = (c.S0 & c2 & c3 & c4) | (c.S2 & g2) | (c.S2 & g3) | (c.S2 & g4)
	c.S2 = (c.S2 & c2 & c3 & c4) | (c.S5 & g2) | (c.S5 & g3) | (c.S5 & g4)
	c.S5 = (c.S5 & c2 & c3 & c4) | bits.RotateLeft64((s4&g0)|(s4&g6)|(s4&g7), 32)

	// Reverse
	// c.S3 = bits.RotateLeft64(c.S3, -16)

	// c.S4 = (c.S4 & c0 & c6 & c7) | bits.RotateLeft64((c.S5&g2)|(c.S5&g3)|(c.S5&g4), 32)
	// c.S5 = (c.S5 & c2 & c3 & c4) | (c.S2 & g2) | (c.S2 & g3) | (c.S2 & g4)
	// c.S2 = (c.S2 & c2 & c3 & c4) | (c.S0 & g2) | (c.S0 & g3) | (c.S0 & g4)
	// c.S0 = (c.S0 & c2 & c3 & c4) | bits.RotateLeft64((s4&g0)|(s4&g6)|(s4&g7), 32)

	return c
}

func RotateS4(c *Cube) *Cube {
	s1 := c.S1

	c.S4 = bits.RotateLeft64(c.S4, 16)

	c.S1 = (c.S1 & c0 & c6 & c7) | bits.RotateLeft64((c.S0&g0)|(c.S0&g1)|(c.S0&g2), -16)
	c.S0 = (c.S0 & c0 & c1 & c2) | bits.RotateLeft64((c.S3&g2)|(c.S3&g3)|(c.S3&g4), -16)
	c.S3 = (c.S3 & c2 & c3 & c4) | bits.RotateLeft64((c.S5&g4)|(c.S5&g5)|(c.S5&g6), -16)
	c.S5 = (c.S5 & c4 & c5 & c6) | bits.RotateLeft64((s1&g0)|(s1&g6)|(s1&g7), -16)

	// Reverse
	// c.S4 = bits.RotateLeft64(c.S4, -16)

	// c.S1 = (c.S1 & c0 & c6 & c7) | bits.RotateLeft64((c.S5&g4)|(c.S5&g5)|(c.S5&g6), 16)
	// c.S5 = (c.S5 & c4 & c5 & c6) | bits.RotateLeft64((c.S3&g2)|(c.S3&g3)|(c.S3&g4), 16)
	// c.S3 = (c.S3 & c2 & c3 & c4) | bits.RotateLeft64((c.S0&g0)|(c.S0&g1)|(c.S0&g2), 16)
	// c.S0 = (c.S0 & c0 & c1 & c2) | bits.RotateLeft64((s1&g0)|(s1&g6)|(s1&g7), 16)

	return c
}

func RotateS5(c *Cube) *Cube {
	s3 := c.S3

	c.S5 = bits.RotateLeft64(c.S5, 16)

	c.S3 = (c.S3 & c4 & c5 & c6) | (c.S2 & g4) | (c.S2 & g5) | (c.S2 & g6)
	c.S2 = (c.S2 & c4 & c5 & c6) | (c.S1 & g4) | (c.S1 & g5) | (c.S1 & g6)
	c.S1 = (c.S1 & c4 & c5 & c6) | (c.S4 & g4) | (c.S4 & g5) | (c.S4 & g6)
	c.S4 = (c.S4 & c4 & c5 & c6) | (s3 & g4) | (s3 & g5) | (s3 & g6)

	// Reverse
	// c.S5 = bits.RotateLeft64(c.S5, -16)

	// c.S3 = (c.S3 & c4 & c5 & c6) | (c.S4 & g4) | (c.S4 & g5) | (c.S4 & g6)
	// c.S4 = (c.S4 & c4 & c5 & c6) | (c.S1 & g4) | (c.S1 & g5) | (c.S1 & g6)
	// c.S1 = (c.S1 & c4 & c5 & c6) | (c.S2 & g4) | (c.S2 & g5) | (c.S2 & g6)
	// c.S2 = (c.S2 & c4 & c5 & c6) | (s3 & g4) | (s3 & g5) | (s3 & g6)

	return c
}

func RotateS0Twice(c *Cube) *Cube {
	RotateS0(c)
	RotateS0(c)

	return c
}

func RotateS1Twice(c *Cube) *Cube {
	RotateS1(c)
	RotateS1(c)

	return c
}

func RotateS2Twice(c *Cube) *Cube {
	RotateS2(c)
	RotateS2(c)

	return c
}

func RotateS3Twice(c *Cube) *Cube {
	RotateS3(c)
	RotateS3(c)

	return c
}

func RotateS4Twice(c *Cube) *Cube {
	RotateS4(c)
	RotateS4(c)

	return c
}

func RotateS5Twice(c *Cube) *Cube {
	RotateS5(c)
	RotateS5(c)

	return c
}

func RotateS0Thrice(c *Cube) *Cube {
	RotateS0Twice(c)
	RotateS0(c)

	return c
}

func RotateS1Thrice(c *Cube) *Cube {
	RotateS1Twice(c)
	RotateS1(c)

	return c
}

func RotateS2Thrice(c *Cube) *Cube {
	RotateS2Twice(c)
	RotateS2(c)

	return c
}

func RotateS3Thrice(c *Cube) *Cube {
	RotateS3Twice(c)
	RotateS3(c)

	return c
}

func RotateS4Thrice(c *Cube) *Cube {
	RotateS4Twice(c)
	RotateS4(c)

	return c
}

func RotateS5Thrice(c *Cube) *Cube {
	RotateS5Twice(c)
	RotateS5(c)

	return c
}
