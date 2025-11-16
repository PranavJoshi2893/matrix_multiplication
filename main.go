package main

import (
	"fmt"
	"math/rand"
	"time"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const N = 1024
const RAND_MAX = 100

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	if err := WriteFlatMatrix("matrix_a.txt", N, rng); err != nil {
		fmt.Println("write A error:", err)
		return
	}
	if err := WriteFlatMatrix("matrix_b.txt", N, rng); err != nil {
		fmt.Println("write B error:", err)
		return
	}

	a, err := LoadFlatMatrix("matrix_a.txt", N)
	if err != nil {
		fmt.Println("load A failed:", err)
		return
	}

	b, err := LoadFlatMatrix("matrix_b.txt", N)
	if err != nil {
		fmt.Println("load B failed:", err)
		return
	}

	start := time.Now()
	c := MatmulOptimized(a, b)
	elapsed := time.Since(start)
	fmt.Printf("Time: %v\n", elapsed)

	if err := WriteFlatResult("matrix_result.txt", c); err != nil {
		fmt.Println("write result failed:", err)
		return
	}
}
