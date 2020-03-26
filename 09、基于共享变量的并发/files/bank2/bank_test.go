package bank_test

import (
	"sync"
	"testing"

	bank "gostudy/09、基于共享变量的并发/files/bank2"
)

func TestBank(t *testing.T) {
	// 1000 个并发存款
	var n sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			bank.Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := bank.Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
