package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/psidh/Ganges/src/lexer"
	"github.com/psidh/Ganges/src/token"
)

const PROMPT = "वदः >>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

	}
}
