package main

import "fmt"

func main() {

	var a [10]int
	a[0] = 42
	a[1] = 13
	fmt.Printf("The first element in %d\n", a[0])
	fmt.Printf("The last element in %d\n", a[9])

	// declare array with type int length 3, initialized with 1, 2, 3 for [0],[1],[2] respectively
	array := [3]int{1, 2, 3}
	fmt.Printf("The first element in %d\n", array[0]) //returns 1
	fmt.Printf("The last element in %d\n", array[2])  //returns 3
	//length of array cannot change after declaration

	// Define 2-D array with 2 elements, and each element has 4 elements
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	fmt.Printf("The %d\n", doubleArray[1])

}
