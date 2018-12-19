package chance

import "math/rand"

// Chance is interface of chance
type Chance interface {
	Bool(options ...func(Chance)) bool
	Character(options ...func(Chance)) string
}

type chance struct {
	seed int64
	r    *rand.Rand
	// Bool
	likelihood int64
	// Character
	pool    *string
	alpha   bool
	casing  string
	numeric bool
	symbols bool
}

// Casing
const (
	CasingLower = "lower"
	CasingUpper = "upper"
)

// Characters
const (
	CharsLower = "abcdefghijklmnopqrstuvwxyz"
	CharsUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers    = "0123456789"
	Symbols    = "!@#$%^&*()"
)

// Option is a type
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

	ResetLikelihood()(ch)

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
