package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(100)
				fmt.Println("Amount of balance now:", account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance:", account.GetBalance())
}
