package chance

import (
	"math"
	"math/rand"
)

// Float32Option is a type
type Float32Option func(*Float32Options)

// Float32Options is float32 options
type Float32Options struct {
	min float32
	max float32
}

func (ch *chance) Float32(options ...Float32Option) float32 {
	ops := Float32Options{min: -math.MaxFloat32, max: math.MaxFloat32}

	for i := range options {
		options[i](&ops)
	}

	ch.r.Seed(ch.seed)
	// TODO: handle error on bad options
	return ops.min + rand.Float32()*(ops.max-ops.min)
}

// Float32 returns a random float32
func Float32(options ...Float32Option) float32 {
	return defaultChance.Float32(options...)
}

// SetFloat32Min sets min of random float32
func SetFloat32Min(min float32) Float32Option {
	return func(iOpts *Float32Options) {
		iOpts.min = min
	}
}

// SetFloat32Max sets max of random float32
func SetFloat32Max(max float32) Float32Option {
	return func(iOpts *Float32Options) {
		iOpts.max = max
	}
}
