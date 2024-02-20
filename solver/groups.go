package solver

import (
	"github.com/AloySobek/Rubik/model"
)

var G0 map[string]func(*model.Cube) *model.Cube = map[string]func(*model.Cube) *model.Cube{
	"U":  U,
	"U2": func(c *model.Cube) *model.Cube { U(c); return U(c) },
	"U'": func(c *model.Cube) *model.Cube { U(c); U(c); return U(c) },

	"L":  L,
	"L2": func(c *model.Cube) *model.Cube { L(c); return L(c) },
	"L'": func(c *model.Cube) *model.Cube { L(c); L(c); return L(c) },

	"F":  F,
	"F2": func(c *model.Cube) *model.Cube { F(c); return F(c) },
	"F'": func(c *model.Cube) *model.Cube { F(c); F(c); return F(c) },

	"R":  R,
	"R2": func(c *model.Cube) *model.Cube { R(c); return R(c) },
	"R'": func(c *model.Cube) *model.Cube { R(c); R(c); return R(c) },

	"B":  B,
	"B2": func(c *model.Cube) *model.Cube { B(c); return B(c) },
	"B'": func(c *model.Cube) *model.Cube { B(c); B(c); return B(c) },

	"D":  D,
	"D2": func(c *model.Cube) *model.Cube { D(c); return D(c) },
	"D'": func(c *model.Cube) *model.Cube { D(c); D(c); return D(c) },
}

var G1 map[string]func(*model.Cube) *model.Cube = map[string]func(*model.Cube) *model.Cube{
	"U": U,

	"L":  L,
	"L2": func(c *model.Cube) *model.Cube { L(c); return L(c) },
	"L'": func(c *model.Cube) *model.Cube { L(c); L(c); return L(c) },

	"F":  F,
	"F2": func(c *model.Cube) *model.Cube { F(c); return F(c) },
	"F'": func(c *model.Cube) *model.Cube { F(c); F(c); return F(c) },

	"R":  R,
	"R2": func(c *model.Cube) *model.Cube { R(c); return R(c) },
	"R'": func(c *model.Cube) *model.Cube { R(c); R(c); return R(c) },

	"B":  B,
	"B2": func(c *model.Cube) *model.Cube { B(c); return B(c) },
	"B'": func(c *model.Cube) *model.Cube { B(c); B(c); return B(c) },

	"D": D,
}

var G2 map[string]func(*model.Cube) *model.Cube = map[string]func(*model.Cube) *model.Cube{
	"U": U,

	"L":  L,
	"L2": func(c *model.Cube) *model.Cube { L(c); return L(c) },
	"L'": func(c *model.Cube) *model.Cube { L(c); L(c); return L(c) },

	"F": F,

	"R":  R,
	"R2": func(c *model.Cube) *model.Cube { R(c); return R(c) },
	"R'": func(c *model.Cube) *model.Cube { R(c); R(c); return R(c) },

	"B": B,

	"D": D,
}

var G3 map[string]func(*model.Cube) *model.Cube = map[string]func(*model.Cube) *model.Cube{
	"U": U,

	"L": L,

	"F": F,

	"R": R,

	"B": B,

	"D": D,
}
