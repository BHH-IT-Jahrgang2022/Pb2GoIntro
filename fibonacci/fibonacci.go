// function that creates the n number of fibonacci sequence
package main

import "fmt"

// function that creates the n number of fibonacci sequence
func fibonacci(n int) []int {
	sequence := []int{0, 1}

	for i := 2; i < n; i++ {
		sequence = append(sequence, sequence[i-1]+sequence[i-2])
	}

	return sequence
}

func main() {
	n := 10 // change this to the desired number of fibonacci sequence
	result := fibonacci(n)
	fmt.Println(result)
}
