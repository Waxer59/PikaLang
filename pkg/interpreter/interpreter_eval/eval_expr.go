package interpreter_eval

import (
	"errors"
	"fmt"
	"math"
	compilerErrors "pika/internal/errors"
	"pika/pkg/ast"
	"pika/pkg/ast/ast_types"
	"pika/pkg/interpreter/interpreter_env"
	"pika/pkg/interpreter/interpreter_makers"
	"pika/pkg/interpreter/interpreter_utils"

	"golang.org/x/exp/slices"
)

func evalCallExpr(expr ast.CallExpr, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	args := make([]interpreter_env.RuntimeValue, len(expr.Args))

	for idx, arg := range expr.Args {
		eval, err := Evaluate(arg, env)
		if err != nil {
			return nil, err
		}
		args[idx] = eval
	}

	fn, err := Evaluate(expr.Caller, env)
	fnName := expr.Caller.(ast.Identifier).Symbol
	nativeFn, nativeFnOk := interpreter_utils.IsNativeFunction(fnName)

	if err != nil && nativeFnOk {
		result := nativeFn(args, env)
		return result, nil
	}

	if err != nil || fn.GetType() != interpreter_env.Function {
		return nil, errors.New(compilerErrors.ErrFuncNotFound + fnName)
	}

	function := fn.(interpreter_env.FunctionVal)
	scope := interpreter_env.New(function.DeclarationEnv)

	// Create the variables for the function arguments
	for idx, arg := range function.Params {
		scope.DeclareVar(arg, args[idx], false)
	}

	var result interpreter_env.RuntimeValue = interpreter_makers.MK_NULL()

	// Evaluate the function body line by line
	for _, statement := range function.Body {
		eval, err := Evaluate(statement, scope)
		if err != nil {
			return nil, err
		}
		result = eval
	}

	return result, nil

}

func evalMemberExpr(expr ast.MemberExpr, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	property := expr.Property

	evalObj, err := Evaluate(expr.Object, env)
	if err != nil {
		return nil, err
	}

	valObj := evalObj.GetValue()
	if expr.Computed {
		evalProperty, err := Evaluate(property, env)
		if err != nil {
			return nil, err
		}
		switch obj := valObj.(type) {
		case []interpreter_env.RuntimeValue:
			if evalProperty.GetType() != interpreter_env.Number {
				return nil, errors.New(compilerErrors.ErrIndexNotFound)
			}
			if int(evalProperty.GetValue().(float64)) >= len(obj) {
				return nil, errors.New(compilerErrors.ErrIndexNotFound)
			}
			val := obj[int(evalProperty.GetValue().(float64))]
			return val, nil
		case map[string]interpreter_env.RuntimeValue:
			valProperty := fmt.Sprint(evalProperty.GetValue())
			if _, ok := obj[valProperty]; ok {
				return obj[valProperty], nil
			} else {
				return nil, errors.New(compilerErrors.ErrPropertyNotFound)
			}
		default:
			return nil, errors.New(compilerErrors.ErrIndexNotFound)
		}
	}

	valProperty := fmt.Sprint(property.(ast.Identifier).Symbol)
	val, ok := valObj.(map[string]interpreter_env.RuntimeValue)[valProperty]
	if !ok {
		return nil, errors.New(compilerErrors.ErrPropertyNotFound)
	}
	return val, nil
}

func evalArrayExpr(arrayExpr ast.ArrayLiteral, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	arr := interpreter_env.ArrayVal{
		Type:     interpreter_env.Array,
		Elements: make([]interpreter_env.RuntimeValue, len(arrayExpr.Elements)),
	}

	for idx, element := range arrayExpr.Elements {
		eval, err := Evaluate(element, env)

		if err != nil {
			return nil, err
		}

		if element.GetKind() == ast_types.Identifier {
			eval, err = env.LookupVar(eval.GetValue().(string))

			if err != nil {
				return nil, err
			}

		}

		arr.Elements[idx] = eval
	}

	return arr, nil
}

func evalObjectExpr(objectExpr ast.ObjectLiteral, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	obj := interpreter_env.ObjectVal{
		Type:       interpreter_env.Object,
		Properties: make(map[string]interpreter_env.RuntimeValue),
	}

	for _, property := range objectExpr.Properties {
		key := property.Key
		value := property.Value

		var runtimeValue interpreter_env.RuntimeValue
		var err error

		if value == nil {
			runtimeValue, err = env.LookupVar(key)
		} else {
			runtimeValue, err = Evaluate(value, env)
		}

		if err != nil {
			return nil, err
		}

		obj.Properties[key] = runtimeValue
	}

	return obj, nil
}

func evalAssignment(assignment ast.AssigmentExpr, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	if assignment.Assigne.GetKind() != ast_types.Identifier {
		return nil, errors.New(compilerErrors.ErrSyntaxInvalidAssignment)
	}

	varName := assignment.Assigne.(ast.Identifier).Symbol
	eval, err := Evaluate(assignment.Value, env)

	if err != nil {
		return nil, err
	}

	variable, err := env.AssignVar(varName, eval)

	return variable, err
}

func evalIdentifier(ident ast.Identifier, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	val, err := env.LookupVar(ident.Symbol)
	return val, err
}

