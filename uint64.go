package chance

import (
	"math"
)

// UInt64Option is a type
type UInt64Option func(*UInt64Options)

// UInt64Options is uint64 options
type UInt64Options struct {
	min uint64
	max uint64
}

func (ch *chance) UInt64(options ...UInt64Option) uint64 {
	ops := UInt64Options{min: 0, max: math.MaxUint64 / 2} // FIXME: Use full range

	for i := range options {
		options[i](&ops)
	}

	ch.r.Seed(ch.seed)
	// TODO: handle error on bad options
	return uint64(randomHelper64(uint64(ops.max-ops.min)) + uint64(ops.min))
}

// UInt64 returns a random uint64
func UInt64(options ...UInt64Option) uint64 {
	return defaultChance.UInt64(options...)
}

// SetUInt64Min sets min of random uint64
func SetUInt64Min(min uint64) UInt64Option {
	return func(iOpts *UInt64Options) {
		iOpts.min = min
	}
}

// SetUInt64Max sets max of random uint64
func SetUInt64Max(max uint64) UInt64Option {
	return func(iOpts *UInt64Options) {
		iOpts.max = max
	}
}
