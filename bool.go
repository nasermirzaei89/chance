package chance

// BoolOption is a type
type BoolOption func(*BoolOptions)

// BoolOptions is bool options
type BoolOptions struct {
}

func (ch *chance) Bool(options ...BoolOption) bool {
	ops := BoolOptions{}

	for i := range options {
		options[i](&ops)
	}

	ch.r.Seed(ch.seed)

	return ch.r.Intn(2) == 1
}

// Bool returns a random boolean
func Bool(options ...BoolOption) bool {
	return defaultChance.Bool(options...)
}
