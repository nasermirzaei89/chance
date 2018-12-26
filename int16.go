package chance

import "math"

// Int16Option is a type
type Int16Option func(*Int16Options)

// Int16Options is int16 options
type Int16Options struct {
	min int16
	max int16
}

func (ch *chance) Int16(options ...Int16Option) int16 {
	ops := Int16Options{min: -math.MaxInt16, max: math.MaxInt16}

	for i := range options {
		options[i](&ops)
	}

	ch.r.Seed(ch.seed)
	// TODO: handle error on bad options
	return int16(ch.r.Int31n(int32(ops.max) - int32(ops.min) + int32(ops.min)))
}

// Int16 returns a random int16
func Int16(options ...Int16Option) int16 {
	return defaultChance.Int16(options...)
}

// SetInt16Min sets min of random int16
func SetInt16Min(min int16) Int16Option {
	return func(iOpts *Int16Options) {
		iOpts.min = min
	}
}

// SetInt16Max sets max of random int16
func SetInt16Max(max int16) Int16Option {
	return func(iOpts *Int16Options) {
		iOpts.max = max
	}
}
