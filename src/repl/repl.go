package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/psidh/Ganges/src/eval"
	"github.com/psidh/Ganges/src/lexer"
	"github.com/psidh/Ganges/src/object"
	"github.com/psidh/Ganges/src/parser"
)

const PROMPT = "वदः >> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	const SWASTIKA = `
	..       ..________
	||       ||
	||       ||
	||       ||
	||_______||________
                 ||      ||
                 ||      ||
                 ||      ||
	_________||      ||
	`

	// fmt.Println(SWASTIKA)
	io.WriteString(out, SWASTIKA)
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

		if evaluated != nil && evaluated.Type() != object.NULL_OBJ {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")

		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "\nत्रुटियाँ प्राप्त हुईं 👇\n")
	for _, msg := range errors {
		io.WriteString(out, "\t❗ दोषः: "+translateToShuddhHindi(msg)+"\n")
	}
	io.WriteString(out, "\nकृपया सुधार करके पुनः प्रयास करें।\n")
}

func translateToShuddhHindi(msg string) string {
	replacements := map[string]string{
		"Expected next token to be":    "अगला चिह्न होने की अपेक्षा थी",
		"instead got":                  "इसके बजाय मिला",
		"no prefix parse function for": "कोई प्रारंभिक व्याख्या कार्य नहीं मिला",
		"could not parse":              "व्याख्या करने में असमर्थ",
		"as integer":                   "पूर्णांक के रूप में",
	}

	for eng, hindi := range replacements {
		msg = replaceAllIgnoreCase(msg, eng, hindi)
	}

	return msg
}

func replaceAllIgnoreCase(s, old, new string) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(s, strings.ToLower(old), new),
		strings.ToUpper(old), new)
}
