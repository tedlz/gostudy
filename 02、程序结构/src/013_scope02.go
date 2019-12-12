package main

import (
	"log"
	"os"
)

// 013、作用域 02
// go run 013_scope02.go
// 输出：
// 2019/12/12 14:26:00 Working directory = /data/go/src/gostudy/02、程序结构/src
// exit status 1
func main() {
	init1()
	init2()
}

var cwd string

func init1() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Fatalf("Working directory = %s", cwd) // 不写这句会报错：cwd declared and not used
}

func init2() {
	var err error         // 单独声明 err
	cwd, err = os.Getwd() // 且不用简短声明（:=）就没问题
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}
