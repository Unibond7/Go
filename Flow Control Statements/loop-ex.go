/*
Exercise to create a function that implements the square root
*/

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, float64) {
	z := float64(x)
	diff, y, i := 0.0, 0.0, 0.0
	for diff > 1e-10 && diff > -1e-10 {
		y -= (z*z - x) / (2 * z)
		diff = (z - y)
		i++
		z = y
	}
	return z, i
}

func main() {
	z, i := Sqrt(2)
	fmt.Printf("Ans from own func = %g after %g iterations\n", z, i)
	fmt.Println("Ans from  Go func =", math.Sqrt(2))
}
