
![Intro to Go](https://github.com/rabocse/go-intro/blob/dev/pictures/IntroToGo.png)    

<p>&nbsp;</p>
<p>&nbsp;</p>

---
## Index 
---
<p>&nbsp;</p>

- Go (Golang)
- Basics:
	- hello.go
	- packages.go
	- imports.go
	- exported-name.go
	- functions.go
    	- multiple-results.go
	- variables.go
    	- variables-with-initializers.go
    	- short-variable-declaration.go
	- basic-types.go
	- zero-value.go
	- type-conversions.go
	- type-inference.go
	- constants.go
  
- Flowcontrol:
    - for.go
        - for-continued.go
        - while.go
        - forever.go
    - if.go
        - if-with-a-short-declartion.go
    - switch.go
    - functions.go
        - blank-identifier.go
        - error-handling.go // NOT DEFINED !!!!!!!!!!!!!
    - defer.go
    - stacking.defers.go
	

<p>&nbsp;</p>
<p>&nbsp;</p>

---
## __0-Go(Golang)__
---
<p>&nbsp;</p>

- "Golang" is a nickname. "Go" is the official name.
- __Statically typed__.
- __Compiled__
- Syntactically "similar" to C. 
- Memory safety, garbage collection.
- Concurrency.
- ["Go is the most _pythonic_ language"](https://www.youtube.com/watch?v=aYbNh3NS7jA&t=1471s) By Guido van Rossum (Inventor of Python)

<p>&nbsp;</p>

---
## __1-Basics__
---
<p>&nbsp;</p>

__hello.go__

```go

package main

import "fmt"

func main() {

	fmt.Println("Hello, Hola")
}
```
&nbsp;


 


__packages.go__ 
   

```go
package main 

import ( 
	"fmt"      
	"math/rand" 
)

func main() {

	fmt.Println("My favorite number is", rand.Intn(10)) 
}

```
- A __package__ is a collection of code that all does similar things like formatting strings or drawing images.
- The __package__ clause gives the name of the package that this code will become a part of.
- The package _main_ is required if this code is going to be run from terminal. 
   
<p>&nbsp;</p>


__imports.go__

```go

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("Now you have %g problems. \n ", math.Sqrt(16)) 
}
```
- Import name vs Import path.
- Example below shows the import path reveals the directory path where rand package is located.

![ImportPath](https://github.com/rabocse/go-intro/blob/dev/pictures/ImportPath.png) 

<p>&nbsp;</p>
   
__exported-name.go__

```go

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(math.Pi)
}


```
- Notice how all the first letters from the constants are capitalized in const.go.
- If changed to "pi" (instead of "Pi"), the constant would become _unexported_.

![ExportedPi](https://github.com/rabocse/go-intro/blob/dev/pictures/ExportedPi.png)

<p>&nbsp;</p>

__functions.go__ 


```go

package main

import (
	"fmt"
)

func add(x int, y int) int { 

	return x + y

}

func main() {

	fmt.Println(add(3, 15))
}


```
![ClassicFunction](https://github.com/rabocse/go-intro/blob/dev/pictures/ClassicFunction.png)

<p>&nbsp;</p>


__multiple-results.go__


```go 

package main

import "fmt"

func calc(x int, y int) (a int, b string) {

	a = x + y

	if a > 0 {

		b = "The result is a positive number"
		return a, b
	} else if a < 0 {
		b = "The result is a negative number"
		return a, b
	} else {

		b = "The result is zero"
		return a, b
	}
}

func main() {

	total, info := calc(-5, 5)

	fmt.Println("========================================")
	fmt.Printf("The result was: %v \n", total)
	fmt.Println("========================================")
	fmt.Println(info)
}

```

![MultipleOutputs](https://github.com/rabocse/go-intro/blob/dev/pictures/MultipleOutputs.png)

<p>&nbsp;</p>


__variables.go__

```go
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
```
<p>&nbsp;</p>


__variable-with-initializers.go__ 


```go
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
	fmt.Println("============= Lisbon rocks ? ==================================")
	fmt.Print(LisbonRocks)

}

```
<p>&nbsp;</p>

__zero-values.go__

```go

package main

import (
	"fmt"
)

func main() {

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

```
<p>&nbsp;</p>


__short-variable-declaration.go__


```go

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
```

<p>&nbsp;</p>

__type-inference.go__ 

```go

package main

import (
	"fmt"
)

func main() {

	value := 29 // change me !

	fmt.Printf("The variable \"value\" is of type: %T \n", value)
}


```

<p>&nbsp;</p>

__basic-types.go__

// Use a more adecuate Rune example instead of 'K'. Also explain why '' instead of "".

```go
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
```
<p>&nbsp;</p>

__type-conversions.go__

```go
package main

import (
	"fmt"
)

func main() {

	var value1 int = 150

	var value2 int = 30

	var value3 float32 = 30.5

	result1 := value1 + value2

	result2 := value1 + int(value3) // Remove the conversion for value3 and the program will not compile.

	fmt.Println(result1)
	fmt.Println(result2)

}
```
<p>&nbsp;</p>


__constants.go__    


```go
package main

import "fmt"

func main() {

	const myAge = 31
	const Pi = 3.14
	const ciscoRocks = true

	fmt.Printf("I am %v years old\n", myAge)
	fmt.Printf("This is the value of Pi: %v\n", Pi)
	fmt.Println("Cisco rocks ?", ciscoRocks)

}
```
<p>&nbsp;</p>
<p>&nbsp;</p>

---
## __2-Flowcontrol__
---
<p>&nbsp;</p>

__for.go__

```go
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
```
<p>&nbsp;</p>



__for-continued.go__

```go
package main

import "fmt"

func main() {

	fmt.Println("======================= LOOPING... ===========================")

	sum := 1

	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}

	fmt.Println("======================= TOTAL =================================")
	fmt.Println(sum)
}
```

<p>&nbsp;</p>

__while.go__

```go
package main

import "fmt"

func main() {

	sum := 1

	for sum < 10 { // "while" sum < 10, then ...

		sum += sum
	}

	fmt.Println(sum)
}
```
<p>&nbsp;</p>


__forever.go__

```go
package main

import "fmt"

func main() {

	for {

		fmt.Println("Looping forever ... !!! XD ")

	}
}
```
<p>&nbsp;</p>


__if.go__

```go
package main

import (
	"fmt"
)

func main() {

	num := 0

	if num < 0 {

		fmt.Println("It is a negative number")
	} else if num > 0 {
		fmt.Println("it is a positive number")
	} else {
		fmt.Println("The number is zero !")
	}

}
```
<p>&nbsp;</p>

__if-with-a-short-declaration.go__

```go
package main

import "fmt"

// printB4 returns a "delimiter string" ( =================================== )
func printB4() string {

	y := "=========================="

	return y
}

func main() {

	var num int = 50 // Change the value to play with the condition num > 100

	if delimiter := printB4(); num > 100 {

		fmt.Println(delimiter) // Only reason why we can print "delimiter" is because is executed before the condition  (x > 100). In other words, it does not depend on it.
		fmt.Println("Condition was \"TRUE\"")

	} else {
		fmt.Println(delimiter)
		fmt.Println("Condition was \"FALSE\"")
	}

	// fmt.Println(delimiter) // Variables declared by the statement are only in scope until the end of the if.

}
```
<p>&nbsp;</p>


__switch.go__

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(" ")
	fmt.Println("============================= 1- Switch ==================================")

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
	fmt.Println("================= 2.1- Switch with no condition =======================")

	// Switch with no condition = switch true

	number := 500

	switch false {

	case number < 100:
		fmt.Println("The number is less than 100")

	case number > 100:
		fmt.Println("The number is higher than 100")
	default:
		fmt.Println("The number is 100")
	}

	fmt.Println(" ")
	fmt.Println("=========== 2.2- Switch with no condition =============================")

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

```
<p>&nbsp;</p>

__functions.go__

```go
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
```
<p>&nbsp;</p>



__function-with_mult-return.go__

```go
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
```
<p>&nbsp;</p>


__Defer (pending)__

```go
package main

import "fmt"

func main() {

	defer fmt.Println("World")

	fmt.Println("Hello")
}
```
<p>&nbsp;</p>

__Stacking defers__

```go
package main

import "fmt"

func main() {

	fmt.Println("Counting...")

	for num := 1; num < 10; num++ {

		defer fmt.Println(num)
	}

	fmt.Println("Done")

}
```


