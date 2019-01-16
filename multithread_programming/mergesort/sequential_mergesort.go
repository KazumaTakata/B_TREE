package main

import (
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
)

func main() {

	s := make([]int, 1000000)

	for i := 0; i < 1000000; i++ {
		s[i] = rand.Intn(100000)
	}

	// s := []int{1, 2, 54, 4, 3, 34, 653, 433, 654, 44, 5437}

	start := time.Now()

	mergesortParallel(s, 0, len(s)-1)

	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
	// fmt.Printf("%v\n", s)
}

func mergesortParallel(A []int, p int, r int) {
	if p < r {
		q := (p + r) / 2

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			wg.Done()
			mergesort(A, p, q)
		}()
		go func() {
			wg.Done()
			mergesort(A, q+1, r)
		}()

		wg.Wait()
		merge(A, p, q, r)
	}
}

func mergesort(A []int, p int, r int) {
	if p < r {
		q := (p + r) / 2

		mergesort(A, p, q)

		mergesort(A, q+1, r)

		merge(A, p, q, r)
	}
}

func merge(A []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	L := make([]int, n1+1)
	R := make([]int, n2+1)

	for i := 0; i < n1; i++ {
		L[i] = A[p+i]
	}

	for i := 0; i < n2; i++ {
		R[i] = A[q+i+1]
	}

	L[n1] = math.MaxInt64
	R[n2] = math.MaxInt64

	i := 0
	j := 0

	for k := p; k <= r; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i = i + 1
		} else {
			A[k] = R[j]
			j = j + 1
		}
	}

}
