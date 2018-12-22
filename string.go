package chance

func (ch *chance) String(options ...func(Chance)) string {
	chp := *ch
	ch2 := &chp

	for i := range options {
		options[i](ch2)
	}

	ch2.r.Seed(ch2.seed)

	str := ""

	for i := 0; i < ch2.stringLength; i++ {
		str += string(ch2.stringPool[ch2.r.Intn(len(ch2.stringPool))])
	}

	return str
}

// String returns a random string
func String(options ...func(Chance)) string {
	return defaultChance.String(options...)
}

// SetStringLength sets length of random string
func SetStringLength(length int) Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.stringLength = length
	}
}

// ResetStringLength resets length of random string
func ResetStringLength() Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.stringLength = 5
	}
}

// SetStringPool sets pool of random string
func SetStringPool(pool string) Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.stringPool = pool
	}
}

// ResetStringPool resets pool of random string
func ResetStringPool() Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.stringPool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()[]"
	}
}
