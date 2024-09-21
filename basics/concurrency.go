package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
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

func CheckURL(url string, t time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

type response struct {
	fileName string
	size     int64
	error    string
}

func worker_file(fileName string, ch chan<- response) {
	resp := response{fileName: fileName, size: -1, error: ""}
	st, err := os.Stat(fileName)
	if err == nil {
		resp.size = st.Size()
	} else {
		resp.error = err.Error()
	}
	ch <- resp
}

func main() {
	for i := 0; i < 10; i++ {
		go worker(i)
	}
	fmt.Println("main")
	time.Sleep(2 * time.Second)

	// for i := 0; i < 3; i++ {
	// 	go worker_hard(i)
	// }
	// time.Sleep(5 * time.Second)

	fmt.Println("-------------------------")

	ch := make(chan int)
	// ch <- 99 // send  fatal error: all goroutines are asleep - deadlock!
	go func() {
		ch <- 99 // send
	}()
	val := <-ch // receive
	fmt.Printf("got %d\n", val)

	// go func() {
	// 	val := <-ch // receive
	// 	fmt.Printf("got %d\n", val)
	// }()
	// time.Sleep(1 * time.Second)

	fmt.Println("-------------------------")

	var wg sync.WaitGroup
	ch2 := make(chan string)

	for i := 0; i < 3; i++ {
		go func(id int) {
			for msg := range ch2 {
				fmt.Printf("%d started %s\n", id, msg)
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("%d finished %s\n", id, msg)
				wg.Done()
			}
		}(i)
	}
	time.Sleep(500 * time.Millisecond)

	for _, msg := range []string{"A", "B", "C", "D", "E", "F"} {
		wg.Add(1)
		ch2 <- msg
	}
	wg.Wait()
	fmt.Println("all jobs done")

	fmt.Println("-------------------------")

	url := "https://www.linkedin.com/learning/"
	fmt.Println(CheckURL(url, 5*time.Second))
	fmt.Println(CheckURL(url, 5*time.Millisecond))

	fmt.Println("-------------------------")

	files := []string{
		"text.go",
		"structs.go",
		"numbers.go",
	}
	ch3 := make(chan response)

	for _, f := range files {
		go worker_file(f, ch3)
	}

	for range files {
		r := <-ch3
		fmt.Printf("%s -> %d %s\n", r.fileName, r.size, r.error)
	}
}
