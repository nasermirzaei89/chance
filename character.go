package chance

func (ch *chance) Character(options ...func(Chance)) string {
	chp := *ch
	ch2 := &chp

	for i := range options {
		options[i](ch2)
	}

	pool := ""
	if ch2.pool != nil {
		pool = *ch2.pool
	} else {
		letters := ""

		if ch2.casing == CasingLower {
			letters += CharsLower
		} else if ch2.casing == CasingUpper {
			letters += CharsUpper
		} else {
			letters += CharsLower + CharsUpper
		}

		if ch2.alpha {
			pool += letters
		}

		if ch2.numeric {
			pool += Numbers
		}

		if ch2.symbols {
			pool += Symbols
		}

		if pool == "" {
			pool = letters + Numbers + Symbols
		}
	}

	return string(pool[ch2.r.Intn(len(pool))])
}

// Character returns a random character
func Character(options ...func(Chance)) string {
	return defaultChance.Character(options...)
}

// SetPool sets pool
func SetPool(pool string) Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.pool = &pool
	}
}

// ResetPool resets pool
func ResetPool() Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.pool = nil
	}
}

// SetAlpha sets whether use alphabets or not
func SetAlpha(use bool) Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.alpha = use
	}
}

// SetUseUpperCase sets letters casing to upper
func SetUseUpperCase() Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.casing = CasingUpper
	}
}

// SetUseLowerCase sets letters casing to lower
func SetUseLowerCase() Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.casing = CasingLower
	}
}

// ResetCasing sets letters casing to both upper and lower
func ResetCasing() Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.casing = ""
	}
}

// SetNumeric sets whether use numbers or not
func SetNumeric(use bool) Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.numeric = use
	}
}

// SetSymbols sets whether use symbols or not
func SetSymbols(use bool) Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.symbols = use
	}
}