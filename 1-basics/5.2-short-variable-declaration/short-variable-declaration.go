package main

import "fmt"

// NumberOfCars := 2

func main() {

	// short-declaration notation ( := ) is used instead of "var" keyword.

	NumberOfApples := 5
	NumberOfOranges := 10
	tempKrakow := 32.5
	tempLisbon := 25.7
	carColor := "White"
	wallColor := "Green"
	CiscoRocks := true
	KrakowRocks := false
	NumberOfPears := 7
	tempBrussels := 21.3
	bloodColor := "Red"
	LisbonRocks := false

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
