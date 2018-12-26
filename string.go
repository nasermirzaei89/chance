package chance

// StringOption is a type
type StringOption func(*StringOptions)

// StringOptions is string options
type StringOptions struct {
	length int
	pool   string
}

func (ch *chance) String(options ...StringOption) string {
	ops := StringOptions{length: 5, pool: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()[]"}

	for i := range options {
		options[i](&ops)
	}

	ch.r.Seed(ch.seed)

	str := ""
	for i := 0; i < ops.length; i++ {
		str += string(ops.pool[ch.r.Intn(len(ops.pool))])
	}

	return str
}

// String returns a random string
func String(options ...StringOption) string {
	return defaultChance.String(options...)
}

// SetStringLength sets length of random string
func SetStringLength(length int) StringOption {
	return func(sch *StringOptions) {
		sch.length = length
	}
}

// SetStringPool sets pool of random string
func SetStringPool(pool string) StringOption {
	return func(sch *StringOptions) {
		sch.pool = pool
	}
}
