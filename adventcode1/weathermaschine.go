package main

import (
	"fmt"
)

// a word 1abc2 ; find first and last digit (12) ; sum them all up.

func sumFirstAndLastDigit(word string) int {
	var first int
	var last int
	for i := 0; i < len(word); i++ {
		if word[i] >= '0' && word[i] <= '9' {
			if first == 0 {
				first = int(word[i] - '0')
			}
			last = int(word[i] - '0')
		}
	}
	return first*10 + last
}

func main() {
	fmt.Println(sumFirstAndLastDigit("1ab6c2"))
}
