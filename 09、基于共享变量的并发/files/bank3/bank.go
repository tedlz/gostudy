// Package bank 为一个并发安全的银行提供一个账户
package bank

import "sync"

var (
	mu      sync.Mutex // 守护 balance
	balance int
)

// Deposit 存款
func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

// Balance 查询余额
func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
