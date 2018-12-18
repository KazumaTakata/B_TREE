package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	p := []int{2, 5, 5, 6, 9, 10, 13, 15, 16, 19, 20, 21, 22, 25, 26, 27, 28, 30, 31, 34, 35, 37, 40, 41, 42, 43, 44, 47, 50}

	start := time.Now()

	max_price := cut_rod(p, 22)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
	fmt.Printf("%d", max_price)
}

func cut_rod(p []int, n int) int {
	if n == 0 {
		return 0
	}

	q := -10000

	for i := 0; i < n; i++ {
		q = Max(q, p[i]+cut_rod(p, n-i-1))
	}
	return q
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
