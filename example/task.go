package main

import (
	"fmt"
)

func prime(n uint64) bool {
	switch {
	case n <= 1:
		return false
	case n == 2:
		return true
	case n%2 == 0:
		return false
	default:
		for i := uint64(3); i*i <= n; i += 2 {
			if n%uint64(i) == 0 {
				return false
			}
		}
		return true
	}
}

func fibonacci(n int) []uint64 {
	sequence := []uint64{0, 1}

	for i := 2; i < n; i++ {
		sequence = append(sequence, sequence[i-1]+sequence[i-2])
	}

	return sequence
}

func main() {
	fmt.Printf("Looking for prime fibonacci numbers\n")

	fibonacci := fibonacci(100)

	for _, v := range fibonacci {
		if prime(v) {
			fmt.Printf("%d ", v)
		}
	}

	fmt.Println("\nDone")
}
