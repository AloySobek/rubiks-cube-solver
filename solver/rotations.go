package solver

import (
	"github.com/AloySobek/Rubik/model"
)

func U(c *model.Cube) *model.Cube {
	ur := c.EO[model.UR]

	c.EO[model.UR] = !c.EO[model.UB]
	c.EO[model.UB] = !c.EO[model.UL]
	c.EO[model.UL] = !c.EO[model.UF]
	c.EO[model.UF] = !ur

	ul := c.EP[model.UL]

	c.EP[model.UL] = c.EP[model.UF]
	c.EP[model.UF] = c.EP[model.UR]
	c.EP[model.UR] = c.EP[model.UB]
	c.EP[model.UB] = ul

	urf := c.CO[model.URF]

	c.CO[model.URF] = (1 + c.CO[model.URB]) % 3
	c.CO[model.URB] = (1 + c.CO[model.ULB]) % 3
	c.CO[model.ULB] = (1 + c.CO[model.ULF]) % 3
	c.CO[model.ULF] = (1 + urf) % 3

	return c
}

func L(c *model.Cube) *model.Cube {
	lf := c.EO[model.LF]

	c.EO[model.LF] = c.EO[model.UL]
	c.EO[model.UL] = c.EO[model.LB]
	c.EO[model.LB] = c.EO[model.DL]
	c.EO[model.DL] = lf

	lb := c.EP[model.LB]

	c.EP[model.LB] = c.EP[model.DL]
	c.EP[model.DL] = c.EP[model.LF]
	c.EP[model.LF] = c.EP[model.UL]
	c.EP[model.UL] = lb

	ulb := c.CO[model.ULB]

	c.CO[model.ULB] = c.CO[model.DLB]
	c.CO[model.DLB] = c.CO[model.DLF]
	c.CO[model.DLF] = c.CO[model.ULF]
	c.CO[model.ULF] = ulb

	return c
}

func F(c *model.Cube) *model.Cube {
	rf := c.EO[model.RF]

	c.EO[model.RF] = c.EO[model.UF]
	c.EO[model.UF] = c.EO[model.LF]
	c.EO[model.LF] = c.EO[model.DF]
	c.EO[model.DF] = rf

	lf := c.EP[model.LF]

	c.EP[model.LF] = c.EP[model.DF]
	c.EP[model.DF] = c.EP[model.RF]
	c.EP[model.RF] = c.EP[model.UF]
	c.EP[model.UF] = lf

	drf := c.CO[model.DRF]

	c.CO[model.DRF] = (1 + c.CO[model.URF]) % 3
	c.CO[model.URF] = (1 + c.CO[model.ULF]) % 3
	c.CO[model.ULF] = (1 + c.CO[model.DLF]) % 3
	c.CO[model.DLF] = (1 + drf) % 3

	return c
}

func R(c *model.Cube) *model.Cube {
	rb := c.EO[model.RB]

	c.EO[model.RB] = c.EO[model.UR]
	c.EO[model.UR] = c.EO[model.RF]
	c.EO[model.RF] = c.EO[model.DR]
	c.EO[model.DR] = rb

	rf := c.EP[model.RF]

	c.EP[model.RF] = c.EP[model.DR]
	c.EP[model.DR] = c.EP[model.RB]
	c.EP[model.RB] = c.EP[model.UR]
	c.EP[model.UR] = rf

	urf := c.CO[model.URF]

	c.CO[model.URF] = c.CO[model.DRF]
	c.CO[model.DRF] = c.CO[model.DRB]
	c.CO[model.DRB] = c.CO[model.URB]
	c.CO[model.URB] = urf

	return c
}

func B(c *model.Cube) *model.Cube {
	lb := c.EO[model.LB]

	c.EO[model.LB] = c.EO[model.UB]
	c.EO[model.UB] = c.EO[model.RB]
	c.EO[model.RB] = c.EO[model.DB]
	c.EO[model.DB] = lb

	rb := c.EP[model.RB]

	c.EP[model.RB] = c.EP[model.DB]
	c.EP[model.DB] = c.EP[model.LB]
	c.EP[model.LB] = c.EP[model.UB]
	c.EP[model.UB] = rb

	dlb := c.CO[model.DLB]

	c.CO[model.DLB] = (1 + c.CO[model.ULB]) % 3
	c.CO[model.ULB] = (1 + c.CO[model.URB]) % 3
	c.CO[model.URB] = (1 + c.CO[model.DRB]) % 3
	c.CO[model.DRB] = (1 + dlb) % 3

	return c
}

func D(c *model.Cube) *model.Cube {
	dr := c.EO[model.DR]

	c.EO[model.DR] = !c.EO[model.DF]
	c.EO[model.DF] = !c.EO[model.DL]
	c.EO[model.DL] = !c.EO[model.DB]
	c.EO[model.DB] = !dr

	dl := c.EP[model.DL]

	c.EP[model.DL] = c.EP[model.DB]
	c.EP[model.DB] = c.EP[model.DR]
	c.EP[model.DR] = c.EP[model.DF]
	c.EP[model.DF] = dl

	drb := c.CO[model.DRB]

	c.CO[model.DRB] = (1 + c.CO[model.DRF]) % 3
	c.CO[model.DRF] = (1 + c.CO[model.DLF]) % 3
	c.CO[model.DLF] = (1 + c.CO[model.DLB]) % 3
	c.CO[model.DLB] = (1 + drb) % 3

	return c
}
