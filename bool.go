package chance

func (ch *chance) Bool(options ...func(Chance)) bool {
	chp := *ch
	ch2 := &chp

	for i := range options {
		options[i](ch2)
	}

	ch2.r.Seed(ch2.seed)

	return ch2.r.Intn(2) == 1
}

// Bool returns a random boolean
func Bool(options ...func(Chance)) bool {
	return defaultChance.Bool(options...)
}
