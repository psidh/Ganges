package eval

import (
	"github.com/psidh/Ganges/src/ast"
	"github.com/psidh/Ganges/src/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	}
	return nil
}

func evalStatements(statements []ast.Statement) object.Object {

	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)
	}
	return result

}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusOperatorExpression(right)
	default:
		return NULL
	}
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return SATYA
	}
	return ASATYA
}

func evalBangOperatorExpression(right object.Object) object.Object {

	switch right {
	case SATYA:
		return ASATYA
	case ASATYA:
		return SATYA
	case NULL:
		return SATYA
	default:
		return ASATYA
	}
}

func evalMinusOperatorExpression(right object.Object) object.Object{
	if right.Type() != object.INTEGER_OBJ{
		return NULL
	}

	literalValue := right.(*object.Integer).Value
	return &object.Integer{Value: literalValue * -1}
}

var (
	NULL   = &object.Null{}
	SATYA  = &object.Boolean{Value: true}
	ASATYA = &object.Boolean{Value: false}
)
