//control structures:  sequence, conditional, iterative

// == 3 for loops: ==
// for init; condition; post {}
// for condition {}
// for {}

package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		// i := 0 is init | i <= 100 is condition | i++ post
		fmt.Println(i)
	}
}
