package eval

import (
	"testing"

	"github.com/psidh/Ganges/src/lexer"
	"github.com/psidh/Ganges/src/object"
	"github.com/psidh/Ganges/src/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, test := range tests {
		eval := testEval(test.input)
		testIntegerObject(t, eval, test.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	env := object.NewEnvironment()
	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)

	if !ok {
		t.Errorf("object is not Integer. got=%T, (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has incorrect value. got=%d, want=%d", result.Value, expected)
		return false
	}
	return true
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"satya", true},
		{"asatya", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"satya == satya", true},
		{"asatya == asatya", true},
		{"satya == asatya", false},
		{"satya != asatya", true},
		{"asatya != satya", true},
		{"(1 < 2) == satya", true},
		{"(1 < 2) == asatya", false},
		{"(1 > 2) == satya", false},
		{"(1 > 2) == asatya", true},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!satya", false},
		{"!asatya", true},
		{"!5", false},
		{"!!satya", true},
		{"!!asatya", false},
		{"!!5", true},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		testBooleanObject(t, evaluated, test.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)

	if !ok {
		t.Errorf("object is not Boolean got=%T, (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t",
			result.Value, expected)
		return false
	}
	return true
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"yadi (satya) { 10 }", 10},
		{"yadi (asatya) { 10 }", nil},
		{"yadi (1) { 10 }", 10},
		{"yadi (1 < 2) { 10 }", 10},
		{"yadi (1 > 2) { 10 }", nil},
		{"yadi (1 > 2) { 10 } anyatha { 20 }", 20},
		{"yadi (1 < 2) { 10 } anyatha { 20 }", 10},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		integer, ok := test.expected.(int)

		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"daan 10;", 10},
		{"daan 10; 9;", 10},
		{"daan 2 * 5; 9;", 10},
		{"9; daan 2 * 5; 9;", 10},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input                string
		expectedErrorMessage string
	}{
		{"5 + satya;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{"5 + satya; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-satya",
			"unknown operator: -BOOLEAN",
		},
		{
			"satya + asatya;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; satya + asatya; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"yadi (10 > 1) { satya + asatya; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{`yadi (10 > 1) {
			yadi (10 > 1) {
			daan satya + asatya;
			}
			daan 1;
			}`,
			"unknown operator: BOOLEAN + BOOLEAN",
		}, {
			"foobar",
			"identifier not found: foobar",
		},
		{
			`"Hello" - "World"`,
			"unknown operator: STRING - STRING",
		},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)

		errorObject, ok := evaluated.(*object.Error)

		if !ok {
			t.Errorf("no error object returned. got=%T (%+v)", evaluated, evaluated)
			continue
		}

		if errorObject.Message != test.expectedErrorMessage {
			t.Errorf("wrong error message. Expected Message=%q, got=%q", test.expectedErrorMessage, errorObject.Message)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
	}

	for _, test := range tests {
		testIntegerObject(t, testEval(test.input), test.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "kriya(x) { x + 2;};"

	evaluated := testEval(input)

	kriya, ok := evaluated.(*object.Function)

	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if (len(kriya.Parameters)) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v",
			kriya.Parameters)
	}
	if kriya.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", kriya.Parameters[0])
	}

	expectedBody := "(x + 2)"

	if kriya.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, kriya.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let identity = kriya(x) { x; }; identity(5);", 5},
		{"let identity = kriya(x) { daan x; }; identity(5);", 5},
		{"let double = kriya(x) { x * 2; }; double(5);", 10},
		{"let add = kriya(x, y) { x + y; }; add(5, 5);", 10},
		{"let add = kriya(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"kriya(x) { x; }(5)", 5},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
		let newAdder = kriya(x) {
		kriya(y) { x + y };
		};
		let addTwo = newAdder(2);
		addTwo(2);
	`

	testIntegerObject(t, testEval(input), 4)
}

func TestStringLiteral(t *testing.T) {
	input := `"ram ram"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	println("_____________________________")
	println(str)
	println("_____________________________")
	if !ok {
		t.Fatalf("object is not String. got=%T, (%+v)", evaluated, evaluated)
	}

	if str.Value != "ram ram" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"ram" + " siya " + "ram"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	println("_____________________________")
	println(str)
	println("_____________________________")
	if !ok {
		t.Fatalf("object is not String. got=%T, (%+v)", evaluated, evaluated)
	}

	if str.Value != "ram siya ram" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{

		{`dairghya("")`, 0},
		{`dairghya("four")`, 4},
		{`dairghya("hello world")`, 11},
		{`dairghya(1)`, "argument to `dairghya` not supported, got INTEGER"},
		{`dairghya("one", "two")`, "wrong number of arguments. got=2, want=1"},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)

		switch expected := test.expected.(type) {

		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)

			if !ok {
				t.Errorf("object is not Error. got=%T, (%+v)", evaluated, evaluated)
				continue
			}

			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Message)
			}
		}
	}
}
