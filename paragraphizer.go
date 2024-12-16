package paragraphizer

import "context"

type Paragraphizer interface {
	Paragraphize(context.Context, string) ([]string, error)
}

func NewParagraphizer(opts ...Option) Paragraphizer {
	c := buildConfig(opts)

	return &paragraphizer{c}
}
