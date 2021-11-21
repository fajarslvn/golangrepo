package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "James",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Mike",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(3 * time.Second)

	fmt.Println("User:", user1.Name, ", Balance:", user1.Balance)
	fmt.Println("User:", user2.Name, ", Balance:", user2.Balance)
}
