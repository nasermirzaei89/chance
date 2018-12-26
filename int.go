package chance

import (
	"math"
	"strconv"
)

// IntOption is a type
type IntOption func(*IntOptions)

// IntOptions is int options
type IntOptions struct {
	min int
	max int
}

func (ch *chance) Int(options ...IntOption) int {
	ops := IntOptions{}
	// Check machine architecture
	if strconv.IntSize == 32 {
		ops.min = -math.MaxInt32 / 2 // FIXME: Use full range
		ops.max = math.MaxInt32 / 2  // FIXME: Use full range
	} else {
		ops.min = -math.MaxInt64 / 2 // FIXME: Use full range
		ops.max = math.MaxInt64 / 2  // FIXME: Use full range
	}

	for i := range options {
		options[i](&ops)
	}

	ch.r.Seed(ch.seed)
	// TODO: handle error on bad options

	return ch.r.Intn(ops.max-ops.min) + ops.min
}

// Int returns a random int
func Int(options ...IntOption) int {
	return defaultChance.Int(options...)
}

// SetIntMin sets min of random int
func SetIntMin(min int) IntOption {
	return func(iOpts *IntOptions) {
		iOpts.min = min
	}
}

// SetIntMax sets max of random int
func SetIntMax(max int) IntOption {
	return func(iOpts *IntOptions) {
		iOpts.max = max
	}
}
