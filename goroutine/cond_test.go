package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	locker = sync.Mutex{}
	cond   = sync.NewCond(&locker)
	group  = sync.WaitGroup{}
)

func WaitCondition(value int) {
	group.Add(1)
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
	group.Done()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	// value di eksekusi satu-satu
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// value di eksekusi secara bersamaan
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()

	group.Wait()
}
