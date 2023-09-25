package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {

	samples, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Incorrect sample rate")
		os.Exit(1)
	}

	start := time.Now()
	pi := estimate(samples)
	elapsed_time := time.Since(start)

	fmt.Println("Single Core")
	fmt.Println("PI  : ", pi)
	fmt.Println("Time: ", elapsed_time)

	start = time.Now()
	N := runtime.NumCPU()
	pi = spread_mean(samples, estimate, N)
	elapsed_time = time.Since(start)

	fmt.Println(N, "Cores")
	fmt.Println("PI  : ", pi)
	fmt.Println("Time: ", elapsed_time)
}

func spread_mean(samples int, f func(int) float64, P int) (estimated float64) {
	counts := make(chan float64)

	for i := 0; i < P; i++ {
		go func() { counts <- f(samples / P) }()
	}

	for i := 0; i < P; i++ {
		estimated += <-counts
	}
	return estimated / float64(P)
}

func estimate(n int) float64 {
	const radius = 1.0

	var (
		seed   = rand.NewSource(time.Now().UnixNano())
		random = rand.New(seed)
		inside int
	)

	for i := 0; i < n; i++ {
		x, y := random.Float64(), random.Float64()

		if num := math.Sqrt(x*x + y*y); num < radius {
			inside++
		}
	}
	return 4 * float64(inside) / float64(n)
}
