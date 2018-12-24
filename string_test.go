package chance_test

import (
	"fmt"

	"github.com/nasermirzaei89/chance"
)

func ExampleString() {
	fmt.Print(chance.String())
	// Output: BhLvh
}

// Returns a random string with length of 12
func ExampleString_setLength() {
	fmt.Print(chance.String(chance.SetStringLength(12)))
	// Output: BhLvhGNKqc#5
}

// Returns a random string from custom pool
func ExampleString_setPool() {
	fmt.Print(chance.String(chance.SetStringPool("aB1")))
	// Output: 1a11B
}
