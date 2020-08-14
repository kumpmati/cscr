package runtime

import (
	"github.com/kumpmati/cscr/pkg/ast"
	"log"
	"strconv"
)

type EvalFunc func(n ast.Node, r *R) Value

func DefaultEvalFunc(node ast.Node, r *R) Value {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}()
	switch node.Type.Type {
	case ast.Keyword:
		// todo: keyword support
	case ast.Block:
		// accumulate return value over every child
		ret := Constant{}
		for _, c := range node.Body {
			ret.value = DefaultEvalFunc(c, r).Get()
		}
		return ret
	case ast.Constant:
		// convert string representation to actual float32
		f, err := strconv.ParseFloat(node.Value, 64)
		if err != nil {
			panic(err)
		}
		return Constant{value: float32(f)}
	case ast.Reference:
		// return a reference to the variable
		if variable := r.memory.GetVar(node.Value); variable != nil {
			return Reference{variable}
		}
		panic(nilReferenceErr)
	case ast.Operator:
		switch node.Value {
		case "!":
			val := DefaultEvalFunc(node.Body[1], r)
			if val.Get() != 0 {
				return Constant{0}
			}
			return Constant{1}
		case ">", ">=", "<", "<=", "==", "!=":
			return logicEval(node.Value, node, r)
		case "+", "-", "*", "/":
			return mathOpEval(node.Value, node, r)
		case "+=", "-=", "*=", "/=":
			targetName := node.Body[0].Value
			target := r.memory.GetVar(targetName)
			if target == nil {
				panic(variableNotDefinedErr)
			}
			mathOp := node.Value[0 : len(node.Value)-1]
			target.value = mathOpEval(mathOp, node, r)
			return target.value
		case ":=":
			return assignOpEval(node, r)
		case "=":
			return setOpEval(node, r)
		}
	}
	return Constant{0}
}

// math operations
func mathOpEval(mathOp string, opNode ast.Node, r *R) Value {
	lhs, rhs := DefaultEvalFunc(opNode.Body[0], r), DefaultEvalFunc(opNode.Body[1], r)
	switch mathOp {
	case "+":
		return Constant{lhs.Get() + rhs.Get()}
	case "-":
		return Constant{lhs.Get() - rhs.Get()}
	case "*":
		return Constant{lhs.Get() * rhs.Get()}
	case "/":
		return Constant{lhs.Get() / rhs.Get()}
	default:
		return lhs
	}
}

func logicEval(op string, node ast.Node, r *R) Value {
	lhs, rhs := DefaultEvalFunc(node.Body[0], r), DefaultEvalFunc(node.Body[1], r)
	switch {
	case op == ">" && lhs.Get() > rhs.Get(),
		op == ">=" && lhs.Get() >= rhs.Get(),
		op == "<" && lhs.Get() < rhs.Get(),
		op == "<=" && lhs.Get() <= rhs.Get(),
		op == "==" && lhs.Get() == rhs.Get(),
		op == "!=" && lhs.Get() != rhs.Get():
		return Constant{1}
	}
	return Constant{0}
}

// create a new variable based on the left hand side,
// and assign it the evaluated value of the right hand side
func assignOpEval(node ast.Node, r *R) Value {
	varName := node.Body[0].Value
	target := r.memory.NewVariable(varName, DefaultEvalFunc(node.Body[1], r))
	return Reference{target}
}

func setOpEval(node ast.Node, r *R) Value {
	// get the variable from memory
	target := r.memory.GetVar(node.Body[0].Value)
	if target == nil {
		panic(variableNotDefinedErr)
	}
	// evaluate the right hand side
	target.value = DefaultEvalFunc(node.Body[1], r)
	return Reference{target}
}
