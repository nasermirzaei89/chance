package chance_test

import (
	"fmt"

	"github.com/nasermirzaei89/chance"
)

func ExampleInt() {
	fmt.Print(chance.Int())
	// Output: 965320773520391507
}

// Returns a random int with max of 10
func ExampleInt_setMax() {
	fmt.Print(chance.Int(chance.SetIntMax(10)))
	// Output: -574485224192377852
}

// Returns a random int with min of 10
func ExampleInt_setMin() {
	fmt.Print(chance.Int(chance.SetIntMin(10)))
	// Output: 965320773520391527
}

// Returns a random int in range -10 to 10
func ExampleInt_setRange() {
	fmt.Print(chance.Int(chance.SetIntMin(-10), chance.SetIntMax(10)))
	// Output: -9
}
