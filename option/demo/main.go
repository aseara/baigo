package demo

// Option 可配选项
type Option func(*Foo) Option

// Verbosity option for verbosity
func Verbosity(v int) Option {
	return func(f *Foo) Option {
		previous := f.verbosity
		f.verbosity = v
		return Verbosity(previous)
	}
}

// Foo type for demo
type Foo struct {
	verbosity int
}

// Option apply opts to f
func (f *Foo) Option(opts ...Option) Option {
	l := len(opts)
	if l == 0 {
		return func(f *Foo) Option {
			return nil
		}
	}
	r := make([]Option, l)
	for _, opt := range opts {
		l--
		r[l] = opt(f)
	}
	return func(f *Foo) Option {
		for _, opt := range r {
			_ = opt(f)
		}
		return nil
	}
}
