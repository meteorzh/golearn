package eval

import (
	"fmt"
	"math"
)

// Expr: 算术表达式
type Expr interface {
	// Eval 返回表达式在 env 上下文下的值
	Eval(env Env) float64
}

// Var表示一个变量, 比如x
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

// literal是一个数字常量, 比如3.141
type literal float64

func (l literal) Eval(env Env) float64 {
	return float64(l)
}

// unary 表示一元操作符表达式
type unary struct {
	op rune // '+', '-' 中的一个
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

// binary 表示二元操作符表达式, 比如x+y
type binary struct {
	op   rune // '+', '-', '*', '/'中的一个
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

// call 表示函数调用表达式
type call struct {
	fn   string // one of "pow", "sin", "sqrt"中的一个
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

type Env map[Var]float64
