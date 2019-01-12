package chance_test

import (
	"fmt"

	"github.com/nasermirzaei89/chance"
)

func ExampleUInt() {
	fmt.Print(chance.UInt())
	// Output: 0
}

// Returns a random uint with max of 10
func ExampleUInt_setMax() {
	fmt.Print(chance.UInt(chance.SetUIntMax(10)))
	// Output: 1
}

// Returns a random uint with min of 10
func ExampleUInt_setMin() {
	fmt.Print(chance.UInt(chance.SetUIntMin(10)))
	// Output: 9828766684487745576
}

// Returns a random uint in range 1 to 10
func ExampleUInt_setRange() {
	fmt.Print(chance.UInt(chance.SetUIntMin(1), chance.SetUIntMax(10)))
	// Output: 9
}
