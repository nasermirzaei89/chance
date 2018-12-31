package chance

import "math"

// UInt16Option is a type
type UInt16Option func(*UInt16Options)

// UInt16Options is uint16 options
type UInt16Options struct {
	min uint16
	max uint16
}

func (ch *chance) UInt16(options ...UInt16Option) uint16 {
	ops := UInt16Options{min: 0, max: math.MaxUint16}

	for i := range options {
		options[i](&ops)
	}

	ch.r.Seed(ch.seed)
	// TODO: handle error on bad options
	return uint16(ch.r.Int31n(int32(ops.max) - int32(ops.min) + int32(ops.min)))
}

// UInt16 returns a random uint16
func UInt16(options ...UInt16Option) uint16 {
	return defaultChance.UInt16(options...)
}

// SetUInt16Min sets min of random uint16
func SetUInt16Min(min uint16) UInt16Option {
	return func(iOpts *UInt16Options) {
		iOpts.min = min
	}
}

// SetUInt16Max sets max of random uint16
func SetUInt16Max(max uint16) UInt16Option {
	return func(iOpts *UInt16Options) {
		iOpts.max = max
	}
}
