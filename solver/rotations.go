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

	return c
}

func L(c *model.Cube) *model.Cube {
	lf := c.EO[model.LF]

	c.EO[model.LF] = c.EO[model.UL]
	c.EO[model.UL] = c.EO[model.LB]
	c.EO[model.LB] = c.EO[model.DL]
	c.EO[model.DL] = lf

	return c
}

func F(c *model.Cube) *model.Cube {
	rf := c.EO[model.RF]

	c.EO[model.RF] = c.EO[model.UF]
	c.EO[model.UF] = c.EO[model.LF]
	c.EO[model.LF] = c.EO[model.DF]
	c.EO[model.DF] = rf

	return c
}

func R(c *model.Cube) *model.Cube {
	rb := c.EO[model.RB]

	c.EO[model.RB] = c.EO[model.UR]
	c.EO[model.UR] = c.EO[model.RF]
	c.EO[model.RF] = c.EO[model.DR]
	c.EO[model.DR] = rb

	return c
}

func B(c *model.Cube) *model.Cube {
	lb := c.EO[model.LB]

	c.EO[model.LB] = c.EO[model.UB]
	c.EO[model.UB] = c.EO[model.RB]
	c.EO[model.RB] = c.EO[model.DB]
	c.EO[model.DB] = lb

	return c
}

func D(c *model.Cube) *model.Cube {
	dr := c.EO[model.DR]

	c.EO[model.DR] = !c.EO[model.DF]
	c.EO[model.DF] = !c.EO[model.DL]
	c.EO[model.DL] = !c.EO[model.DB]
	c.EO[model.DB] = !dr

	return c
}
