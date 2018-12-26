package chance

import "math"

// Int32Option is a type
type Int32Option func(*Int32Options)

// Int32Options is int32 options
type Int32Options struct {
	min int32
	max int32
}

func (ch *chance) Int32(options ...Int32Option) int32 {
	ops := Int32Options{min: -math.MaxInt32, max: math.MaxInt32}

	for i := range options {
		options[i](&ops)
	}

	ch.r.Seed(ch.seed)
	// TODO: handle error on bad options
	return int32(ch.r.Int63n(int64(ops.max)-int64(ops.min)) + int64(ops.min))
}

// Int32 returns a random int32
func Int32(options ...Int32Option) int32 {
	return defaultChance.Int32(options...)
}

// SetInt32Min sets min of random int32
func SetInt32Min(min int32) Int32Option {
	return func(iOpts *Int32Options) {
		iOpts.min = min
	}
}

// SetInt32Max sets max of random int32
func SetInt32Max(max int32) Int32Option {
	return func(iOpts *Int32Options) {
		iOpts.max = max
	}
}
