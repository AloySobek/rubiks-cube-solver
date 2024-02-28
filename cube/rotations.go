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

func U(c *Cube) *Cube {
	ur := c.EO[UR]

	c.EO[UR] = !c.EO[UB]
	c.EO[UB] = !c.EO[UL]
	c.EO[UL] = !c.EO[UF]
	c.EO[UF] = !ur

	ul := c.EP[UL]

	c.EP[UL] = c.EP[UF]
	c.EP[UF] = c.EP[UR]
	c.EP[UR] = c.EP[UB]
	c.EP[UB] = ul

	urf := c.CO[URF]

	c.CO[URF] = (c.CO[URB] + 1) % 3
	c.CO[URB] = (c.CO[ULB] + 2) % 3
	c.CO[ULB] = (c.CO[ULF] + 1) % 3
	c.CO[ULF] = (urf + 2) % 3

	urf2 := c.CP[URF]

	c.CP[URF] = c.CP[URB]
	c.CP[URB] = c.CP[ULB]
	c.CP[ULB] = c.CP[ULF]
	c.CP[ULF] = urf2

	r := c.R

	c.U = bits.RotateLeft64(c.U, 16)

	c.R = (c.R & c0 & c1 & c2) | (c.B & g0) | (c.B & g1) | (c.B & g2)
	c.B = (c.B & c0 & c1 & c2) | (c.L & g0) | (c.L & g1) | (c.L & g2)
	c.L = (c.L & c0 & c1 & c2) | (c.F & g0) | (c.F & g1) | (c.F & g2)
	c.F = (c.F & c0 & c1 & c2) | (r & g0) | (r & g1) | (r & g2)

	return c
}

func L(c *Cube) *Cube {
	lf := c.EO[LF]

	c.EO[LF] = c.EO[UL]
	c.EO[UL] = c.EO[LB]
	c.EO[LB] = c.EO[DL]
	c.EO[DL] = lf

	lb := c.EP[LB]

	c.EP[LB] = c.EP[DL]
	c.EP[DL] = c.EP[LF]
	c.EP[LF] = c.EP[UL]
	c.EP[UL] = lb

	dlf := c.CO[DLF]

	c.CO[DLF] = c.CO[ULF]
	c.CO[ULF] = c.CO[ULB]
	c.CO[ULB] = c.CO[DLB]
	c.CO[DLB] = dlf

	dlf2 := c.CP[DLF]

	c.CP[DLF] = c.CP[ULF]
	c.CP[ULF] = c.CP[ULB]
	c.CP[ULB] = c.CP[DLB]
	c.CP[DLB] = dlf2

	f := c.F

	c.L = bits.RotateLeft64(c.L, 16)

	c.F = (c.F & c0 & c6 & c7) | (c.U & g0) | (c.U & g6) | (c.U & g7)
	c.U = (c.U & c0 & c6 & c7) | bits.RotateLeft64((c.B&g2)|(c.B&g3)|(c.B&g4), -32)
	c.B = (c.B & c2 & c3 & c4) | bits.RotateLeft64((c.D&g0)|(c.D&g6)|(c.D&g7), 32)
	c.D = (c.D & c0 & c6 & c7) | (f & g0) | (f & g6) | (f & g7)

	return c
}

func F(c *Cube) *Cube {
	rf := c.EO[RF]

	c.EO[RF] = c.EO[UF]
	c.EO[UF] = c.EO[LF]
	c.EO[LF] = c.EO[DF]
	c.EO[DF] = rf

	lf := c.EP[LF]

	c.EP[LF] = c.EP[DF]
	c.EP[DF] = c.EP[RF]
	c.EP[RF] = c.EP[UF]
	c.EP[UF] = lf

	drf := c.CO[DRF]

	c.CO[DRF] = (c.CO[URF] + 1) % 3
	c.CO[URF] = (c.CO[ULF] + 2) % 3
	c.CO[ULF] = (c.CO[DLF] + 1) % 3
	c.CO[DLF] = (drf + 2) % 3

	drf2 := c.CP[DRF]

	c.CP[DRF] = c.CP[URF]
	c.CP[URF] = c.CP[ULF]
	c.CP[ULF] = c.CP[DLF]
	c.CP[DLF] = drf2

	r := c.R

	c.F = bits.RotateLeft64(c.F, 16)

	c.R = (c.R & c0 & c6 & c7) | bits.RotateLeft64((c.U&g4)|(c.U&g5)|(c.U&g6), 16)
	c.U = (c.U & c4 & c5 & c6) | bits.RotateLeft64((c.L&g2)|(c.L&g3)|(c.L&g4), 16)
	c.L = (c.L & c2 & c3 & c4) | bits.RotateLeft64((c.D&g0)|(c.D&g1)|(c.D&g2), 16)
	c.D = (c.D & c0 & c1 & c2) | bits.RotateLeft64((r&g0)|(r&g6)|(r&g7), 16)

	return c
}

