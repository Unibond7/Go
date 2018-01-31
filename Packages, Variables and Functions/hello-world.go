/* Start of my Go Programming experience
Simple Hello World example which has been modified to show use of own library*/

package main

import (
	"fmt"

	"github.com/Unibond7/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
}
