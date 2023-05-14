package interpreter_eval

import (
	"pika/pkg/ast"
	"pika/pkg/ast/ast_types"
	"pika/pkg/interpreter/interpreter_env"
	"pika/pkg/interpreter/interpreter_makers"
)

func evalCallExpr(expr ast.CallExpr, env interpreter_env.Environment) interpreter_env.RuntimeValue {
	args := make([]interpreter_env.RuntimeValue, len(expr.Args))

	for idx, arg := range expr.Args {
		args[idx] = Evaluate(arg, env)
	}

	fn := Evaluate(expr.Caller, env)

	switch fn.GetType() {
	case interpreter_env.NativeFn:
		result := fn.(interpreter_env.NativeFnVal).Call(args, env)
		return result
	case interpreter_env.Function:
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
			result = Evaluate(statement, scope)
		}

		return result
	}

	panic("Function not found")
}

func evalObjectExpr(objectExpr ast.ObjectLiteral, env interpreter_env.Environment) interpreter_env.RuntimeValue {
	obj := interpreter_env.ObjectVal{
		Type:       interpreter_env.Object,
		Properties: make(map[string]interpreter_env.RuntimeValue),
	}

	for _, property := range objectExpr.Properties {
		key := property.Key
		value := property.Value

		var runtimeValue interpreter_env.RuntimeValue

		if value == nil {
			runtimeValue = env.LookupVar(key)
		} else {
			runtimeValue = Evaluate(value, env)
		}

		obj.Properties[key] = runtimeValue
	}

	return obj
}

func evalAssignment(assignment ast.AssigmentExpr, env interpreter_env.Environment) interpreter_env.RuntimeValue {
	if assignment.Assigne.GetKind() != ast_types.Identifier {
		panic("Invalid assignment target")
	}

	varName := assignment.Assigne.(ast.Identifier).Symbol
	return env.AssignVar(varName, Evaluate(assignment.Value, env))
}

func evalIdentifier(ident ast.Identifier, env interpreter_env.Environment) interpreter_env.RuntimeValue {
	val := env.LookupVar(ident.Symbol)
	return val
}

func evaluateNumericBinaryExpr(operator string, lhs interpreter_env.RuntimeValue, rhs interpreter_env.RuntimeValue) interpreter_env.RuntimeValue {
	result := 0

	valLhs, okLhs := lhs.(interpreter_env.NumberVal)
	valRhs, okRhs := rhs.(interpreter_env.NumberVal)

	if !okLhs || !okRhs {
		panic("Left and right hand side of binary expression must be of type number")
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

	return interpreter_env.NumberVal{Value: result, Type: interpreter_env.Number}
}

func evalBinaryExpr(binop ast.BinaryExpr, env interpreter_env.Environment) interpreter_env.RuntimeValue {
	lhs := Evaluate(binop.Left, env)
	rhs := Evaluate(binop.Right, env)

	if lhs.GetType() == interpreter_env.Number && rhs.GetType() == interpreter_env.Number {
		return evaluateNumericBinaryExpr(binop.Operator, lhs, rhs)
	}

	return interpreter_makers.MK_NULL()
}
