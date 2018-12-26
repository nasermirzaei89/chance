package chance

import "math"

// Int64Option is a type
type Int64Option func(*Int64Options)

// Int64Options is int64 options
type Int64Options struct {
	min int64
	max int64
}

func (ch *chance) Int64(options ...Int64Option) int64 {
	ops := Int64Options{min: -math.MaxInt64 / 2, max: math.MaxInt64 / 2} // FIXME: Use full range

	for i := range options {
		options[i](&ops)
	}

	ch.r.Seed(ch.seed)
	// TODO: handle error on bad options
	return ch.r.Int63n(ops.max-ops.min) + ops.min
}

// Int64 returns a random int64
func Int64(options ...Int64Option) int64 {
	return defaultChance.Int64(options...)
}

// SetInt64Min sets min of random int64
func SetInt64Min(min int64) Int64Option {
	return func(iOpts *Int64Options) {
		iOpts.min = min
	}
}

// SetInt64Max sets max of random int64
func SetInt64Max(max int64) Int64Option {
	return func(iOpts *Int64Options) {
		iOpts.max = max
	}
}
