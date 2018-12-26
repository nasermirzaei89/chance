package chance

import "math"

// Int8Option is a type
type Int8Option func(*Int8Options)

// Int8Options is int8 options
type Int8Options struct {
	min int8
	max int8
}

func (ch *chance) Int8(options ...Int8Option) int8 {
	ops := Int8Options{min: -math.MaxInt8, max: math.MaxInt8}

	for i := range options {
		options[i](&ops)
	}

	ch.r.Seed(ch.seed)
	// TODO: handle error on bad options
	return int8(ch.r.Int31n(int32(ops.max)-int32(ops.min)) + int32(ops.min))
}

// Int8 returns a random int8
func Int8(options ...Int8Option) int8 {
	return defaultChance.Int8(options...)
}

// SetInt8Min sets min of random int8
func SetInt8Min(min int8) Int8Option {
	return func(iOpts *Int8Options) {
		iOpts.min = min
	}
}

// SetInt8Max sets max of random int8
func SetInt8Max(max int8) Int8Option {
	return func(iOpts *Int8Options) {
		iOpts.max = max
	}
}
