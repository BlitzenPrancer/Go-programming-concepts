// enumerations can be defined by types and constants in go
package main

import "fmt"

// DayOfWeek represents a day
type DayOfWeek int

// constants of type DayOfWeek
const (
	// sunday DayOfWeek = 1 << iota - this assigns constants exponential value
	sunday DayOfWeek = 1 + iota // starts from 1
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

func main() {
	fmt.Println(sunday)
	fmt.Println(tuesday)
}
