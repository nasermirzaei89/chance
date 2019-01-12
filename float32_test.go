package chance_test

import (
	"fmt"

	"github.com/nasermirzaei89/chance"
)

func ExampleFloat32() {
	fmt.Print(chance.Float32())
	// Output: +Inf
}

// Returns a random float32 with max of 10
func ExampleFloat32_setMax() {
	fmt.Print(chance.Float32(chance.SetFloat32Max(10)))
	// Output: -2.0243711e+37
}

// Returns a random float32 with min of 10
func ExampleFloat32_setMin() {
	fmt.Print(chance.Float32(chance.SetFloat32Min(10)))
	// Output: 2.2613806e+38
}

// Returns a random float32 in range -10 to 10
func ExampleFloat32_setRange() {
	fmt.Print(chance.Float32(chance.SetFloat32Min(-10), chance.SetFloat32Max(10)))
	// Output: -1.2457161
}
