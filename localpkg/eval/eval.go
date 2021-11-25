package eval

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

// wraps control text line wrap
var wraps int

type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64
	Check(vars map[Var]bool) error
	String() string
}

type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (v Var) String() string {
	return fmt.Sprintf("Var: %s\n", string(v))
}

type Env map[Var]float64

type literal float64

func (l literal) Eval(Env) float64 {
	return float64(l)
}

func (l literal) Check(map[Var]bool) error {
	return nil
}

func (l literal) String() string {
	return fmt.Sprintf("Literal: %f\n", l)
}

type unary struct {
	op rune // one of '+', '-'
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

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

func (u unary) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%*s\n", wraps, "unary"))
	wraps += len("unary") / 2
	defer func() {
		wraps -= len("unary") / 2
	}()
	// 2 line branches
	drawBranch(&buf)
	buf.WriteString(fmt.Sprintf("Op: %c\n", u.op))
	drawBranch(&buf)
	drawSubExpr(&buf, u.x)
	return buf.String()
}

type binary struct {
	op   rune // one of '+', '-', '*', '/'
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

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (b binary) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%*s\n", wraps, "binary"))
	wraps += len("binary") / 2
	defer func() {
		wraps -= len("binary") / 2
	}()
	// 2 line branches
	drawBranch(&buf)
	drawSubExpr(&buf, b.x)

	drawBranch(&buf)
	buf.WriteString(fmt.Sprintf("Op: %c\n", b.op))

	drawBranch(&buf)
	drawSubExpr(&buf, b.y)
	return buf.String()
}

type call struct {
	fn   string // one of "pow", "sqrt", "sin"
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
	panic(fmt.Sprintf("unsupported function call: %q", c.fn))
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d", c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (c call) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%*s\n", wraps, "call"))
	wraps += len("call") / 2
	defer func() {
		wraps -= len("call") / 2
	}()
	// 2 line branches
	drawBranch(&buf)
	buf.WriteString(fmt.Sprintf("Func: %v\n", c.fn))
	for _, arg := range c.args {
		drawBranch(&buf)
		drawSubExpr(&buf, arg)
	}
	return buf.String()
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

func drawBranch(buf *bytes.Buffer) {
	for i := 0; i < 2; i++ {
		buf.WriteString(fmt.Sprintf("%*s\n", wraps, "|"))
	}
	buf.WriteString(fmt.Sprintf("%*s", wraps+1, "--"))
}

func drawSubExpr(buf *bytes.Buffer, s fmt.Stringer) {
	wraps += 2
	defer func() {
		wraps -= 2
	}()
	buf.WriteString(fmt.Sprintf("%v", s))
}
