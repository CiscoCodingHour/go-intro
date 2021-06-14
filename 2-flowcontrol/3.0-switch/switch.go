package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(" ")
	fmt.Println("=============================== 1- Switch ===========================================")

	// Switch case

	num := 1

	fmt.Print("Write ", num, " as ")

	switch num {

	case 1:
		fmt.Println("\"one\"")
	case 2:
		fmt.Println("\"two\"")
	case 3:
		fmt.Println("three")
	}

	fmt.Println(" ")
	fmt.Println("================= 2.1- Switch with no condition ==========================================")

	// Switch with no condition = switch true

	number := 500

	switch true {

	case number < 100:
		fmt.Println("The number is less than 100")

	case number > 100:
		fmt.Println("The number is higher than 100")
	default:
		fmt.Println("The number is 100")
	}

	fmt.Println(" ")
	fmt.Println("================= 2.2- Switch with no condition ==========================================")

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}
