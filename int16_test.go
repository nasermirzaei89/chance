package chance_test

import (
	"fmt"

	"github.com/nasermirzaei89/chance"
)

func ExampleInt16() {
	fmt.Print(chance.Int16())
	// Output: 7405
}

// Returns a random int16 with max of 10
func ExampleInt16_setMax() {
	fmt.Print(chance.Int16(chance.SetInt16Max(10)))
	// Output: 1
}

// Returns a random int16 with min of 10
func ExampleInt16_setMin() {
	fmt.Print(chance.Int16(chance.SetInt16Min(10)))
	// Output: 7405
}

// Returns a random int16 in range -10 to 10
func ExampleInt16_setRange() {
	fmt.Print(chance.Int16(chance.SetInt16Min(-10), chance.SetInt16Max(10)))
	// Output: 1
}
