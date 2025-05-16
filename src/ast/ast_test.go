package ast

import (
	"testing"

	"github.com/psidh/Ganges/src/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&RamaStatement{
				Token: token.Token{Type: token.RAMA, Literal: "rama"},
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
	if program.String() != "rama krishna = balaram;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