func evalConditionalExpr(conditionalExpr ast.ConditionalExpr, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	evalCondition, err := Evaluate(conditionalExpr.Condition, env)

	if err != nil {
		return nil, err
	}

	val := EvaluateTruthyFalsyValues(evalCondition)

	if val {
		return Evaluate(conditionalExpr.Consequent, env)
	}

	return Evaluate(conditionalExpr.Alternate, env)
}

func evalStringBinaryExpr(operator string, lhs interpreter_env.RuntimeValue, rhs interpreter_env.RuntimeValue) (interpreter_env.RuntimeValue, error) {
	var result string = ""
	valLhs, okLhs := lhs.(interpreter_env.StringVal)
	valRhs, okRhs := rhs.(interpreter_env.StringVal)
	if !okLhs || !okRhs {
		return nil, errors.New(compilerErrors.ErrBinaryInvalidBinaryExpr)
	}
	switch operator {
	case "+":
		result = valLhs.Value + valRhs.Value
	}

	return interpreter_makers.MK_String(result), nil
}

func evaluateNumericBinaryExpr(operator string, lhs interpreter_env.RuntimeValue, rhs interpreter_env.RuntimeValue) (interpreter_env.RuntimeValue, error) {
	var result float64 = 0

	valLhs, okLhs := lhs.(interpreter_env.NumberVal)
	valRhs, okRhs := rhs.(interpreter_env.NumberVal)

	if !okLhs || !okRhs {
		return nil, errors.New(compilerErrors.ErrBinaryInvalidBinaryExpr)
	}

	switch operator {
	case "+":
		result = valLhs.Value + valRhs.Value
	case "-":
		result = valLhs.Value - valRhs.Value
	case "*":
		result = valLhs.Value * valRhs.Value
	case "/":
		if valRhs.Value == 0 {
			result = math.Inf(0)
		}
		result = valLhs.Value / valRhs.Value
	case "%":
		if valRhs.Value == 0 {
			return nil, errors.New(compilerErrors.ErrBinaryDivisionByZero)
		}
		result = float64(int(valLhs.Value) % int(valRhs.Value))
	case "**", "^":
		result = math.Pow(valLhs.Value, valRhs.Value)
	}

	return interpreter_makers.MK_Number(result), nil
}

func evalLogicalExpr(logicalExpr ast.LogicalExpr, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	result := false
	evalLhs, err := Evaluate(logicalExpr.Left, env)

	if err != nil {
		return nil, err
	}

	evalRhs, err := Evaluate(logicalExpr.Right, env)

	if err != nil {
		return nil, err
	}

	valLhs := EvaluateTruthyFalsyValues(evalLhs.GetValue())
	valRhs := EvaluateTruthyFalsyValues(evalRhs.GetValue())

	switch logicalExpr.Operator {
	case "&&":
		result = valLhs && valRhs
	case "||":
		result = valLhs || valRhs
	}

	return interpreter_makers.MK_Boolean(result), nil
}

func evalUnaryExpr(expr ast.UnaryExpr, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	result := false
	eval, err := Evaluate(expr.Argument, env)

	if err != nil {
		return nil, err
	}

	boolVal := EvaluateTruthyFalsyValues(eval.GetValue())

	switch expr.Operator {
	case "!":
		result = !boolVal
	}

	return interpreter_makers.MK_Boolean(result), nil
}

func evalComparisonBinaryExpr(operator string, lhs interpreter_env.RuntimeValue, rhs interpreter_env.RuntimeValue) (interpreter_env.RuntimeValue, error) {
	var result bool = false

	numValLhs, _ := lhs.(interpreter_env.NumberVal)
	numValRhs, _ := rhs.(interpreter_env.NumberVal)

	switch operator {
	case "==":
		result = lhs.GetValue() == rhs.GetValue()
	case "!=":
		result = lhs.GetValue() != rhs.GetValue()
	case "<":
		result = numValLhs.Value < numValRhs.Value
	case ">":
		result = numValLhs.Value > numValRhs.Value
	case "<=":
		result = numValLhs.Value <= numValRhs.Value
	case ">=":
		result = numValLhs.Value >= numValRhs.Value
	}
	return interpreter_makers.MK_Boolean(result), nil
}

func evalBinaryExpr(binop ast.BinaryExpr, env interpreter_env.Environment) (interpreter_env.RuntimeValue, error) {
	lhs, err := Evaluate(binop.Left, env)
	if err != nil {
		return nil, err
	}

	rhs, err := Evaluate(binop.Right, env)
	if err != nil {
		return nil, err
	}

	// EVAL < < >= <= == !=
	if slices.Contains(ast_types.BoolExpr, binop.Operator) {
		eval, err := evalComparisonBinaryExpr(binop.Operator, lhs, rhs)
		return eval, err
	}

	// EVAL + - * / % ** (numbers)
	if lhs.GetType() == interpreter_env.Number && rhs.GetType() == interpreter_env.Number {
		eval, err := evaluateNumericBinaryExpr(binop.Operator, lhs, rhs)
		return eval, err
	}

	// EVAL + (strings)
	if lhs.GetType() == interpreter_env.String && rhs.GetType() == interpreter_env.String {
		eval, err := evalStringBinaryExpr(binop.Operator, lhs, rhs)
		return eval, err
	}

	return interpreter_makers.MK_NULL(), nil
}
