package chance

import "math"

// UInt8Option is a type
type UInt8Option func(*UInt8Options)

// UInt8Options is uint8 options
type UInt8Options struct {
	min uint8
	max uint8
}

func (ch *chance) UInt8(options ...UInt8Option) uint8 {
	ops := UInt8Options{min: 0, max: math.MaxUint8}

	for i := range options {
		options[i](&ops)
	}

	ch.r.Seed(ch.seed)
	// TODO: handle error on bad options
	return uint8(ch.r.Int31n(int32(ops.max)-int32(ops.min)) + int32(ops.min))
}

// UInt8 returns a random uint8
func UInt8(options ...UInt8Option) uint8 {
	return defaultChance.UInt8(options...)
}

// SetUInt8Min sets min of random uint8
func SetUInt8Min(min uint8) UInt8Option {
	return func(iOpts *UInt8Options) {
		iOpts.min = min
	}
}

// SetUInt8Max sets max of random uint8
func SetUInt8Max(max uint8) UInt8Option {
	return func(iOpts *UInt8Options) {
		iOpts.max = max
	}
}
