/*
A function can take zero or more arguments.
Notice that the type comes after the variable name.

Each block should only declare main as a function once.
Files within the same folder are considered within teh block.
*/

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
