package cube

import (
	"log"
	"math/rand"
	"strings"
)

type Cube struct {
	Yellow uint64
	Orange uint64
	Green  uint64
	White  uint64
	Blue   uint64
	Red    uint64
}

const (
	Yellow = 1
	Orange = 2
	Green  = 4
	White  = 8
	Blue   = 16
	Red    = 32
)

const (
	SolvedYellow = uint64(0) | Yellow | (Yellow << 8) | (Yellow << 16) | (Yellow << 24) | (Yellow << 32) | (Yellow << 40) | (Yellow << 48) | (Yellow << 56)
	SolvedOrange = uint64(0) | Orange | (Orange << 8) | (Orange << 16) | (Orange << 24) | (Orange << 32) | (Orange << 40) | (Orange << 48) | (Orange << 56)
	SolvedGreen  = uint64(0) | Green | (Green << 8) | (Green << 16) | (Green << 24) | (Green << 32) | (Green << 40) | (Green << 48) | (Green << 56)
	SolvedWhite  = uint64(0) | White | (White << 8) | (White << 16) | (White << 24) | (White << 32) | (White << 40) | (White << 48) | (White << 56)
	SolvedBlue   = uint64(0) | Blue | (Blue << 8) | (Blue << 16) | (Blue << 24) | (Blue << 32) | (Blue << 40) | (Blue << 48) | (Blue << 56)
	SolvedRed    = uint64(0) | Red | (Red << 8) | (Red << 16) | (Red << 24) | (Red << 32) | (Red << 40) | (Red << 48) | (Red << 56)
)

var PossibleMoves []string = []string{
	"F", "R", "L", "B", "U", "D",
	"F'", "R'", "L'", "B'", "U'", "D'",
	"F2", "R2", "L2", "B2", "U2", "D2",
}

var NotationMatrix map[string]func(*Cube, bool) = map[string]func(*Cube, bool){
	"F":  Front,
	"B":  Back,
	"R":  Right,
	"L":  Left,
	"U":  Up,
	"D":  Down,
	"F'": RFront,
	"B'": RBack,
	"R'": RRight,
	"L'": RLeft,
	"U'": RUp,
	"D'": RDown,
	"F2": func(c *Cube, _ bool) { Front(c, false); Front(c, false) },
	"R2": func(c *Cube, _ bool) { Right(c, false); Right(c, false) },
	"L2": func(c *Cube, _ bool) { Left(c, false); Left(c, false) },
	"B2": func(c *Cube, _ bool) { Back(c, false); Back(c, false) },
	"U2": func(c *Cube, _ bool) { Up(c, false); Up(c, false) },
	"D2": func(c *Cube, _ bool) { Down(c, false); Down(c, false) },
}

var NotationMatrix1 map[string]func(*Cube, bool) = map[string]func(*Cube, bool){
	"F":  Front,
	"B":  Back,
	"R":  Right,
	"L":  Left,
	"F'": RFront,
	"B'": RBack,
	"R'": RRight,
	"L'": RLeft,
	"F2": func(c *Cube, _ bool) { Front(c, false); Front(c, false) },
	"R2": func(c *Cube, _ bool) { Right(c, false); Right(c, false) },
	"L2": func(c *Cube, _ bool) { Left(c, false); Left(c, false) },
	"B2": func(c *Cube, _ bool) { Back(c, false); Back(c, false) },
	"U2": func(c *Cube, _ bool) { Up(c, false); Up(c, false) },
	"D2": func(c *Cube, _ bool) { Down(c, false); Down(c, false) },
}

var NotationMatrix2 map[string]func(*Cube, bool) = map[string]func(*Cube, bool){
	// "F": Front,
	// "B": Back,
	// "F'": RFront,
	// "B'": RBack,
	"R":  Right,
	"L":  Left,
	"R'": RRight,
	"L'": RLeft,
	"F2": func(c *Cube, _ bool) { Front(c, false); Front(c, false) },
	"R2": func(c *Cube, _ bool) { Right(c, false); Right(c, false) },
	"L2": func(c *Cube, _ bool) { Left(c, false); Left(c, false) },
	"B2": func(c *Cube, _ bool) { Back(c, false); Back(c, false) },
	"U2": func(c *Cube, _ bool) { Up(c, false); Up(c, false) },
	"D2": func(c *Cube, _ bool) { Down(c, false); Down(c, false) },
}

