package chance

import "math/rand"

// Chance is interface of chance
type Chance interface {
	Bool(options ...func(Chance)) bool
}

type chance struct {
	seed int64
	r    *rand.Rand
	// Bool
	likelihood int64
}

type Option func(Chance)

var defaultChance Chance

func init() {
	defaultChance = New()
}

// New returns new instance of chance
func New(options ...Option) Chance {
	ch := new(chance)

	// defaults
	ch.seed = 1
	ch.r = rand.New(rand.NewSource(ch.seed))

	DefaultLikelihood()(ch)

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
