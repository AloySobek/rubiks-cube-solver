package cube

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

const (
	Yellow = 1
	Orange = 2
	Green  = 4
	White  = 8
	Blue   = 16
	Red    = 32
)

type Cube struct {
	Yellow uint64 // Up side
	Orange uint64 // Left side
	Green  uint64 // Back side
	White  uint64 // Down side
	Blue   uint64 // Front side
	Red    uint64 // Right side
}

var PossibleMoves []string = []string{
	"F", "R", "L", "B", "U", "D",
	"F'", "R'", "L'", "B'", "U'", "D'",
	"F2", "R2", "L2", "B2", "U2", "D2",
}
var NotationMatrix map[string]func(*Cube, bool, bool) = map[string]func(*Cube, bool, bool){
	"F": Front, "R": Right, "L": Left, "B": Back, "U": Up, "D": Down,
	"F'": RFront, "R'": RRight, "L'": RLeft, "B'": RBack, "U'": RUp, "D'": RDown,
	"F2": DFront, "R2": DRight, "L2": DLeft, "B2": DBack, "U2": DUp, "D2": DDown,
}

func Create() *Cube {
	return &Cube{
		Yellow: uint64(0) | Yellow | (Yellow << 8) | (Yellow << 16) | (Yellow << 24) | (Yellow << 32) | (Yellow << 40) | (Yellow << 48) | (Yellow << 56),
		Orange: uint64(0) | Orange | (Orange << 8) | (Orange << 16) | (Orange << 24) | (Orange << 32) | (Orange << 40) | (Orange << 48) | (Orange << 56),
		Green:  uint64(0) | Green | (Green << 8) | (Green << 16) | (Green << 24) | (Green << 32) | (Green << 40) | (Green << 48) | (Green << 56),
		White:  uint64(0) | White | (White << 8) | (White << 16) | (White << 24) | (White << 32) | (White << 40) | (White << 48) | (White << 56),
		Blue:   uint64(0) | Blue | (Blue << 8) | (Blue << 16) | (Blue << 24) | (Blue << 32) | (Blue << 40) | (Blue << 48) | (Blue << 56),
		Red:    uint64(0) | Red | (Red << 8) | (Red << 16) | (Red << 24) | (Red << 32) | (Red << 40) | (Red << 48) | (Red << 56),
	}
}

func GetRandomMixSequence() (sequence string) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < 100; i += 1 {
		sequence += PossibleMoves[rand.Intn(len(PossibleMoves))]

		if i+1 != 100 {
			sequence += " "
		}
	}

	return
}

func Mix(cube *Cube, sequence string) *Cube {
	for _, move := range strings.Split(sequence, " ") {
		if function, ok := NotationMatrix[strings.ToUpper(move)]; ok {
			function(cube, false, false)
		} else {
			log.Fatalf("Unsupported move: %s", move)
		}
	}

	return cube
}
