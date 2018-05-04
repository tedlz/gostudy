package main

import "fmt"

// 001、Go 语言的关键字
// go run 001_keywords.go
// 输出：
// 关键字：
// break       case        chan        const       continue
// default     defer       else        fallthrough for
// func        go          goto        if          import
// interface   map         package     range       return
// select      struct      switch      type        var

// 常量：
// true        false       iota        nil

// 类型：
// int         int8        int16       int32       int64
// uint        uint8       uint16      uint32      uint64      uintptr
// float32     float64     complex64   complex128
// bool        byte        rune        string      error

// 函数：
// make        len         cap         new         append
// copy        close       delete      complex     real
// imag        panic       recover
func main() {
	// 关键字
	var keywords = []string{
		"break",
		"case",
		"chan",
		"const",
		"continue",
		"default",
		"defer",
		"else",
		"fallthrough",
		"for",
		"func",
		"go",
		"goto",
		"if",
		"import",
		"interface",
		"map",
		"package",
		"range",
		"return",
		"select",
		"struct",
		"switch",
		"type",
		"var",
	}
	fmt.Println("关键字：")
	for k, v := range keywords {
		fmt.Print(v)
		if (k+1)%5 == 0 {
			fmt.Println()
		} else {
			split := "\t"
			if len(v) <= 7 {
				split += "\t"
			}
			fmt.Print(split)
		}
	}
	fmt.Println()

	// 常量
	var constant = []string{
		"true",
		"false",
		"iota",
		"nil",
	}
	fmt.Println("常量：")
	for k, v := range constant {
		fmt.Print(v)
		if k < len(constant)-1 {
			fmt.Print("\t\t")
		}
	}
	fmt.Print("\n\n")

	// 类型
	var t = []string{
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"uintptr",
		"float32",
		"float64",
		"complex64",
		"complex128",
		"bool",
		"byte",
		"rune",
		"string",
		"error",
	}
	fmt.Println("类型：")
	for k, v := range t {
		fmt.Print(v)
		if k == 4 || k == 10 || k == 14 {
			fmt.Println()
		} else {
			split := "\t"
			if len(v) <= 7 {
				split += split
			}
			fmt.Print(split)
		}
	}
	fmt.Print("\n\n")

	// 函数
	var function = []string{
		"make",
		"len",
		"cap",
		"new",
		"append",
		"copy",
		"close",
		"delete",
		"complex",
		"real",
		"imag",
		"panic",
		"recover",
	}
	fmt.Println("函数：")
	for k, v := range function {
		fmt.Print(v)
		if (k+1)%5 == 0 {
			fmt.Println()
		} else if k < len(function)-1 {
			fmt.Print("\t\t")
		}
	}
	fmt.Println()
}
