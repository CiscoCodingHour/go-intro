package main

import "fmt"

func main() {

	fmt.Println("Counting...")

	for num := 1; num < 10; num++ {

		defer fmt.Println(num)
	}

	fmt.Println("Done")

}
