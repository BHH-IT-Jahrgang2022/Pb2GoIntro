package main

import "fmt"

func calc_fac(n int) int {
	if n == 0 {
		return 1
	}
	return n * calc_fac(n-1)
}

func main() {
	fmt.Println("Hello, Go!")
}
