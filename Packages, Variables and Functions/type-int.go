/*
When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax),
the variable's type is inferred from the value on the right hand side.
*/

package main

import "fmt"

func main() {
	v := 63.2 //Change to see what type I am!!
	fmt.Printf("v is of type %T\n", v)
}
