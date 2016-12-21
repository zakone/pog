package bank_test

import (
	"fmt"
	"testing"
	"./bank"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Withdraw(200), true; got != want {
		t.Errorf("Withdraw Error: %t, want %t", got, want)
	}

	if got, want := bank.Withdraw(200), false; got != want {
		t.Errorf("Withdraw Error: %t, want %t", got, want)
	}
}
