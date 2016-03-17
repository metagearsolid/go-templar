package main

import (
	"fmt"
)

func main() {

	m := "Hello World"
	n := &m
	fmt.Println(m, n) //prints string "Hello World" and memory location of string
	fmt.Println("m - ", m)
	fmt.Println([]byte(m))
	fmt.Printf("%T\n", []byte(m)[0])
	fmt.Printf("%b\n", []byte(m)[0])
}
