package chance

import "math"

// IntOption is a type
type IntOption func(*IntOptions)

// IntOptions is int options
type IntOptions struct {
	intMin int
	intMax int
}

func (ch *chance) Int(options ...IntOption) int {
	ops := IntOptions{-math.MaxInt32, math.MaxInt32}

	for i := range options {
		options[i](&ops)
	}

	ch.r.Seed(ch.seed)
	// TODO: handle error on bad options
	return ch.r.Intn(ops.intMax-ops.intMin) + ops.intMin
}

// Int returns a random int
func Int(options ...IntOption) int {
	return defaultChance.Int(options...)
}

// SetIntMin sets min of random int
func SetIntMin(min int) IntOption {
	return func(iOpts *IntOptions) {
		iOpts.intMin = min
	}
}

// SetIntMax sets max of random int
func SetIntMax(max int) IntOption {
	return func(iOpts *IntOptions) {
		iOpts.intMax = max
	}
}
