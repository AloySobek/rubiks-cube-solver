package main

import (
	"math/rand"
	"strings"
)

func GetRandomMove(n int) (move string) {
	possibleMoves := strings.Split("F R L B U D F' R' L' B' U' D' F2 R2 L2 B2 U2 D2", " ")

	for i := 0; i < n; i += 1 {

		move += possibleMoves[rand.Intn(len(possibleMoves))]

		if i+1 != n {
			move += " "
		}
	}

	return
}