func R(c *Cube) *Cube {
	rb := c.EO[RB]

	c.EO[RB] = c.EO[UR]
	c.EO[UR] = c.EO[RF]
	c.EO[RF] = c.EO[DR]
	c.EO[DR] = rb

	rf := c.EP[RF]

	c.EP[RF] = c.EP[DR]
	c.EP[DR] = c.EP[RB]
	c.EP[RB] = c.EP[UR]
	c.EP[UR] = rf

	drb := c.CO[DRB]

	c.CO[DRB] = c.CO[URB]
	c.CO[URB] = c.CO[URF]
	c.CO[URF] = c.CO[DRF]
	c.CO[DRF] = drb

	drb2 := c.CP[DRB]

	c.CP[DRB] = c.CP[URB]
	c.CP[URB] = c.CP[URF]
	c.CP[URF] = c.CP[DRF]
	c.CP[DRF] = drb2

	b := c.B

	c.R = bits.RotateLeft64(c.R, 16)

	c.B = (c.B & c0 & c6 & c7) | bits.RotateLeft64((c.U&g2)|(c.U&g3)|(c.U&g4), 32)
	c.U = (c.U & c2 & c3 & c4) | (c.F & g2) | (c.F & g3) | (c.F & g4)
	c.F = (c.F & c2 & c3 & c4) | (c.D & g2) | (c.D & g3) | (c.D & g4)
	c.D = (c.D & c2 & c3 & c4) | bits.RotateLeft64((b&g0)|(b&g6)|(b&g7), 32)

	return c
}

func B(c *Cube) *Cube {
	lb := c.EO[LB]

	c.EO[LB] = c.EO[UB]
	c.EO[UB] = c.EO[RB]
	c.EO[RB] = c.EO[DB]
	c.EO[DB] = lb

	rb := c.EP[RB]

	c.EP[RB] = c.EP[DB]
	c.EP[DB] = c.EP[LB]
	c.EP[LB] = c.EP[UB]
	c.EP[UB] = rb

	dlb := c.CO[DLB]

	c.CO[DLB] = (c.CO[ULB] + 1) % 3
	c.CO[ULB] = (c.CO[URB] + 2) % 3
	c.CO[URB] = (c.CO[DRB] + 1) % 3
	c.CO[DRB] = (dlb + 2) % 3

	dlb2 := c.CP[DLB]

	c.CP[DLB] = c.CP[ULB]
	c.CP[ULB] = c.CP[URB]
	c.CP[URB] = c.CP[DRB]
	c.CP[DRB] = dlb2

	l := c.L

	c.B = bits.RotateLeft64(c.B, 16)

	c.L = (c.L & c0 & c6 & c7) | bits.RotateLeft64((c.U&g0)|(c.U&g1)|(c.U&g2), -16)
	c.U = (c.U & c0 & c1 & c2) | bits.RotateLeft64((c.R&g2)|(c.R&g3)|(c.R&g4), -16)
	c.R = (c.R & c2 & c3 & c4) | bits.RotateLeft64((c.D&g4)|(c.D&g5)|(c.D&g6), -16)
	c.D = (c.D & c4 & c5 & c6) | bits.RotateLeft64((l&g0)|(l&g6)|(l&g7), -16)

	return c
}

func D(c *Cube) *Cube {
	dr := c.EO[DR]

	c.EO[DR] = !c.EO[DF]
	c.EO[DF] = !c.EO[DL]
	c.EO[DL] = !c.EO[DB]
	c.EO[DB] = !dr

	dl := c.EP[DL]

	c.EP[DL] = c.EP[DB]
	c.EP[DB] = c.EP[DR]
	c.EP[DR] = c.EP[DF]
	c.EP[DF] = dl

	drb := c.CO[DRB]

	c.CO[DRB] = (c.CO[DRF] + 1) % 3
	c.CO[DRF] = (c.CO[DLF] + 2) % 3
	c.CO[DLF] = (c.CO[DLB] + 1) % 3
	c.CO[DLB] = (drb + 2) % 3

	drb2 := c.CP[DRB]

	c.CP[DRB] = c.CP[DRF]
	c.CP[DRF] = c.CP[DLF]
	c.CP[DLF] = c.CP[DLB]
	c.CP[DLB] = drb2

	r := c.R

	c.D = bits.RotateLeft64(c.D, 16)

	c.R = (c.R & c4 & c5 & c6) | (c.F & g4) | (c.F & g5) | (c.F & g6)
	c.F = (c.F & c4 & c5 & c6) | (c.L & g4) | (c.L & g5) | (c.L & g6)
	c.L = (c.L & c4 & c5 & c6) | (c.B & g4) | (c.B & g5) | (c.B & g6)
	c.B = (c.B & c4 & c5 & c6) | (r & g4) | (r & g5) | (r & g6)

	return c
}
