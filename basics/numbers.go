package main

import (
	"fmt"
	"math/cmplx"
	"sort"
)

func main() {
	count, total := 0, 0
	for n := 1; n <= 100; n++ {
		if n%3 == 0 || n%5 == 0 {
			count++
			total += n
		}
	}

	fmt.Println(float32(total) / float32(count))

	nums := []float64{2, 1, 3}

	sort.Float64s(nums)
	var median float64
	i := len(nums) / 2
	if len(nums)%2 == 1 {
		median = nums[i]
	} else {
		median = (nums[i-1] + nums[i]) / 2
	}
	fmt.Println(median)

	var x complex128
	x = 3 + 4i
	fmt.Println(cmplx.Abs(x))

}
