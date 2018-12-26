package chance_test

import (
	"fmt"

	"github.com/nasermirzaei89/chance"
)

func ExampleInt64() {
	fmt.Print(chance.Int64())
	// Output: 965320773520391507
}

// Returns a random int64 with max of 10
func ExampleInt64_setMax() {
	fmt.Print(chance.Int64(chance.SetInt64Max(10)))
	// Output: -574485224192377852
}

// Returns a random int64 with min of 10
func ExampleInt64_setMin() {
	fmt.Print(chance.Int64(chance.SetInt64Min(10)))
	// Output: 965320773520391527
}

// Returns a random int64 in range -10 to 10
func ExampleInt64_setRange() {
	fmt.Print(chance.Int64(chance.SetInt64Min(-10), chance.SetInt64Max(10)))
	// Output: 0
}
