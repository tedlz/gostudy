package main

import (
	"fmt"
	"gostudy/07、接口/files"
	"math"
	"testing"
)

// 013、示例：表达式求值
func main() {

}

// TestEval 是对求值器的一个测试，它使用了我们会在第 11 章讲解的 testing 包，
// 但是现在知道调用 t.Errorf 会报告一个错误就足够了。这个函数循环遍历一个表格中的输入
// 这个表格中定义了三个表达式和针对每个表达式不同的环境变量
// 第一个表达式根据给定圆的面积 A 计算它的半径
// 第二个表达式通过两个变量 x 和 y 计算两个立方体的体积之和
// 第三个表达式将华氏度转为摄氏度
// 对于表格中的每一条记录，这个测试会解析它的表达式然后在环境变量中计算它并输出结果
// go test -v 013_eval_test.go
// 输出：
// === RUN   TestEval
//
// sqrt(A / pi)
//         map[A:87616 pi:3.141592653589793] => 167
//
// pow(x, 3) + pow(y, 3)
//         map[x:12 y:1] => 1729
//         map[x:9 y:10] => 1729
//
// 5 / 9 * (F - 32)
//         map[F:-40] => -40
//         map[F:32] => 0
//         map[F:212] => 100
// --- PASS: TestEval (0.00s)
func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  files.Env
		want string
	}{
		{"sqrt(A / pi)", files.Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", files.Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", files.Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", files.Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", files.Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", files.Env{"F": 212}, "100"},
	}
	var prevExpr string
	for _, test := range tests {
		// 仅在更改时打印 expr
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := files.Parse(test.expr)
		if err != nil {
			t.Error(err) // 解析错误
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

// TestErrors *
// go test -v 013_eval_test.go
// 输出：
// === RUN   TestErrors
// x % 2               unexpected '%'
// math.Pi             unexpected '.'
// !true               unexpected '!'
// "hello"             unexpected '"'
// log(10)             unknown function "log"
// sqrt(1, 2)          call to sqrt has 2 args, want 1
// --- PASS: TestErrors (0.00s)
func TestErrors(t *testing.T) {
	for _, test := range []struct{ expr, wantErr string }{
		{"x % 2", "unexpected '%'"},
		{"math.Pi", "unexpected '.'"},
		{"!true", "unexpected '!'"},
		{`"hello"`, "unexpected '\"'"},
		{"log(10)", `unknown function "log"`},
		{"sqrt(1, 2)", "call to sqrt has 2 args, want 1"},
	} {
		expr, err := files.Parse(test.expr)
		if err == nil {
			vars := make(map[files.Var]bool)
			err = expr.Check(vars)
			if err == nil {
				t.Errorf("unexpected success: %s", test.expr)
				continue
			}
		}
		fmt.Printf("%-20s%v\n", test.expr, err)
		if err.Error() != test.wantErr {
			t.Errorf("got error %s, want %s", err, test.wantErr)
		}
	}
}
