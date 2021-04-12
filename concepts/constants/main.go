package main

import "fmt"

// the value of this constant is not going to change
// this constant can be used throughout the application
const Pi float32 = 3.1415926

// untyped constant
const z = 321

// defining multiple constants
const (
	a = 1
	b = 2
)

func main() {
	var x float64
	// explicit type conversion
	x = float64(Pi)
	fmt.Println(x)
	fmt.Println(z)
}
