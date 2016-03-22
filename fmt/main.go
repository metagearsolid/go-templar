/* value formatting for package fmt
 */

package main

import "fmt"

func main() {
	a := 12
	b := false
	c := 9000
	d := 3.141
	e := "JOJO'S BIZARRE ADVENTURES"
	f := 0.0001

	// GENERAL
	fmt.Printf("Variable a has value of %v\n", a)  // %v\n prints value: default format
	fmt.Printf("Variable a has value of %#v\n", a) // go-syntax representation of value
	fmt.Printf("Variable a has value of %T\n", a)  // go-syntax representation of value type
	fmt.Printf("A percentage sign %%\n")           // %t\n prints percentage sign

	/*
	  Default format for %v:
	  bool:                    %t
	  int, int8 etc.:          %d
	  uint, uint8 etc.:        %d, %x if printed with %#v
	  float32, complex64, etc: %g
	  string:                  %s
	  chan:                    %p
	  pointer:                 %p
	*/

	// BOOLEAN
	fmt.Printf("Variable b has value of %t\n", b) // %t\n prints value: true or false

	// INTEGER
	fmt.Printf("Variable c has value of %b\n", c) //base 2
	fmt.Printf("Variable c has value of %c\n", c) //unicode character
	fmt.Printf("Variable c has value of %d\n", c) //base 10
	fmt.Printf("Variable c has value of %o\n", c) //base 8
	fmt.Printf("Variable c has value of %q\n", c) //single-quoted character literal safely escaped with Go syntax
	fmt.Printf("Variable c has value of %x\n", c) //base 16, lower-case
	fmt.Printf("Variable c has value of %X\n", c) //base 16, upper-case
	fmt.Printf("Variable c has value of %U\n", c) //Unicode format

	// FLOAT + COMPLEX
	fmt.Printf("Variable d has value of %b\n", d) //decimalless scientific notation with exponent a power of two in the manner of strconv
	fmt.Printf("Variable d has value of %e\n", d) //scientific notation "e"
	fmt.Printf("Variable d has value of %E\n", d) //scientific notation "E"
	fmt.Printf("Variable d has value of %f\n", d) // decimal point, no exponent
	fmt.Printf("Variable d has value of %F\n", d) // synonym for %f
	fmt.Printf("Variable d has value of %g\n", d) // %e for large exponents, %f otherwise
	fmt.Printf("Variable d has value of %e\n", d) // %E for large exponents, %F otherwise

	// string
	fmt.Printf("Variable e has value of %s\n", e) // uninterpreted bytes of the string or slice
	fmt.Printf("Variable e has value of %q\n", e) // a double-quoted string safely escaped with Go syntax
	fmt.Printf("Variable e has value of %x\n", e) // base 16, lower-case, 2 characters per byte
	fmt.Printf("Variable e has value of %X\n", e) // base 16, upper-case, 2 characters per byte

	// Width: specified by an optional decimal number immediately preceding the verb
	// Precision:  specified after the optional width
	fmt.Printf("Variable d has value of %2f\n", f)   // decimal point, no exponent, width 2
	fmt.Printf("Variable d has value of %2.2f\n", f) // decimal point, no exponent, width 2, precision 2

	//pointer
	//fmt.Printf("Variable e has value of %p\n", e) // base 16, upper-case, 2 characters per byte
}
