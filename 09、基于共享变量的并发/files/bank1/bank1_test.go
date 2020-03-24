package bank_test

import (
	"fmt"
	"testing"

	bank "gostudy/09、基于共享变量的并发/files/bank1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// 等待这两笔交易
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
