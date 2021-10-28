package goroutines

import (
	"testing"
	"time"
)

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name: "James",
	}

	user2 := UserBalance{
		Name: "Mike",
	}

	go Transfer(&user1, &user2, 1000)
	go Transfer(&user2, &user1, 1000)

	time.Sleep(5 * time.Second)
}
