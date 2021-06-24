package main

import "fmt"

// Variables are declared and initialized.

var NumberOfApples int = 5

var NumberOfOranges int = 10

var tempKrakow float32 = 32.5

var tempLisbon float32 = 25.7

var carColor string = "White"

var wallColor string = "Green"

var CiscoRocks bool = true

var KrakowRocks bool = false

func main() {

	var NumberOfPears int = 7
	var tempBrussels float32 = 21.3
	var bloodColor string = "Red"
	var LisbonRocks bool = false

	fmt.Println("============= Apples and Oranges =============================")
	fmt.Println(NumberOfApples)
	fmt.Println(NumberOfOranges)
	fmt.Println("============= Temp. Krakow and Lisbon ========================")
	fmt.Println(tempKrakow)
	fmt.Println(tempLisbon)
	fmt.Println("============= Car and Wall Color =============================")
	fmt.Println(carColor)
	fmt.Println(wallColor)
	fmt.Println("============= Cisco & Krakow. Do they rock ? =================")
	fmt.Println(CiscoRocks)
	fmt.Println(KrakowRocks)
	fmt.Println("============= Pears ==========================================")
	fmt.Println(NumberOfPears)
	fmt.Println("============= Temp. Brussels ==================================")
	fmt.Println(tempBrussels)
	fmt.Println("============= Blood Color =====================================")
	fmt.Println(bloodColor)
	fmt.Println("============= Lisbon rocks ? ===================================")
	fmt.Print(LisbonRocks)

}
