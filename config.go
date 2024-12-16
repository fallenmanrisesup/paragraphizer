package paragraphizer

type Config struct {
	Preprocessors []Preprocessor
}

type Option func(c *Config)

func buildConfig(options []Option) Config {
	c := Config{
		Preprocessors: make([]Preprocessor, 0),
	}

	for _, fn := range options {
		fn(&c)
	}

	return c
}

func WithPreprocessor(s Preprocessor) Option {
	return func(c *Config) {
		c.Preprocessors = append(c.Preprocessors, s)
	}
}
