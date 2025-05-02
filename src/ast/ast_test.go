package ast

import (
	"testing"

	"github.com/psidh/Ganges/src/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "krishna"},
					Value: "krishna",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "balaram"},
					Value: "balaram",
				},
			},
		},
	}
	if program.String() != "let krishna = balaram;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
