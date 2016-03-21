/*Statements:
should be declared on seperate line for readability and go will auto insert ";"
*/
/*
Integer Types:
int8, int16, int32, int64, rune,
uint8, uint16, uint32, uint64, byte
Note that rune is alias of int32 and byte is alias of uint8

Float types:
float32, float64, complex128, complex64

Constants:
true, false, iota

Types: 	bool, byte, complex64, complex128, error, float32,
  	float64, int, int8, int16, int32, int64, rune, string,
  	uint, uint8, uint16, uint32, uint64, uintptr
Constants: 	true, false, iota
Null or no value: 	nil
Functions: 	append, cap, close, complex, copy, delete, imag, len,
  	make, new, panic, print, println, real, recover
*/
/* UTF-8
In the UTF-8 world characters are sometimes called runes. Mostly, when people talk about characters, they mean 8 bit characters. As UTF-8 characters may be up to 32 bits the word rune is used.  The Go language defines the word rune as an alias for the type int32.
*/

package main

import "fmt"

func main() {

	//define variables using shorthand :=, without defining type (letting go autodetect type)
	//can only be used inside functions
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	//define multiple variables
	v1, v2, v3 := 1, 2, 3
	v4 := v1 + v2 + v3

	// iota
	const (
		x = iota // x == 0
		y        // y == 1
		z        // z == 2
		w        // If there is no expression after the constants name, it uses the last expression,
		//so it's saying w = iota implicitly. Therefore w == 3, and y and z both can omit "= iota" as well.
	)

	const v = iota //once iota meets keyword "const", it resets to '0'

	//calling variable type
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)

	//calling variable value
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)

	//calling total from multiple variables
	fmt.Printf("%v \n", v4)

	//calling iota
	fmt.Printf("%v \n", x)
	fmt.Printf("%v \n", y)
	fmt.Printf("%v \n", z)
	fmt.Printf("%v \n", w)
	fmt.Printf("%v \n", v)
}