var NotationMatrix3 map[string]func(*Cube, bool) = map[string]func(*Cube, bool){
	"F2": func(c *Cube, _ bool) { Front(c, false); Front(c, false) },
	"R2": func(c *Cube, _ bool) { Right(c, false); Right(c, false) },
	"L2": func(c *Cube, _ bool) { Left(c, false); Left(c, false) },
	"B2": func(c *Cube, _ bool) { Back(c, false); Back(c, false) },
	"U2": func(c *Cube, _ bool) { Up(c, false); Up(c, false) },
	"D2": func(c *Cube, _ bool) { Down(c, false); Down(c, false) },
}

func Create() *Cube {
	return &Cube{
		Yellow: SolvedYellow,
		Orange: SolvedOrange,
		Green:  SolvedGreen,
		White:  SolvedWhite,
		Blue:   SolvedBlue,
		Red:    SolvedRed,
	}
}

func Copy(c *Cube) *Cube {
	return &Cube{
		Yellow: c.Yellow,
		Orange: c.Orange,
		Green:  c.Green,
		White:  c.White,
		Blue:   c.Blue,
		Red:    c.Red,
	}
}

func GetRandomMixSequence(n int) (sequence string) {
	for i := 0; i < n; i += 1 {
		sequence += PossibleMoves[rand.Intn(len(PossibleMoves))]

		if i+1 != n {
			sequence += " "
		}
	}

	return
}

func ApplyMoves(cube *Cube, sequence []string, callback func(cube *Cube)) *Cube {
	for _, move := range sequence {
		if function, ok := NotationMatrix[strings.ToUpper(move)]; ok {
			function(cube, false)

			if callback != nil {
				callback(cube)
			}
		} else {
			log.Fatalf("Unsupported move: %s", move)
		}
	}

	return cube
}

func IsSolved(cube *Cube) bool {
	if cube.Yellow == SolvedYellow && cube.Orange == SolvedOrange &&
		cube.Green == SolvedGreen && cube.White == SolvedWhite &&
		cube.Blue == SolvedBlue && cube.Red == SolvedRed {
		return true
	}

	return false
}

func IsGoodEdges(c *Cube) bool {
	var bad uint64 = Orange | Red

	if ((c.Green&GET_1)>>8)&bad > 0 ||
		((c.Green&GET_3)>>24)&bad > 0 ||
		((c.Green&GET_5)>>40)&bad > 0 ||
		((c.Green&GET_7)>>56)&bad > 0 {
		return false
	}

	if ((c.Blue&GET_1)>>8)&bad > 0 ||
		((c.Blue&GET_3)>>24)&bad > 0 ||
		((c.Blue&GET_5)>>40)&bad > 0 ||
		((c.Blue&GET_7)>>56)&bad > 0 {
		return false
	}

	if ((c.Yellow&GET_3)>>24)&bad > 0 ||
		((c.Yellow&GET_7)>>56)&bad > 0 {
		return false
	}

	if ((c.White&GET_3)>>24)&bad > 0 ||
		((c.White&GET_7)>>56)&bad > 0 {
		return false
	}

	return true
}

func IsGoodCorners(c *Cube) bool {
	var good uint64 = Orange | Red
	var edges uint64 = White | Yellow | Blue | Green

	if ((c.Yellow&GET_1)>>8)&edges == 0 ||
		((c.Yellow&GET_5)>>40)&edges == 0 {
		return false
	}

	if ((c.Blue&GET_1)>>8)&edges == 0 ||
		((c.Blue&GET_5)>>40)&edges == 0 {
		return false
	}

	if ((c.Green&GET_1)>>8)&edges == 0 ||
		((c.Green&GET_5)>>40)&edges == 0 {
		return false
	}

	if ((c.White&GET_1)>>8)&edges == 0 ||
		((c.White&GET_5)>>40)&edges == 0 {
		return false
	}

	if (c.Red&GET_0)&good == 0 ||
		((c.Red&GET_2)>>16)&good == 0 ||
		((c.Red&GET_4)>>32)&good == 0 ||
		((c.Red&GET_6)>>48)&good == 0 {
		return false
	}

	if (c.Orange&GET_0)&good == 0 ||
		((c.Orange&GET_2)>>16)&good == 0 ||
		((c.Orange&GET_4)>>32)&good == 0 ||
		((c.Orange&GET_6)>>48)&good == 0 {
		return false
	}

	return true
}

