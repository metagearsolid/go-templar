/* zero values

zero values:
false : boolean
0 : integer
0.0 : float
"" : string
nil : pointer, functions, interfaces, slices, channels, and maps

*/

// b and c are not defined, therefore output is zero value "0"
package main

import "fmt"

func main() {
	//var message string          //initialize message as string
	//var message = "Hello World" //declaring string is not necessary when providing string
	//var a, b, c int             //initialize a, b, c as ints (optional)

	//can only be done in function, type can be dropped when initializing only
	message := "Hello World!" //'assign Hello World to message'
	a, b, c := 1, false, 3

	fmt.Println(message, a, b, c)
}
