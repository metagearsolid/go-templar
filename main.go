package main

import "fmt"

func main() {
	fmt.Printf("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちは世界\n")
	// function Printf called via syntax:  <pkgName><funcName>
}

// Conventions:
// Every executable program must have:
// Only one main package with an entry function main
// without any arguments or return values in the

/*
Package conventions:
    All Go files start with package <something>.
    In Go, package is always first, then import, then everything else.
    package main is required for a standalone executable..
    Every package should have a package comment, a block comment preceding the package clause.
    The package comment should introduce the package and provide information relevant to the package as a whole.
    Package comments should begin with the name of the thing being described and end in a period.
    By convention, packages are given lower case, single-word names; there should be no need for underscores or mixedCaps.
    Go’s convention is that the package name is the last element of the import path: the package imported as “crypto/rot13” should be named rot13.
*/
/*
25 keywords:
break 	default 	func 	interface 	select
case 	defer 	go 	map 	struct
chan 	else 	goto 	package 	switch
const 	fallthrough 	if 	range 	type
continue 	for 	import 	return 	var
*/
