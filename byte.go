package chance

import "math"

// ByteOption is a type
type ByteOption func(*ByteOptions)

// ByteOptions is int8 options
type ByteOptions struct {
	min byte
	max byte
}

func (ch *chance) Byte(options ...ByteOption) byte {
	ops := ByteOptions{min: 0, max: math.MaxUint8}

	for i := range options {
		options[i](&ops)
	}

	ch.r.Seed(ch.seed)
	// TODO: handle error on bad options
	return byte(ch.r.Int31n(int32(ops.max)-int32(ops.min)) + int32(ops.min))
}

// Byte returns a random byte
func Byte(options ...ByteOption) byte {
	return defaultChance.Byte(options...)
}

// SetByteMin sets min of random byte
func SetByteMin(min byte) ByteOption {
	return func(iOpts *ByteOptions) {
		iOpts.min = min
	}
}

// SetByteMax sets max of random byte
func SetByteMax(max byte) ByteOption {
	return func(iOpts *ByteOptions) {
		iOpts.max = max
	}
}
