package chance

// BoolOption is a type
type BoolOption func(*BoolOptions)

type BoolOptions struct {
}

func (ch *chance) Bool(options ...BoolOption) bool {
	ops := BoolOptions{}

	for i := range options {
		options[i](&ops)
	}

	ch.R.Seed(ch.Seed)

	return ch.R.Intn(2) == 1
}

// Bool returns a random boolean
func Bool(options ...BoolOption) bool {
	return defaultChance.Bool(options...)
}
