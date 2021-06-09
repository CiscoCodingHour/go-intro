package main

import "fmt"

func main() {

	fmt.Println("=============== COUNT ===========================")

	sum := 0

	for number := 0; number < 10; number++ {

		fmt.Println(number)

		sum = sum + number // sum += number

	}

	fmt.Println("=============== SUM ============================")
	fmt.Println(sum)

}
