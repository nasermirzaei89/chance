package chance_test

import (
	"fmt"

	"github.com/nasermirzaei89/chance"
)

func ExampleInt() {
	fmt.Print(chance.Int())
	// Output: 583532949
}

// Returns a random int with max of 10
func ExampleInt_setMax() {
	fmt.Print(chance.Int(chance.SetIntMax(10)))
	// Output: -1764108444
}

// Returns a random int with min of 10
func ExampleInt_setMin() {
	fmt.Print(chance.Int(chance.SetIntMin(10)))
	// Output: 1298498091
}
