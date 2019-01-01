package chance_test

import (
	"fmt"

	"github.com/nasermirzaei89/chance"
)

func ExampleByte() {
	fmt.Print(chance.Byte())
	// Output: 86
}

// Returns a random byte with max of 10
func ExampleByte_setMax() {
	fmt.Print(chance.Byte(chance.SetByteMax(10)))
	// Output: 1
}

// Returns a random byte with min of 10
func ExampleByte_setMin() {
	fmt.Print(chance.Byte(chance.SetByteMin(10)))
	// Output: 51
}

// Returns a random byte in range 0 to 10
func ExampleByte_setRange() {
	fmt.Print(chance.Byte(chance.SetByteMin(0), chance.SetByteMax(10)))
	// Output: 1
}
