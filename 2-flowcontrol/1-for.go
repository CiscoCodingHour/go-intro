package main

import "fmt"

func main() {

	fmt.Println("=============== COUNT ===========================")

	sum := 0

	for number := 0; number < 10; number++ {

		sum += number // sum = sum + i

		fmt.Println(number)
	}

	fmt.Println("=============== SUM ===========================")
	fmt.Println(sum)

}

/* for INIT ; CONDITION; POST

INIT: Executed before the first iteration
CONDITION: Evaluated before every iteration
POST: Executed at the end of every iteration

