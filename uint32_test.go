package chance_test

import (
	"fmt"

	"github.com/nasermirzaei89/chance"
)

func ExampleUInt32() {
	fmt.Print(chance.UInt32())
	// Output: 1432518515
}

// Returns a random uint32 with max of 10
func ExampleUInt32_setMax() {
	fmt.Print(chance.UInt32(chance.SetUInt32Max(10)))
	// Output: 0
}

// Returns a random uint32 with min of 10
func ExampleUInt32_setMin() {
	fmt.Print(chance.UInt32(chance.SetUInt32Min(10)))
	// Output: 1532597480
}

// Returns a random uint32 in range 6 to 10
func ExampleUInt32_setRange() {
	fmt.Print(chance.UInt32(chance.SetUInt32Min(6), chance.SetUInt32Max(10)))
	// Output: 8
}
