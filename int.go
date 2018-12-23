package chance

// IntOption is a type
type IntOption func(*IntOptions)

type IntOptions struct {
}

func (ch *chance) Int(options ...IntOption) int {
	ops := IntOptions{}

	for i := range options {
		options[i](&ops)
	}

	ch2.r.Seed(ch2.seed)
	// TODO: handle error on bad options
	return ch2.r.Intn(ops.intMax-ops.intMin) + ops.intMin
}

// Int returns a random int
func Int(options ...IntOption) int {
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
