# Paragraphizer

Paragraphizer — tool for splitting html document blocks into slice of paragraphs

# Install package

```bash
go get github.com/fallenmanrisesup/paragraphizer
```

# Example

```go
package main

import (
	"context"
	"fmt"
	"html"
	"log"
	"paragraphizer"
)

const doc = `
	<div>
		<a href="https://abc.com">link out of paragraph</a>
		<p>next p</p>
		<div>
			<div>
				<i>nested inline elem</i>
			</div>
		</div>
	</div>
`

func main() {
	var (
		unescapePreprocessor = paragraphizer.NewPreprocessor(
			func(ctx context.Context, s string) (string, error) {
				return html.UnescapeString(s), nil
			},
		)
		xssPreprocessor = paragraphizer.NewPreprocessor(
			func(ctx context.Context, s string) (string, error) {
				// remove xss tags
				return s, nil
			},
		)
	)

	p := paragraphizer.NewParagraphizer(
		paragraphizer.WithPreprocessor(
			unescapePreprocessor,
		),
		paragraphizer.WithPreprocessor(
			xssPreprocessor,
		),
	)

	result, err := p.Paragraphize(context.Background(), doc)
	if err != nil {
		log.Fatal(err)
	}

	// [0] <p><a href="https://abc.com">link out of paragraph</a></p>
	// [1] <p>next p</p>
	// [2] <p><i>nested inline elem</i></p>

	for i, r := range result {
		fmt.Printf("[%d] %s\n", i, r)
	}
}
```
