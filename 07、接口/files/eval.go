package files

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"text/scanner"
)

// Expr 是一个算术表达式
type Expr interface {
	// Eval 在 env 环境中返回 Expr 的值
	Eval(env Env) float64
	// 检查报告 Expr 中的错误并将其添加到 vars 集合中
	Check(vars map[Var]bool) error
}

// Var 表示变量，例如 x
type Var string

// literal 是数字常数，例如 3.141
type literal float64

// unary 表示一元运算符表达式，例如 -x
type unary struct {
	op rune
	x  Expr
}

// binary 表示二进制运算符表达式，例如 x+y
type binary struct {
	op   rune
	x, y Expr
}

// call 代表函数调用表达式，例如 sin(x)
type call struct {
	fn   string
	args []Expr
}

// Env 为了计算一个包含变量的表达式，我们需要一个环境变量将变量的名字映射成对应的值
type Env map[Var]float64

// 我们也需要每个表达式去定义一个 Eval 方法，这个方法会根据给定的环境变量返回表达式的值
// 因为每个表达式都必须提供这个方法，我们将它加到 Expr 接口中
// 这个包只会对外公开 Expr、Env 和 Var 类型，调用方不需要获取其它的表达式类型就可以使用这个求值器

// Eval - Var 类型的 Eval 方法对环境变量进行查找，如果变量未在环境中定义过，方法会返回一个零值
func (v Var) Eval(env Env) float64 {
	return env[v]
}

// literal 类型的 Eval 方法简单返回它真实的值
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// unary 和 binary 的 Eval 方法会递归计算它的运算对象，然后将运算符 op 作用到它们身上
// 我们不认为除以零或无穷大是错误的，因为它们会产生结果，即使不是有限的
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
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

// call 类型的 Eval 方法会计算 pow、sin 或者 sqrt 函数的参数值，然后调用对应在 math 包中的函数
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

// 一些方法会失败。例如，一个 call 表达式可能有未知的函数或者错误数量的参数
// 用一个无效的运算符如 ! 或 < 去构建一个 unary 或者 binary 表达式也是可能会发生的（尽管下面提到的 Parse 函数不会这么做）
// 这些错误会让 Eval 方法 panic。其它的错误，像计算一个没有在环境变量中出现过的 Var，只会让 Eval 方法返回一个错误的结果
// 所有的这些错误都可以通过在计算前检查 Expr 来发现，这是我们接下来要讲的 Check 方法的工作，但是让我们先测试 Eval 方法

// 幸运的是目前为止所有的输出都是适合的格式，但是我们的运气不可能一直都有
// 甚至在解释型语言中，为了静态错误检查语法是非常常见的；静态错误就是不用运行程序就可以检测出来的错误
// 通过将静态检查和动态的部分分开，我们可以快速的检查错误并且对于多次检查只执行一次，而不是每次表达式计算的时候都进行检查
// 让我们往 Expr 接口中增加另一个方法，Check 方法在一个表达式语义树检查出静态错误。我们马上会说明它的 vars 参数

// Check - Var 和 literal 的 Check 计算不可能的失败，所以这些类型的 Check 方法会返回一个 nil 值
func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}
func (literal) Check(vars map[Var]bool) error {
	return nil
}

// unary 和 binary 的 Check 方法会首先检查操作符是否有效，然后递归的检查运算单元
func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
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

// call 的 Check 方法首先检查调用的函数是否已知并且有没有正确个数的参数，然后递归检查每一个参数
func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

// Check 方法的参数是一个 Var 类型的集合，这个集合聚集从表达式中找到的变量名
// 为了保证成功的计算，这些变量中的每一个都必须出现在环境变量中。从逻辑上讲，这个集合就是调用 Check 方法返回的结果，
// 但是因为这个方法是递归调用的，所以对于 Check 方法填充结果到一个作为参数传入的集合中会更加方便
// 调用方在初始调用时必须提供一个空的集合

// ==========================================================================
// 以下来自：https://github.com/adonovan/gopl.io/blob/master/ch7/eval/parse.go
// ==========================================================================

// ---- lexer ----

// This lexer is similar to the one described in Chapter 13.
type lexer struct {
	scan  scanner.Scanner
	token rune // current lookahead token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

type lexPanic string

// describe returns a string describing the current token, for use in errors.
func (lex *lexer) describe() string {
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", lex.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token)) // any other rune
}

func precedence(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

// ---- parser ----

// Parse parses the input string as an arithmetic expression.
//
//   expr = num                         a literal number, e.g., 3.14159
//        | id                          a variable name, e.g., x
//        | id '(' expr ',' ... ')'     a function call
//        | '-' expr                    a unary operator (+-)
//        | expr '+' expr               a binary operator (+-*/)
//
func Parse(input string) (_ Expr, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
			// no panic
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			// unexpected panic: resume state of panic.
			panic(x)
		}
	}()
	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	lex.next() // initial lookahead
	e := parseExpr(lex)
	if lex.token != scanner.EOF {
		return nil, fmt.Errorf("unexpected %s", lex.describe())
	}
	return e, nil
}

func parseExpr(lex *lexer) Expr { return parseBinary(lex, 1) }

// binary = unary ('+' binary)*
// parseBinary stops when it encounters an
// operator of lower precedence than prec1.
func parseBinary(lex *lexer, prec1 int) Expr {
	lhs := parseUnary(lex)
	for prec := precedence(lex.token); prec >= prec1; prec-- {
		for precedence(lex.token) == prec {
			op := lex.token
			lex.next() // consume operator
			rhs := parseBinary(lex, prec+1)
			lhs = binary{op, lhs, rhs}
		}
	}
	return lhs
}

// unary = '+' expr | primary
func parseUnary(lex *lexer) Expr {
	if lex.token == '+' || lex.token == '-' {
		op := lex.token
		lex.next() // consume '+' or '-'
		return unary{op, parseUnary(lex)}
	}
	return parsePrimary(lex)
}

// primary = id
//         | id '(' expr ',' ... ',' expr ')'
//         | num
//         | '(' expr ')'
func parsePrimary(lex *lexer) Expr {
	switch lex.token {
	case scanner.Ident:
		id := lex.text()
		lex.next() // consume Ident
		if lex.token != '(' {
			return Var(id)
		}
		lex.next() // consume '('
		var args []Expr
		if lex.token != ')' {
			for {
				args = append(args, parseExpr(lex))
				if lex.token != ',' {
					break
				}
				lex.next() // consume ','
			}
			if lex.token != ')' {
				msg := fmt.Sprintf("got %q, want ')'", lex.token)
				panic(lexPanic(msg))
			}
		}
		lex.next() // consume ')'
		return call{id, args}

	case scanner.Int, scanner.Float:
		f, err := strconv.ParseFloat(lex.text(), 64)
		if err != nil {
			panic(lexPanic(err.Error()))
		}
		lex.next() // consume number
		return literal(f)

	case '(':
		lex.next() // consume '('
		e := parseExpr(lex)
		if lex.token != ')' {
			msg := fmt.Sprintf("got %s, want ')'", lex.describe())
			panic(lexPanic(msg))
		}
		lex.next() // consume ')'
		return e
	}
	msg := fmt.Sprintf("unexpected %s", lex.describe())
	panic(lexPanic(msg))
}
