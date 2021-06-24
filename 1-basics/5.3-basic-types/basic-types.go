package main

import (
	"fmt"
)

// Boolean

var summerRocks bool = true
var winterRocks bool = false

var (

	// Integers (Signed Integers)
	value1 int   = 1990
	value2 int8  = 127
	value3 int16 = 1990
	value4 int32 = 1990
	value5 int64 = 1990

	// Integers (Unsigned Integers)
	value6  uint    = 255
	value7  uint8   = 255
	value8  uint16  = 255
	value9  uint32  = 255
	value10 uint64  = 1<<64 - 1
	value11 uintptr = 255

	// Floats
	value12 float32 = 37.5
	value13 float64 = 37.5

	// Byte and Rune are aliases for uint8 and int32 respectively.
	value14 byte = 255
	value15 rune = 'k'

	// Complex numbers
	value16 complex64  = 2 + 0.5i
	value17 complex128 = 2 + 0.7i
)

func main() {

	fmt.Println(" ")
	fmt.Println("=============== Booleans =========================================")
	fmt.Printf("VALUE: %v				TYPE: %T \n", summerRocks, summerRocks)
	fmt.Printf("VALUE: %v				TYPE: %T \n", winterRocks, winterRocks)
	fmt.Println(" ")
	fmt.Println("=============== Integers (Signed Integers) =======================")
	fmt.Printf("VALUE: %v				TYPE: %T \n", value1, value1)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value2, value2)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value3, value3)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value4, value4)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value5, value5)
	fmt.Println(" ")
	fmt.Println("============== Integers (Unsigned Integers ) ======================")
	fmt.Printf("VALUE: %v				TYPE: %T \n", value6, value6)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value7, value7)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value8, value8)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value9, value9)
	fmt.Printf("VALUE: %v		TYPE: %T \n", value10, value10)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value11, value11)
	fmt.Println(" ")
	fmt.Println("============== Floats ==============================================")
	fmt.Printf("VALUE: %v				TYPE: %T \n", value12, value12)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value13, value13)
	fmt.Println(" ")
	fmt.Println("============== Byte and Rune =======================================")
	fmt.Printf("VALUE: %v				TYPE: %T \n", value14, value14)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value15, value15)
	fmt.Println(" ")
	fmt.Println("============== Complex Numbers =====================================")
	fmt.Printf("VALUE: %v				TYPE: %T \n", value16, value16)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value17, value17)

}
