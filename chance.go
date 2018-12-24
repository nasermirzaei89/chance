// Package chance is a random generator in Go
package chance

import "math/rand"

// Chance is interface of chance
type Chance interface {
	Bool(options ...func(Chance)) bool
	String(options ...func(Chance)) string
	Int(options ...func(Chance)) int
}

type chance struct {
	r    *rand.Rand
	seed int64

	stringLength int
	stringPool   string
	intMin       int
	intMax       int
}

// Option is a type
type Option func(Chance)

var defaultChance Chance

func init() {
	defaultChance = New()
}

// New returns new instance of chance
//  chance.New(chance.SetSeed(time.Now().UnixNano()))
func New(options ...Option) Chance {
	ch := new(chance)

	// defaults
	ResetSeed()(ch)
	ch.r = rand.New(rand.NewSource(ch.seed))
	ResetStringLength()(ch)
	ResetStringPool()(ch)
	ResetIntMax()(ch)
	ResetIntMin()(ch)

	for i := range options {
		options[i](ch)
	}
	return ch
}

// SetSeed sets seed option of instance
func SetSeed(seed int64) Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.seed = seed
	}
}

// ResetSeed resets seed option of instance
func ResetSeed() Option {
	return SetSeed(1)
}
