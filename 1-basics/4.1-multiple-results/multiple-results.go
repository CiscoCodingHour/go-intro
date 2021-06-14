package main

import "fmt"

func calc(x int, y int) (a int, b string) {

	a = x + y

	if a > 0 {

		b = "The result is a positive number"
		return a, b
	} else if a < 0 {
		b = "The result is a negative number"
		return a, b
	} else {

		b = "The result is zero"
		return a, b
	}
}

func main() {

	total, info := calc(-5, 5)

	fmt.Println("========================================")
	fmt.Printf("The result was: %v \n", total)
	fmt.Println("========================================")
	fmt.Println(info)
}
