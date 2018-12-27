package chance

import (
	"math"
	"math/rand"
	"strconv"
)

// UIntOption is a type
type UIntOption func(*UIntOptions)

// UIntOptions is uint options
type UIntOptions struct {
	min uint
	max uint
}

func (ch *chance) UInt(options ...UIntOption) uint {
	ops := UIntOptions{}

	for i := range options {
		options[i](&ops)
	}

	ch.r.Seed(ch.seed)
	// TODO: handle error on bad options

	// Check machine architecture
	if strconv.IntSize == 32 {
		return uint(randomHelper32(uint32(ops.max-ops.min)) + uint32(ops.min))
	}

	return uint(randomHelper64(uint64(ops.max-ops.min)) + uint64(ops.min))
}

// UInt returns a random uint
func UInt(options ...UIntOption) uint {
	return defaultChance.UInt(options...)
}

// SetUIntMin sets min of random uint
func SetUIntMin(min uint) UIntOption {
	return func(iOpts *UIntOptions) {
		iOpts.min = min
	}
}

// SetUIntMax sets max of random uint
func SetUIntMax(max uint) UIntOption {
	return func(iOpts *UIntOptions) {
		iOpts.max = max
	}
}

func randomHelper64(n uint64) uint64 {
	if n < math.MaxInt64 {
		return uint64(rand.Int63n(int64(n + 1)))
	}
	x := rand.Uint64()
	for x > n {
		x = rand.Uint64()
	}
	return x
}

func randomHelper32(n uint32) uint32 {
	if n < math.MaxInt32 {
		return uint32(rand.Int31n(int32(n + 1)))
	}
	x := rand.Uint32()
	for x > n {
		x = rand.Uint32()
	}
	return x
}
