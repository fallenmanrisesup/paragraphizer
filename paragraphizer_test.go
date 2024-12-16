package paragraphizer_test

import (
	"context"
	"fmt"
	"paragraphizer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParagraphize(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input: `
				<p>Лесничему, нашедшему главную новогоднюю елку страны, вручили новый автомобиль Niva Chevrolet. Машина будет использоваться только для работы, <a href="https://ria.ru/20241209/elka-1988128168.html">сообщает</a> «РИА Новости».</p> 
				<div class="article__special_container"> 
					<p>«Она будет исключительно рабочая, в целях осуществления своих функциональных обязанностей», — сказал журналистам директор Бородинского филиала ГКУ МО «Мособллес», лесничий Махач Ширинов.</p> 
				</div>
			`,
			expected: []string{
				`<p>Лесничему, нашедшему главную новогоднюю елку страны, вручили новый автомобиль Niva Chevrolet. Машина будет использоваться только для работы, <a href="https://ria.ru/20241209/elka-1988128168.html">сообщает</a> «РИА Новости».</p>`,

				`<p>«Она будет исключительно рабочая, в целях осуществления своих функциональных обязанностей», — сказал журналистам директор Бородинского филиала ГКУ МО «Мособллес», лесничий Махач Ширинов.</p>`,
			},
		},
		{
			input: `
				<div>
					<a href="https://abc.com">link out of paragraph</a>
					<p>next p</p>
					<div>
						<div>
							<i>nested inline elem</i>
						</div>
					</div>
				</div>
			`,
			expected: []string{
				`<p><a href="https://abc.com">link out of paragraph</a></p>`,
				`<p>next p</p>`,
				`<p><i>nested inline elem</i></p>`,
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("[%d/%d]", i+1, len(tests)), func(t *testing.T) {
			h := paragraphizer.NewParagraphizer()
			result, err := h.Paragraphize(context.Background(), tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
