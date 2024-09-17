package main

import (
	"fmt"
	"time"
)

func worker(n int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("goroutine %d\n", n)
}

func worker_hard(n int) {
	a := 0
	for i := 0; i < 1000000000; i++ {
		a = i / 2
	}
	fmt.Printf("goroutine %d\n", a)
}

func main() {
	for i := 0; i < 10; i++ {
		go worker(i)
	}
	fmt.Println("main")
	// time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		go worker_hard(i)
	}
	time.Sleep(5 * time.Second)

}
