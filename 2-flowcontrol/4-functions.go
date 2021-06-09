package main

import (
	"fmt"
)

// 1- Function with return.
func double(x float32) float32 {

	return x * 2
}

// 2- Function without return.
func doublePrint(x float32) {

	fmt.Println(x * 2)
}

// 3- Func with two arguments and one return
func sum(a float32, b float32) float32 {

	return a + b
}

// 4- func with two arguments and no return
func sub(a float32, b float32) {

	fmt.Println(a - b)
}

func main() {

	fmt.Println("============================ 1- double ============================")

	y := double(3.5)
	fmt.Println(y)

	fmt.Println("============================ 2- doublePrint ======================")

	doublePrint(10)

	fmt.Println("============================ 3- sum ==============================")

	totalSum := sum(35, 67)
	fmt.Println(totalSum)

	fmt.Println("============================ 4- sub ==============================")

	sub(100, 10)

}
