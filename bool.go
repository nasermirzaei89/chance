package chance

func (ch *chance) Bool(options ...func(Chance)) bool {
	chp := *ch
	ch2 := &chp

	for i := range options {
		options[i](ch2)
	}

	if ch2.likelihood <= 0 {
		return false
	}

	if ch2.likelihood >= 100 {
		return true
	}

	ch2.r.Seed(ch2.seed)

	return ch2.r.Int63n(100) < ch2.likelihood
}

// Bool returns a random boolean
func Bool(options ...func(Chance)) bool {
	return defaultChance.Bool(options...)
}

// SetLikelihood sets likelihood for bool
func SetLikelihood(likelihood int64) Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.likelihood = likelihood
	}
}

// ResetLikelihood resets likelihood for bool
func ResetLikelihood() Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.likelihood = 50
	}
}
