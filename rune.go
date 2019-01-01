package chance

import "math"

// RuneOption is a type
type RuneOption func(*RuneOptions)

// RuneOptions is rune options
type RuneOptions struct {
	min rune
	max rune
}

func (ch *chance) Rune(options ...RuneOption) rune {
	ops := RuneOptions{min: -math.MaxInt32, max: math.MaxInt32}

	for i := range options {
		options[i](&ops)
	}

	ch.r.Seed(ch.seed)
	// TODO: handle error on bad options
	return rune(ch.r.Int63n(int64(ops.max)-int64(ops.min)) + int64(ops.min))
}

// Rune returns a random rune
func Rune(options ...RuneOption) rune {
	return defaultChance.Rune(options...)
}

// SetRuneMin sets min of random rune
func SetRuneMin(min rune) RuneOption {
	return func(iOpts *RuneOptions) {
		iOpts.min = min
	}
}

// SetRuneMax sets max of random rune
func SetRuneMax(max rune) RuneOption {
	return func(iOpts *RuneOptions) {
		iOpts.max = max
	}
}
