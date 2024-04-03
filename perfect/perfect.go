package main

import "fmt"

func main() {
	fmt.Printf("Looking for perfect numbers\n")

	for i := 1; i < 10000; i++ {
		if perfect(i) {
			fmt.Printf("%d ", i)
		}
	}

	fmt.Println("\nDone")
}

func perfect(n int) bool {
	sum := 0

	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	return sum == n
}
