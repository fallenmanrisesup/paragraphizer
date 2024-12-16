package paragraphizer

import "context"

var (
	_ Preprocessor = &preprocessor{}
)

type Preprocessor interface {
	Preprocess(context.Context, string) (string, error)
}

type PreprocessFunc = func(context.Context, string) (string, error)

type preprocessor struct {
	fn PreprocessFunc
}

func NewPreprocessor(fn PreprocessFunc) Preprocessor {
	return &preprocessor{fn}
}

func (p *preprocessor) Preprocess(ctx context.Context, s string) (string, error) {
	return p.fn(ctx, s)
}
