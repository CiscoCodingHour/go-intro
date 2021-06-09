package main

import (
	"fmt"
)

func main() {

	var value1 int = 150

	var value2 int = 30

	var value3 float32 = 30.5

	result1 := value1 + value2

	result2 := value1 + int(value3) // Remove the conversion for value3 and the program will not compile.

	fmt.Println(result1)
	fmt.Println(result2)

}
