package model

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

	urb := c.CP[URB]

	c.CP[URB] = c.CP[ULB]
	c.CP[ULB] = c.CP[ULF]
	c.CP[ULF] = c.CP[URF]
	c.CP[URF] = urb

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

	ulf := c.CP[ULF]

	c.CP[ULF] = c.CP[ULB]
	c.CP[ULB] = c.CP[DLB]
	c.CP[DLB] = c.CP[DLF]
	c.CP[DLF] = ulf

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

	urf := c.CP[URF]

	c.CP[URF] = c.CP[ULF]
	c.CP[ULF] = c.CP[DLF]
	c.CP[DLF] = c.CP[DRF]
	c.CP[DRF] = urf

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

	urb := c.CP[URB]

	c.CP[URB] = c.CP[URF]
	c.CP[URF] = c.CP[DRF]
	c.CP[DRF] = c.CP[DRB]
	c.CP[DRF] = urb

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

	ulb := c.CP[ULB]

	c.CP[ULB] = c.CP[URB]
	c.CP[URB] = c.CP[DRB]
	c.CP[DRB] = c.CP[DLB]
	c.CP[DLB] = ulb

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

	drf := c.CP[DRF]

	c.CP[DRF] = c.CP[DLF]
	c.CP[DLF] = c.CP[DLB]
	c.CP[DLB] = c.CP[DRB]
	c.CP[DRB] = drf

	return c
}
