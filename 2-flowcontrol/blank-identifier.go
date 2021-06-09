package main

import "fmt"

func display(a float32, b float32) (c, d float32) {

	item1 := a
	item2 := b

	return item1, item2
}

func main() {

	fmt.Println("============ Getting both results ====================")
	fmt.Println(" ")
	result1, result2 := display(29, 31)

	fmt.Println(result1)
	fmt.Println(result2)

	fmt.Println("============ Ignoring one result ====================")
	fmt.Println(" ")

	result3, _ := display(37, 75)

	fmt.Println(result3)
	// fmt.Println(result4)

}
