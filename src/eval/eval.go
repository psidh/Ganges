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
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(left, node.Operator, right)
	case *ast.BlockStatement:
		return evalStatements(node.Statements)
	case *ast.IfExpression:
		return evalIfExpression(node)
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

func evalMinusOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
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

func evalInfixExpression(left object.Object, operator string, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	default:
		return NULL
	}
}

func evalIntegerInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftValue + rightValue}
	case "-":
		return &object.Integer{Value: leftValue - rightValue}
	case "*":
		return &object.Integer{Value: leftValue * rightValue}
	case "/":
		return &object.Integer{Value: leftValue / rightValue}
	case "<":
		return nativeBoolToBooleanObject(leftValue < rightValue)
	case ">":
		return nativeBoolToBooleanObject(leftValue > rightValue)
	case ">=":
		return nativeBoolToBooleanObject(leftValue >= rightValue)
	case "<=":
		return nativeBoolToBooleanObject(leftValue <= rightValue)
	case "==":
		return nativeBoolToBooleanObject(leftValue == rightValue)
	case "!=":
		return nativeBoolToBooleanObject(leftValue != rightValue)
	default:
		return NULL
	}
}

func evalIfExpression(ie *ast.IfExpression) object.Object {
	condition := Eval(ie.Condition)

	if isTruthy(condition) {
		return Eval(ie.Consequence)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative)
	} else {
		return NULL
	}
}

func isTruthy(obj object.Object) bool{
	switch obj{
	case NULL:
		return false
	case SATYA:
		return true
	case ASATYA:
		return false
	default:
		return true
	}
}
