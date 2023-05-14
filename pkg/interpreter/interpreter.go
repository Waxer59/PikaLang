package interpreter

import (
	"pika/pkg/ast"
	"pika/pkg/ast/ast_types"
)

func evalBinaryExpr(binop ast.BinaryExpr, env Environment) RuntimeValue {
	lhs := Evaluate(binop.Left, env)
	rhs := Evaluate(binop.Right, env)

	if lhs.GetType() == Number && rhs.GetType() == Number {
		return evaluateNumericBinaryExpr(binop.Operator, lhs, rhs)
	}

	return MK_NULL()
}

func evaluateNumericBinaryExpr(operator string, lhs RuntimeValue, rhs RuntimeValue) RuntimeValue {
	result := 0

	valLhs, okLhs := lhs.(NumberVal)
	valRhs, okRhs := rhs.(NumberVal)

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

	return NumberVal{Value: result, Type: Number}
}

func evalProgram(program ast.Program, env Environment) RuntimeValue {
	var lastEvaluated RuntimeValue = MK_NULL()

	for _, statement := range program.Body {
		lastEvaluated = Evaluate(statement, env)
	}

	return lastEvaluated
}

func evalIdentifier(ident ast.Identifier, env Environment) RuntimeValue {
	val := env.LookupVar(ident.Symbol)
	return val
}

func evalVariableDeclaration(variableDeclaration ast.VariableDeclaration, env Environment) RuntimeValue {
	var value RuntimeValue = MK_NULL()

	if variableDeclaration.Value != nil {
		value = Evaluate(variableDeclaration.Value, env)
	}
	return env.DeclareVar(variableDeclaration.Identifier, value, variableDeclaration.Constant)
}

func evalAssignment(assignment ast.AssigmentExpr, env Environment) RuntimeValue {
	if assignment.Assigne.GetKind() != ast_types.Identifier {
		panic("Invalid assignment target")
	}

	varName := assignment.Assigne.(ast.Identifier).Symbol
	return env.AssignVar(varName, Evaluate(assignment.Value, env))
}

func evalObjectExpr(objectExpr ast.ObjectLiteral, env Environment) RuntimeValue {
	obj := ObjectVal{
		Type:       Object,
		Properties: make(map[string]RuntimeValue),
	}

	for _, property := range objectExpr.Properties {
		key := property.Key
		value := property.Value

		var runtimeValue RuntimeValue

		if value == nil {
			runtimeValue = env.LookupVar(key)
		} else {
			runtimeValue = Evaluate(value, env)
		}

		obj.Properties[key] = runtimeValue
	}

	return obj
}

func evalCallExpr(expr ast.CallExpr, env Environment) RuntimeValue {
	args := make([]RuntimeValue, len(expr.Args))

	for idx, arg := range expr.Args {
		args[idx] = Evaluate(arg, env)
	}

	fn := Evaluate(expr.Caller, env)

	switch fn.GetType() {
	case NativeFn:
		result := fn.(NativeFnVal).Call(args, env)
		return result
	case Function:
		function := fn.(FunctionVal)
		scope := NewEnvironment(function.DeclarationEnv)

		// Create the variables for the function arguments
		for idx, arg := range function.Params {
			//TODO: Check the bounds | verify arity of function
			scope.DeclareVar(arg, args[idx], false)
		}

		var result RuntimeValue = MK_NULL()

		// Evaluate the function body line by line
		for _, statement := range function.Body {
			result = Evaluate(statement, scope)
		}

		return result
	}

	panic("Function not found")
}

func evalFunctionDeclaration(declaration ast.FunctionDeclaration, env Environment) RuntimeValue {

	fn := FunctionVal{
		Type:           Function,
		Name:           declaration.Name,
		Params:         declaration.Params,
		DeclarationEnv: &env,
		Body:           declaration.Body,
	}

	return env.DeclareVar(declaration.Name, fn, true)
}

func Evaluate(astNode ast.Stmt, env Environment) RuntimeValue {
	switch astNode.GetKind() {
	case ast_types.NumericLiteral:
		value := astNode.(ast.NumericLiteral).GetValue().(int)
		return NumberVal{Value: value, Type: Number}
	case ast_types.BinaryExpr:
		return evalBinaryExpr(astNode.(ast.BinaryExpr), env)
	case ast_types.Program:
		return evalProgram(astNode.(ast.Program), env)
	case ast_types.Identifier:
		return evalIdentifier(astNode.(ast.Identifier), env)
	case ast_types.VariableDeclaration:
		return evalVariableDeclaration(astNode.(ast.VariableDeclaration), env)
	case ast_types.ObjectLiteral:
		return evalObjectExpr(astNode.(ast.ObjectLiteral), env)
	case ast_types.CallExpr:
		return evalCallExpr(astNode.(ast.CallExpr), env)
	case ast_types.FunctionDeclaration:
		return evalFunctionDeclaration(astNode.(ast.FunctionDeclaration), env)
	case ast_types.AssigmentExpr:
		return evalAssignment(astNode.(ast.AssigmentExpr), env)
	default:
		panic("This AST node is not supported")
	}
}
