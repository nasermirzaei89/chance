package chance_test

import (
	"fmt"

	"github.com/nasermirzaei89/chance"
)

func ExampleInt32() {
	fmt.Print(chance.Int32())
	// Output: 583532949
}

// Returns a random int32 with max of 10
func ExampleInt32_setMax() {
	fmt.Print(chance.Int32(chance.SetInt32Max(10)))
	// Output: -1764108444
}

// Returns a random int32 with min of 10
func ExampleInt32_setMin() {
	fmt.Print(chance.Int32(chance.SetInt32Min(10)))
	// Output: 783690945
}

// Returns a random int32 in range -10 to 10
func ExampleInt32_setRange() {
	fmt.Print(chance.Int32(chance.SetInt32Min(-10), chance.SetInt32Max(10)))
	// Output: 0
}
