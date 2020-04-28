package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"text/scanner"
)

// Unmarshal 解析 S 表达式并填充地址在非空指针之外的变量
func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next() // 得到第一个 token
	defer func() {
		// 注意：不是良好的错误处理示例
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

// !+ lexer
type lexer struct {
	scan  scanner.Scanner
	token rune // 当前标记
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { // 注意：不是良好的错误处理示例
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

// !- lexer

// !+ read
// read 函数是一个解码器，用于解码一小部分形式良好的 S 表达式的子集，为了简化我们的例子，我们采用了许多可疑的捷径
// 解析器假设：
// - S 表达式格式输入正确，它不做错误检查
// - S 表达式输入对应于变量的类型
// - 输入中所有的数字均为非负的十进制整数
// - ((key value) ...) 结构中的所有 key 都是未加引号的符号
// - 输入中不包含虚线列表，例如 (1 2 . 3)
// - 输入中不包含 Lisp 阅读器宏，例如 'x 和 #'x
// 反射逻辑假设：
// - v 始终是 S 表达式值的适当类型的变量，例如，v 不能是 boolean、interface、channel 或 function，
//   并且如果 v 是 array，则输入的元素数量必须正确
// - 顶层调用 read 中的 v 具有其类型的零值，不需要清除
// - 如果 v 是数字变量，那么它是一个有符号的整数
func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		// 唯一有效的标识符是 "nil" 和 struct 字段名
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) // 注意：忽略错误
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text()) // 注意：忽略错误
		v.SetInt(int64(i))
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next() // comsume ')'
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

// !- read

// !+ readList
func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}
	case reflect.Slice: // (item ...)
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}
	case reflect.Struct: // ((name value) ...)
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}
	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}
	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

// !- readList

// !+ endList
func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}

// !- endList
