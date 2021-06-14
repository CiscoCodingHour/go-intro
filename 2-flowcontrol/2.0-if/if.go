package main

import (
	"fmt"
)

func main() {

	num := 0

	if num < 0 {

		fmt.Println("It is a negative number")
	} else if num > 0 {
		fmt.Println("it is a positive number")
	} else {
		fmt.Println("The number is zero !")
	}

}

// Similar than "for" keyword
// After "if" keyword, there is no need of parenthesis for the expression.
// Braces are required to define scope.
