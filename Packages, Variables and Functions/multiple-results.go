/*
A function can return any number of results.

The swap function returns two strings.
*/

package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("world", "hello")
	fmt.Println(a, b)
}
