package chance

import "math"

// UInt32Option is a type
type UInt32Option func(*UInt32Options)

// UInt32Options is uint32 options
type UInt32Options struct {
	min uint32
	max uint32
}

func (ch *chance) UInt32(options ...UInt32Option) uint32 {
	ops := UInt32Options{min: 0, max: math.MaxUint32}

	for i := range options {
		options[i](&ops)
	}

	ch.r.Seed(ch.seed)
	// TODO: handle error on bad options
	return uint32(ch.r.Int63n(int64(ops.max)-int64(ops.min)) + int64(ops.min))
}

// UInt32 returns a random uint32
func UInt32(options ...UInt32Option) uint32 {
	return defaultChance.UInt32(options...)
}

// SetUInt32Min sets min of random uint32
func SetUInt32Min(min uint32) UInt32Option {
	return func(iOpts *UInt32Options) {
		iOpts.min = min
	}
}

// SetUInt32Max sets max of random uint32
func SetUInt32Max(max uint32) UInt32Option {
	return func(iOpts *UInt32Options) {
		iOpts.max = max
	}
}
