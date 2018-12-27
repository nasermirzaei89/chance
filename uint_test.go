package chance_test

import (
	"fmt"

	"github.com/nasermirzaei89/chance"
)

func ExampleUInt() {
	fmt.Print(chance.UInt())
	// Output: 0
}

// Returns a random int64 with max of 10
func ExampleUInt_setMax() {
	fmt.Print(chance.UInt(chance.SetUIntMax(10)))
	// Output: 3
}

// Returns a random int64 with min of 10
func ExampleUInt_setMin() {
	fmt.Print(chance.UInt(chance.SetUIntMin(10)))
	// Output: 15352856648520921639
}

// Returns a random int64 in range -10 to 10
func ExampleUInt_setRange() {
	fmt.Print(chance.UInt(chance.SetUIntMin(1), chance.SetUIntMax(10)))
	// Output: 2
}
