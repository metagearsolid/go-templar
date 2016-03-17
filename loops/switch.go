package loops

import "fmt"

func friends() {
	switch "Mario" {
	case "Sonic":
		fmt.Println("Wassup Toad")
	case "Mario":
		fmt.Println("Wassup Mario")
	case "Peach":
		fmt.Println("Wassup Koopa")
	default:
		fmt.Println("Have you no friends?")
	}
}
