package chance

import "math"

func (ch *chance) Int(options ...func(Chance)) int {
	chp := *ch
	ch2 := &chp

	for i := range options {
		options[i](ch2)
	}

	ch2.r.Seed(ch2.seed)
	// TODO: handle error on bad options
	return ch2.r.Intn(ch2.intMax-ch2.intMin) + ch2.intMin
}

// Int returns a random int
func Int(options ...func(Chance)) int {
	return defaultChance.Int(options...)
}

// SetIntMin sets min of random int
func SetIntMin(min int) Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.intMin = min
	}
}

// SetIntMax sets max of random int
func SetIntMax(max int) Option {
	return func(ich Chance) {
		ch := ich.(*chance)
		ch.intMax = max
	}
}

// ResetIntMax resets max of random int
func ResetIntMax() Option {
	//TODO: set max int64 for 64bit systems
	return SetIntMax(math.MaxInt32)
}

// ResetIntMin resets min of random int
func ResetIntMin() Option {
	//TODO: set max int64 for 64bit systems
	return SetIntMin(math.MaxInt32 * -1)
}
