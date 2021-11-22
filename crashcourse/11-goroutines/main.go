package main

import (
	"fmt"
	"time"
)

func doSomething(thingToDo string) {
	for i := 0; i <= 10; i++ {
		fmt.Println("Step:", i, "I am:", thingToDo)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go doSomething("eating lunch")
	go doSomething("watching movie")
	doSomething("drinking coffee")
} 