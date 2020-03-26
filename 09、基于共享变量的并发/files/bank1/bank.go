// Package bank 为一个并发安全的银行提供一个账户
package bank

var deposits = make(chan int) // 汇款到存款
var balances = make(chan int) // 收到余额

// Deposit 存款
func Deposit(amount int) { deposits <- amount }

// Balance 查询余额
func Balance() int { return <-balances }

func teller() {
	var balance int // balance 只限于 teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // 启动 monitor goroutine
}
