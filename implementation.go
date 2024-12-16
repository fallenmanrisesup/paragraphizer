package paragraphizer

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

var (
	_ Paragraphizer = &paragraphizer{}
)

type paragraphizer struct {
	c Config
}

func (p *paragraphizer) Paragraphize(ctx context.Context, doc string) ([]string, error) {
	pr, err := newTokenProcessor(ctx, p.c, doc)
	if err != nil {
		return nil, err
	}

	return pr.Process()
}

type tokenProcessor struct {
	tokenizer        *html.Tokenizer
	currentParagraph bytes.Buffer
	paragraphs       []string
	handlers         map[html.TokenType]func(html.Token)
}

func newTokenProcessor(ctx context.Context, c Config, doc string) (*tokenProcessor, error) {
	var err error
	for _, s := range c.Preprocessors {
		doc, err = s.Preprocess(ctx, doc)
		if err != nil {
			return nil, err
		}
	}

	processor := &tokenProcessor{
		tokenizer: html.NewTokenizer(bytes.NewReader([]byte(doc))),
		handlers:  make(map[html.TokenType]func(html.Token)),
	}

	processor.handlers = map[html.TokenType]func(html.Token){
		html.StartTagToken: processor.handleStartTag,
		html.EndTagToken:   processor.handleEndTag,
		html.TextToken:     processor.handleTextToken,
	}

	return processor, nil
}

func (p *tokenProcessor) Process() ([]string, error) {
	for {
		tokenType := p.tokenizer.Next()
		if tokenType == html.ErrorToken {
			return finalizeParagraphs(&p.paragraphs, &p.currentParagraph)
		}

		if handler, exists := p.handlers[tokenType]; exists {
			token := p.tokenizer.Token()
			handler(token)
		}
	}
}

func (p *tokenProcessor) handleStartTag(token html.Token) {
	if isAllowedInline(token.Data) {
		p.currentParagraph.WriteString(formatStartTag(token))
	} else if isBlockTag(token.Data) {
		flushCurrentParagraph(&p.paragraphs, &p.currentParagraph)
	}
}

func (p *tokenProcessor) handleEndTag(token html.Token) {
	if isAllowedInline(token.Data) {
		p.currentParagraph.WriteString(fmt.Sprintf("</%s>", token.Data))
	} else if isBlockTag(token.Data) {
		flushCurrentParagraph(&p.paragraphs, &p.currentParagraph)
	}
}

func (p *tokenProcessor) handleTextToken(token html.Token) {
	p.currentParagraph.WriteString(token.Data)
}

func finalizeParagraphs(paragraphs *[]string, currentParagraph *bytes.Buffer) ([]string, error) {
	flushCurrentParagraph(paragraphs, currentParagraph)
	return *paragraphs, nil
}

func flushCurrentParagraph(paragraphs *[]string, currentParagraph *bytes.Buffer) {
	if currentParagraph.Len() > 0 && strings.TrimSpace(currentParagraph.String()) != "" {
		*paragraphs = append(*paragraphs, fmt.Sprintf("<p>%s</p>", strings.TrimSpace(currentParagraph.String())))
		currentParagraph.Reset()
	}
}
