package chance_test

import (
	"fmt"

	"github.com/nasermirzaei89/chance"
)

func ExampleUInt16() {
	fmt.Print(chance.UInt16())
	// Output: 53126
}

// Returns a random uint16 with max of 10
func ExampleUInt16_setMax() {
	fmt.Print(chance.UInt16(chance.SetUInt16Max(10)))
	// Output: 1
}

// Returns a random uint16 with min of 10
func ExampleUInt16_setMin() {
	fmt.Print(chance.UInt16(chance.SetUInt16Min(10)))
	// Output: 53126
}

// Returns a random uint16 in range 0 to 10
func ExampleUInt16_setRange() {
	fmt.Print(chance.UInt16(chance.SetUInt16Min(0), chance.SetUInt16Max(10)))
	// Output: 1
}
