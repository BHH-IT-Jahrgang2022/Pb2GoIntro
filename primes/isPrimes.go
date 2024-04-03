package main

import "fmt"

func main() {
	isPrime := prime(97)
	fmt.Println(isPrime)
}

func prime(n int) bool {
	switch {
	case n <= 1:
		return false
	case n == 2:
		return true
	case n%2 == 0:
		return false
	default:
		for i := 3; i*i <= n; i += 2 {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
}
