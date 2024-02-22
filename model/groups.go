package model

var G0 map[string]func(*Cube) *Cube = map[string]func(*Cube) *Cube{
	"U":  U,
	"U2": func(c *Cube) *Cube { U(c); return U(c) },
	"U'": func(c *Cube) *Cube { U(c); U(c); return U(c) },

	"L":  L,
	"L2": func(c *Cube) *Cube { L(c); return L(c) },
	"L'": func(c *Cube) *Cube { L(c); L(c); return L(c) },

	"F":  F,
	"F2": func(c *Cube) *Cube { F(c); return F(c) },
	"F'": func(c *Cube) *Cube { F(c); F(c); return F(c) },

	"R":  R,
	"R2": func(c *Cube) *Cube { R(c); return R(c) },
	"R'": func(c *Cube) *Cube { R(c); R(c); return R(c) },

	"B":  B,
	"B2": func(c *Cube) *Cube { B(c); return B(c) },
	"B'": func(c *Cube) *Cube { B(c); B(c); return B(c) },

	"D":  D,
	"D2": func(c *Cube) *Cube { D(c); return D(c) },
	"D'": func(c *Cube) *Cube { D(c); D(c); return D(c) },
}

var G1 map[string]func(*Cube) *Cube = map[string]func(*Cube) *Cube{
	"U2": func(c *Cube) *Cube { U(c); return U(c) },

	"L":  L,
	"L2": func(c *Cube) *Cube { L(c); return L(c) },
	"L'": func(c *Cube) *Cube { L(c); L(c); return L(c) },

	"F":  F,
	"F2": func(c *Cube) *Cube { F(c); return F(c) },
	"F'": func(c *Cube) *Cube { F(c); F(c); return F(c) },

	"R":  R,
	"R2": func(c *Cube) *Cube { R(c); return R(c) },
	"R'": func(c *Cube) *Cube { R(c); R(c); return R(c) },

	"B":  B,
	"B2": func(c *Cube) *Cube { B(c); return B(c) },
	"B'": func(c *Cube) *Cube { B(c); B(c); return B(c) },

	"D2": func(c *Cube) *Cube { D(c); return D(c) },
}

var G2 map[string]func(*Cube) *Cube = map[string]func(*Cube) *Cube{
	"U2": func(c *Cube) *Cube { U(c); return U(c) },

	"L":  L,
	"L2": func(c *Cube) *Cube { L(c); return L(c) },
	"L'": func(c *Cube) *Cube { L(c); L(c); return L(c) },

	"F2": func(c *Cube) *Cube { F(c); return F(c) },

	"R":  R,
	"R2": func(c *Cube) *Cube { R(c); return R(c) },
	"R'": func(c *Cube) *Cube { R(c); R(c); return R(c) },

	"B2": func(c *Cube) *Cube { B(c); return B(c) },

	"D2": func(c *Cube) *Cube { D(c); return D(c) },
}

var G3 map[string]func(*Cube) *Cube = map[string]func(*Cube) *Cube{
	"U2": func(c *Cube) *Cube { U(c); return U(c) },

	"L2": func(c *Cube) *Cube { L(c); return L(c) },

	"F2": func(c *Cube) *Cube { F(c); return F(c) },

	"R2": func(c *Cube) *Cube { R(c); return R(c) },

	"B2": func(c *Cube) *Cube { B(c); return B(c) },

	"D2": func(c *Cube) *Cube { D(c); return D(c) },
}