func IsGoodSides(c *Cube) bool {
	if (c.Yellow&GET_0)&(Yellow|White) == 0 ||
		((c.Yellow&GET_1)>>8)&(Yellow|White) == 0 ||
		((c.Yellow&GET_2)>>16)&(Yellow|White) == 0 ||
		((c.Yellow&GET_3)>>24)&(Yellow|White) == 0 ||
		((c.Yellow&GET_4)>>32)&(Yellow|White) == 0 ||
		((c.Yellow&GET_5)>>40)&(Yellow|White) == 0 ||
		((c.Yellow&GET_6)>>48)&(Yellow|White) == 0 ||
		((c.Yellow&GET_7)>>56)&(Yellow|White) == 0 {
		return false
	}

	if (c.White&GET_0)&(Yellow|White) == 0 ||
		((c.White&GET_1)>>8)&(Yellow|White) == 0 ||
		((c.White&GET_2)>>16)&(Yellow|White) == 0 ||
		((c.White&GET_3)>>24)&(Yellow|White) == 0 ||
		((c.White&GET_4)>>32)&(Yellow|White) == 0 ||
		((c.White&GET_5)>>40)&(Yellow|White) == 0 ||
		((c.White&GET_6)>>48)&(Yellow|White) == 0 ||
		((c.White&GET_7)>>56)&(Yellow|White) == 0 {
		return false
	}

	if (c.Green&GET_0)&(Green|Blue) == 0 ||
		((c.Green&GET_1)>>8)&(Green|Blue) == 0 ||
		((c.Green&GET_2)>>16)&(Green|Blue) == 0 ||
		((c.Green&GET_3)>>24)&(Green|Blue) == 0 ||
		((c.Green&GET_4)>>32)&(Green|Blue) == 0 ||
		((c.Green&GET_5)>>40)&(Green|Blue) == 0 ||
		((c.Green&GET_6)>>48)&(Green|Blue) == 0 ||
		((c.Green&GET_7)>>56)&(Green|Blue) == 0 {
		return false
	}

	if (c.Blue&GET_0)&(Green|Blue) == 0 ||
		((c.Blue&GET_1)>>8)&(Green|Blue) == 0 ||
		((c.Blue&GET_2)>>16)&(Green|Blue) == 0 ||
		((c.Blue&GET_3)>>24)&(Green|Blue) == 0 ||
		((c.Blue&GET_4)>>32)&(Green|Blue) == 0 ||
		((c.Blue&GET_5)>>40)&(Green|Blue) == 0 ||
		((c.Blue&GET_6)>>48)&(Green|Blue) == 0 ||
		((c.Blue&GET_7)>>56)&(Green|Blue) == 0 {
		return false
	}

	if (c.Orange&GET_0)&(Orange|Red) == 0 ||
		((c.Orange&GET_1)>>8)&(Orange|Red) == 0 ||
		((c.Orange&GET_2)>>16)&(Orange|Red) == 0 ||
		((c.Orange&GET_3)>>24)&(Orange|Red) == 0 ||
		((c.Orange&GET_4)>>32)&(Orange|Red) == 0 ||
		((c.Orange&GET_5)>>40)&(Orange|Red) == 0 ||
		((c.Orange&GET_6)>>48)&(Orange|Red) == 0 ||
		((c.Orange&GET_7)>>56)&(Orange|Red) == 0 {
		return false
	}

	if (c.Red&GET_0)&(Orange|Red) == 0 ||
		((c.Red&GET_1)>>8)&(Orange|Red) == 0 ||
		((c.Red&GET_2)>>16)&(Orange|Red) == 0 ||
		((c.Red&GET_3)>>24)&(Orange|Red) == 0 ||
		((c.Red&GET_4)>>32)&(Orange|Red) == 0 ||
		((c.Red&GET_5)>>40)&(Orange|Red) == 0 ||
		((c.Red&GET_6)>>48)&(Orange|Red) == 0 ||
		((c.Red&GET_7)>>56)&(Orange|Red) == 0 {
		return false
	}

	return true
}
