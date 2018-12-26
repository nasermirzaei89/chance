package chance_test

import (
	"fmt"

	"github.com/nasermirzaei89/chance"
)

func ExampleInt8() {
	fmt.Print(chance.Int8())
	// Output: -84
}

// Returns a random int8 with max of 10
func ExampleInt8_setMax() {
	fmt.Print(chance.Int8(chance.SetInt8Max(10)))
	// Output: -102
}

// Returns a random int8 with min of 10
func ExampleInt8_setMin() {
	fmt.Print(chance.Int8(chance.SetInt8Min(10)))
	// Output: 33
}

// Returns a random int8 in range -10 to 10
func ExampleInt8_setRange() {
	fmt.Print(chance.Int8(chance.SetInt8Min(-10), chance.SetInt8Max(10)))
	// Output: -9
}
