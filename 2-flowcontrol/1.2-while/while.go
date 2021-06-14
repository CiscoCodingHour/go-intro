package main

import "fmt"

func main() {

	sum := 1

	for sum < 10 { // "while" sum < 10, then ...

		sum += sum
	}

	fmt.Println(sum)
}
