package main

import "fmt"

// printB4 returns a "delimiter string" ( =================================== )
func printB4() string {

	y := "=========================="

	return y
}

func main() {

	var num int = 500 // Change the value to play with the condition num > 100

	if delimiter := printB4(); num > 100 {

		fmt.Println(delimiter) // Only reason why we can print "delimiter" is because is executed before the condition  (x > 100). In other words, it does not depend on it.
		fmt.Println("Condition was \"TRUE\"")

	} else {
		fmt.Println(delimiter)
		fmt.Println("Condition was \"FALSE\"")
	}

	// fmt.Println(delimiter) // Variables declared by the statement are only in scope until the end of the if.

}
