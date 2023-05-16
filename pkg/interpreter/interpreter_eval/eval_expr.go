package interpreter_eval

import (
	"errors"
	compilerErrors "pika/internal/errors"
	"pika/pkg/ast"
	"pika/pkg/ast/ast_types"
	"pika/pkg/interpreter/interpreter_env"
	"pika/pkg/interpreter/interpreter_makers"
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
	nativeFn, nativeFnOk := interpreter_env.IsNativeFunction(fnName)

	if err != nil && nativeFnOk {
		result := nativeFn(args, env)
		return result, nil
	}

	if err != nil || fn.GetType() != interpreter_env.Function {
		return nil, errors.New(string(compilerErrors.ErrFuncNotFound) + fnName)
	}

	function := fn.(interpreter_env.FunctionVal)
	scope := interpreter_env.New(function.DeclarationEnv)

	// Create the variables for the function arguments
	for idx, arg := range function.Params {
		//TODO: Check the bounds | verify arity of function
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
		return nil, errors.New(string(compilerErrors.ErrSyntaxInvalidAssignment))
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

func evaluateNumericBinaryExpr(operator string, lhs interpreter_env.RuntimeValue, rhs interpreter_env.RuntimeValue) (interpreter_env.RuntimeValue, error) {
	result := 0

	valLhs, okLhs := lhs.(interpreter_env.NumberVal)
	valRhs, okRhs := rhs.(interpreter_env.NumberVal)

	if !okLhs || !okRhs {
		return nil, errors.New(string(compilerErrors.ErrSyntaxInvalidBinaryExpr))
	}

	switch operator {
	case "+":
		result = valLhs.Value + valRhs.Value
	case "-":
		result = valLhs.Value - valRhs.Value
	case "*":
		result = valLhs.Value * valRhs.Value
	case "/":
		//TODO: Division by zero checks
		result = valLhs.Value / valRhs.Value
	case "%":
		result = valLhs.Value % valRhs.Value
	}

	return interpreter_env.NumberVal{Value: result, Type: interpreter_env.Number}, nil
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

	if lhs.GetType() == interpreter_env.Number && rhs.GetType() == interpreter_env.Number {
		eval, err := evaluateNumericBinaryExpr(binop.Operator, lhs, rhs)
		return eval, err
	}

	return interpreter_makers.MK_NULL(), nil
}
