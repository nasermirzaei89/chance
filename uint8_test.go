package chance_test

import (
	"fmt"

	"github.com/nasermirzaei89/chance"
)

func ExampleUInt8() {
	fmt.Print(chance.UInt8())
	// Output: 86
}

// Returns a random int8 with max of 10
func ExampleUInt8_setMax() {
	fmt.Print(chance.UInt8(chance.SetUInt8Max(10)))
	// Output: 1
}

// Returns a random int8 with min of 10
func ExampleUInt8_setMin() {
	fmt.Print(chance.UInt8(chance.SetUInt8Min(10)))
	// Output: 51
}

// Returns a random int8 in range -10 to 10
func ExampleUInt8_setRange() {
	fmt.Print(chance.UInt8(chance.SetUInt8Min(0), chance.SetUInt8Max(10)))
	// Output: 1
}
