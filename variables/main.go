/*
Define three variables without type "type" and without keyword "var", and initialize their values.
vname1 is v1，vname2 is v2，vname3 is v3
*/
// vname1, vname2, vname3 := v1, v2, v3 //can only be used inside functions

package main

import "fmt"

func main() {
	//var message string          //initialize message as string
	//var message = "Hello World" //declaring string is not necessary when providing string
	//var a, b, c int             //initialize a, b, c as ints (optional)

	//can only be done in function, type can be dropped when initializing only
	message := "Hello World!" //'assign Hello World to message'
	a, b, c := 1, false, 3
	const constantName = "value" //constants determined at complie cannot be changed during runtime
	//3 kinds of constants are:
	//boolean, string, constants
	var isActive bool //bool is default false

	/*
		Numerical Types:
		rune, int8, int16, int32, int64, byte, uint8, uint16, uint32, uint64. Note that rune is alias of int32 and byte is alias of uint8

		Float types have the float32 and float64 types and no type called float

		complex numbers: complex128, complex64

	*/
	var d int8
	var e complex64 = 5 + 5i

	/*
													Strings:

													Change character by index:
													var s string = "hello"
													s[0] = 'c'

													s := "hello"
													c := []byte(s)  // convert string to []byte type
													c[0] = 'c'
													s2 := string(c)  // convert back to string type
													fmt.Printf("%s\n", s2)

													Combine:
													s := "hello,"
												m := " world"
												a := s + m
												fmt.Printf("%s\n", a)

												s := "hello"
										s = "c" + s[1:] // you cannot change string values by index, but you can get values instead.
										fmt.Printf("%s\n", s)

										m := `hello
										world`

										Define by group:
										Basic form.

										import "fmt"
										import "os"

										const i = 100
										const pi = 3.1415
										const prefix = "Go_"

										var i int
										var pi float32
										var prefix string

										Group form.

										import(
										    "fmt"
										    "os"
										)

										const(
										    i = 100
										    pi = 3.1415
										    prefix = "Go_"
										)

										var(
										    i int
										    pi float32
										    prefix string
										)

								iota keyword enumerate
										const(
										    x = iota  // x == 0
										    y = iota  // y == 1
										    z = iota  // z == 2
										    w  // If there is no expression after the constants name, it uses the last expression,
										    //so it's saying w = iota implicitly. Therefore w == 3, and y and z both can omit "= iota" as well.
										)

										const v = iota // once iota meets keyword `const`, it resets to `0`, so v = 0.

										const (
										  e, f, g = iota, iota, iota // e=0,f=0,g=0 values of iota are same in one line.
										)

						Arrays:
						var arr [n]type

						var arr [10]int  // an array of type [10]int
						arr[0] = 42      // array is 0-based
						arr[1] = 13      // assign value to element
						fmt.Printf("The first element is %d\n", arr[0])
						// get element value, it returns 42
						fmt.Printf("The last element is %d\n", arr[9])
						//it returns default value of 10th element in this array, which is 0 in this case.

						a := [3]int{1, 2, 3} // define an int array with 3 elements

						b := [10]int{1, 2, 3}
						// define a int array with 10 elements, of which the first three are assigned.
						//The rest of them use the default value 0.

						c := [...]int{4, 5, 6} // use `…` to replace the length parameter and Go will calculate it for you.

						// define a two-dimensional array with 2 elements, and each element has 4 elements.
				doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

				// The declaration can be written more concisely as follows.
				easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

		Slice (dynamic array)

		// just like defining an array, but this time, we exclude the length.
		var fslice []int

		slice := []byte {'a', 'b', 'c', 'd'}

	*/

	fmt.Println(message, a, b, c)
	fmt.Println(constantName)
	fmt.Println(isActive)
	fmt.Println(d)
	fmt.Println(e)
}
