package main

import (
	"fmt"
	"reflect"
)

type Farenheit float64

type Celsius float64

// type aliases
type DefaultUnit = Farenheit

func main() {
	var f Farenheit
	var c Celsius
	var d DefaultUnit = 21
	c = 32
	// explicit type conversion
	f = Farenheit(c)
	fmt.Println(f)
	fmt.Println(d)
	// testing type aliases
	c = 22
	f = 234
	// Farenheit is assigned to its type alias DefaultUnit
	d = f
	fmt.Println(reflect.TypeOf(d))
}
