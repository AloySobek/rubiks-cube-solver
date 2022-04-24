package cube

import "fmt"

func Print(cube *Cube) {
	fmt.Println("Blue side:")
	printSide(cube.Blue, "B")
	fmt.Printf("\n\n")

	fmt.Println("Red side:")
	printSide(cube.Red, "R")
	fmt.Printf("\n\n")

	fmt.Println("Orange side:")
	printSide(cube.Orange, "O")
	fmt.Printf("\n\n")

	fmt.Println("Green side:")
	printSide(cube.Green, "G")
	fmt.Printf("\n\n")

	fmt.Println("Yellow side:")
	printSide(cube.Yellow, "Y")
	fmt.Printf("\n\n")

	fmt.Println("White side:")
	printSide(cube.White, "W")
	fmt.Printf("\n\n")
}

func printSide(side uint64, name string) {
	for i := 0; i < 8; i += 1 {
		switch (side & (uint64(0xFF) << (8 * i))) >> (8 * i) {
		case Blue:
			fmt.Printf("B ")
		case Red:
			fmt.Printf("R ")
		case Orange:
			fmt.Printf("O ")
		case Green:
			fmt.Printf("G ")
		case Yellow:
			fmt.Printf("Y ")
		case White:
			fmt.Printf("W ")
		}

		if i == 3 {
			fmt.Printf("%s ", name)
		}
	}
}
