package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// 007、字符串
// go run 007_string.go
// 输出：
// 12
// 104 119
// hello
// wo
// hello
// world
// hello, world
// goodbye, world
// left foot, right foot
// left foot
// 常见 ASCII 控制代码的转义方式
// \a      响铃
// \b      退格
// \f      换页
// \n      换行
// \r      回车
// \t      制表符
// \v      垂直制表符
// \'      单引号（只用在单引号包裹的内容中）
// \"      双引号（只用在双引号包裹的内容中）
// \\      反斜杠
// 世界
// 世界
// 世界
// 世界
// true
// true
// true
// true
// false
// true
// 13
// 9
// 0       h
// 1       e
// 2       l
// 3       l
// 4       o
// 5       ,
// 6
// 7       世
// 10      界
// 0       'h'     104
// 1       'e'     101
// 2       'l'     108
// 3       'l'     108
// 4       'o'     111
// 5       ','     44
// 6       ' '     32
// 7       '世'    19990
// 10      '界'    30028
// 9
// 68 65 6c 6c 6f 2c 20 e4 b8 96 e7 95 8c
// [68 65 6c 6c 6f 2c 20 4e16 754c]
// hello, 世界
// A
// 京
// �
// �
// c
// a.b
// abc
// 啊
// 你.好
// 你好
// 1,234,567,890
// 12,345,678.90
// +123,456,789
// -123,456.7890
// abc
// [97 98 99]
// abc
// [1, 2, 3]
// 123 123
// 1111011
// xx=1111011
// 123 123
func main() {
	// 字符串是一个不可改变的字节序列
	// 字符串可以包含任意数据，包括 byte 值 0
	// 字符串通常被解释为采用 utf-8 编码的 unicode 码点（rune）序列

	// 内置的 len 函数可以返回一个字符串中的字节数目（非 rune 字符数目）
	// 索引操作 s[i] 返回第 i 个字节的字节值，i 必须满足 0 < i < len(s) 条件约束
	s := "hello, world"
	fmt.Println(len(s))     // 12
	fmt.Println(s[0], s[7]) // 104(h) 119(w)

	// c := s[len(s)] // 试图访问超出字符串索引范围的字节将会导致异常
	// fmt.Println(c) // panic: runtime error: index out of range

	fmt.Println(s[0:5]) // hello，会截取从 h 到 o 的值，不包括 s[5] 本身，即不包括 o 后面的逗号
	fmt.Println(s[7:9]) // wo，不包括 s[9] 本身，即不包括 r

	// s[i:j] 不管是 i 或 j 都能被忽略，当被忽略时，会采用 0 作为起始位置，采用 len(s) 作为结束位置
	fmt.Println(s[:5]) // hello
	fmt.Println(s[7:]) // world
	fmt.Println(s[:])  // hello, world

	// + 操作符可以拼接新的字符串
	fmt.Println("goodbye" + s[5:]) // goodbye, world

	// 字符串可以进行比较，会逐个比较字节，比较的结果是字符串自然编码的顺序

	// 不会导致 t 被改变，t 仍然保留 a 的原始值 left foot，a 变为 left foot, right foot
	a := "left foot"
	t := a
	a += ", right foot"
	fmt.Println(a) // left foot, right foot
	fmt.Println(t) // left foot

	// 字符串包含的字节序列永远不会被改变
	// s[0] = "L"，会报错 cannot assign to s[0]

	// go 语言源文件总是用 utf-8 编码，因此我们可以将 unicode 码点写到字符串面值中
	arr := []string{
		"常见 ASCII 控制代码的转义方式",
		"\\a\t响铃",
		"\\b\t退格",
		"\\f\t换页",
		"\\n\t换行",
		"\\r\t回车",
		"\\t\t制表符",
		"\\v\t垂直制表符",
		"\\'\t单引号（只用在单引号包裹的内容中）",
		"\\\"\t双引号（只用在双引号包裹的内容中）",
		"\\\\\t反斜杠",
	}
	for _, arg := range arr {
		fmt.Println(arg)
	}

	fmt.Println("世界")                           // 世界
	fmt.Println("\xe4\xb8\x96\xe7\x95\x8c")     // 世界，不是合法的 rune 字符，却对应了有效的 utf-8 码点
	fmt.Println("\u4e16\u754c")                 // 世界，对应 16bit 的码点值
	fmt.Println("\U00004e16\U0000754c")         // 世界，对应 32bit 的码点值
	fmt.Println("世" == "\xe4\xb8\x96")          // true
	fmt.Println("世" == "\u4e16")                // true
	fmt.Println("\xe4\xb8\x96" == "\U00004e16") // true

	fmt.Println(HasPrefix(a, t)) // true
	fmt.Println(HasSuffix(a, t)) // false
	fmt.Println(Contains(a, t))  // true

	s = "hello, 世界"
	fmt.Println(len(s))                    // 13（字节）
	fmt.Println(utf8.RuneCountInString(s)) // 9（字符）
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	// range 会隐式解码 utf8 字符串
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	n := 0
	// for _, _ = range s {
	for range s {
		n++
	}
	fmt.Println(n) // 9

	fmt.Printf("% x\n", s) // 68 65 6c 6c 6f 2c 20 e4 b8 96 e7 95 8c
	r := []rune(s)
	fmt.Printf("%x\n", r) // [68 65 6c 6c 6f 2c 20 4e16 754c]

	// 若将 []rune 类型的 unicode 字符 slice 或数组转为 string，则会对它们进行 utf8 编码
	fmt.Println(string(r)) // hello, 世界
	// 将整数转为字符串会生成对应 unicode 码点的 utf8 字符
	fmt.Println(string(65))     // A
	fmt.Println(string(0x4eac)) // 京
	// 如果对应码点的字符是无效的，则用 \uFFFD 替换
	fmt.Println(string(0xfffd))  // �
	fmt.Println(string(1234567)) // �

	// 字符串和 byte 切片
	fmt.Println(basename("a/b/c.go"))  // c
	fmt.Println(basename("a.b.go"))    // a.b
	fmt.Println(basename("abc"))       // abc
	fmt.Println(basename2("你/好/啊.go")) // 啊
	fmt.Println(basename2("你.好.go"))   // 你.好
	fmt.Println(basename2("你好"))       // 你好

	// 给数字加逗号
	fmt.Println(comma("1234567890"))   // 1,234,567,890
	fmt.Println(comma("12345678.90"))  // 12,345,678.90
	fmt.Println(comma("+123456789"))   // +123,456,789
	fmt.Println(comma("-123456.7890")) // -123,456.7890

	// 字符串和字节切片之间可以相互转换
	x := "abc"
	y := []byte(x)
	x2 := string(y)
	fmt.Println(x)  // abc
	fmt.Println(y)  // [97 98 99]
	fmt.Println(x2) // abc

	// int 数组转 string
	fmt.Println(intsToString([]int{1, 2, 3})) // [1, 2, 3]

	// 字符串和数字的转换
	xx := 123
	yy := fmt.Sprintf("%d", xx)
	fmt.Println(yy, strconv.Itoa(xx)) // 123 123
	// FormatInt 和 FormatUint 可以用不同的进制来格式化数字
	fmt.Println(strconv.FormatInt(int64(xx), 2)) // 1111011
	fmt.Printf("xx=%b\n", xx)                    // xx=1111011

	xxx, _ := strconv.Atoi("123")
	yyy, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(xxx, yyy) // 123 123
}

// HasPrefix 测试一个字符串是否是另一个字符串的前缀
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// HasSuffix 测试一个字符串是否是另一个字符串的后缀
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// Contains 测试一个字符串是否包含在另一个字符串内
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

// 获取文件名
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// 获取文件名 2
func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot > 0 {
		s = s[:dot]
	}
	return s
}

// 给数字加逗号
func comma(s string) string {
	prefix := ""
	suffix := ""
	// 处理小数点
	if dot := strings.LastIndex(s, "."); dot > 0 {
		suffix = s[dot:]
		s = s[:dot]
	}
	// 处理正负号
	if HasPrefix(s, "+") || HasPrefix(s, "-") {
		prefix = s[:1]
		s = s[1:]
	}
	// 加逗号
	n := len(s)
	if n <= 3 {
		return s
	}
	return prefix + comma(s[:n-3]) + "," + s[n-3:] + suffix
}

// int 数组转 string
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
