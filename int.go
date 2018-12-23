package chance

func (ch *chance) Int(options ...func(Chance)) int {
	chp := *ch
	ch2 := &chp

	for i := range options {
		options[i](ch2)
	}

	ch2.r.Seed(ch2.seed)

	return ch2.r.Int()
}

// Int returns a random int
func Int(options ...func(Chance)) int {
	return defaultChance.Int(options...)
}