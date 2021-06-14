package main

import "fmt"

// Variables are declared but not manually initialized (nothing was assigned)

var NumberOfApples int

var NumberOfOranges int

var tempKrakow float32

var tempLisbon float32

var carColor string

var wallColor string

var CiscoRocks bool

var KrakowRocks bool

func main() {

	var NumberOfPears int
	var tempBrussels float32
	var bloodColor string
	var LisbonRocks bool

	fmt.Println("============= Apples and Oranges =============================")
	fmt.Println(NumberOfApples)
	fmt.Println(NumberOfOranges)
	fmt.Println("============= Temp. Krakow and Lisbon =========================")
	fmt.Println(tempKrakow)
	fmt.Println(tempLisbon)
	fmt.Println("============= Car and Wall Color =============================")
	fmt.Println(carColor)
	fmt.Println(wallColor)
	fmt.Println("============= Do they rock ? ===================================")
	fmt.Println(CiscoRocks)
	fmt.Println(KrakowRocks)
	fmt.Println("============= Fruit ===========================================")
	fmt.Println(NumberOfPears)
	fmt.Println("============= Temp ============================================")
	fmt.Println(tempBrussels)
	fmt.Println("============= Color ===========================================")
	fmt.Println(bloodColor)
	fmt.Println("============= Do they rock ? ===================================")
	fmt.Print(LisbonRocks)

}
