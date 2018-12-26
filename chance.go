// Package chance is a random generator in Go
package chance

import "math/rand"

// Chance is interface of chance
type Chance interface {
	Bool(options ...BoolOption) bool
	String(options ...StringOption) string
	Int(options ...IntOption) int
	Int8(options ...Int8Option) int8
	Int16(options ...Int16Option) int16
	Int32(options ...Int32Option) int32
	Int64(options ...Int64Option) int64
}

type chance struct {
	r    *rand.Rand
	seed int64
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
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.seed = 1
	}
}
