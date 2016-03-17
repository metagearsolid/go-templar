package main

import (
	"fmt"

	"github.com/metagearsolid/go-templar/visibility/visable"
)

func main() {
	fmt.Println(visable.MyName)
	visable.PrintVar()
}
