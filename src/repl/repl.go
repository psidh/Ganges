package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/psidh/Ganges/src/eval"
	"github.com/psidh/Ganges/src/lexer"
	"github.com/psidh/Ganges/src/object"
	"github.com/psidh/Ganges/src/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")

		evaluated := eval.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "\nParser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t- "+msg+"\n")
	}
}
