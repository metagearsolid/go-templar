package main

import (
	"errors"
	"fmt"
)

func main() {
	// Change character of string
	s := "hello"   //declare string "hello"
	c := []byte(s) //convert string to []byte type
	c[0] = 'c'
	s2 := string(c) //convert back to string type
	fmt.Printf("%s\n", s2)

	// Combine Strings
	m := "hello,"
	n := " world"
	o := m + n
	fmt.Printf("%s\n", o)

	// Modify string
	p := "hello"
	p = "c" + s[1:] // cannot change string values by index, but can get values instead
	fmt.Printf("%s\n", p)

	// Error types
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}

}
