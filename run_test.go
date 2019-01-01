package chance_test

import (
	"fmt"

	"github.com/nasermirzaei89/chance"
)

func ExampleRune() {
	fmt.Print(chance.Rune())
	// Output: 583532949
}

// Returns a random rune with max of 10
func ExampleRune_setMax() {
	fmt.Print(chance.Rune(chance.SetRuneMax(10)))
	// Output: -1764108444
}

// Returns a random rune with min of 10
func ExampleRune_setMin() {
	fmt.Print(chance.Rune(chance.SetRuneMin(10)))
	// Output: 783690945
}

// Returns a random rune in range -10 to 10
func ExampleRune_setRange() {
	fmt.Print(chance.Rune(chance.SetRuneMin(-10), chance.SetRuneMax(10)))
	// Output: 0
}
