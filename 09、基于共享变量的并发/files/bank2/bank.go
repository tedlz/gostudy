// Package bank 为一个并发安全的银行提供一个账户
package bank

var (
	sema    = make(chan struct{}, 1) // 一个保护 balance 的二元信号量
	balance int
)

// Deposit 存款
func Deposit(amount int) {
	sema <- struct{}{} // 获取 token
	balance = balance + amount
	<-sema // 释放 token
}

// Balance 查询余额
func Balance() int {
	sema <- struct{}{} // 获取 token
	b := balance
	<-sema // 释放 token
	return b
}
