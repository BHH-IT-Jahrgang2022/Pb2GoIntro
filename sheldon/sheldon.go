package main

import (
	"fmt"
	"strconv"
)

func main() {
	primes := []int{}
	for i := 2; len(primes) <= 10000; i++ {
		if prime(i) {
			primes = append(primes, i)
		}
	}
	for i, prime := range primes {
		if sheldon(prime, i+1) {
			fmt.Println(prime)
		}
	}
}

func prime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func reverse(n int) int {
	reversed := 0
	for n > 0 {
		reversed = 10*reversed + n%10
		n /= 10
	}
	return reversed
}

func isPermutation(a, b, r int) bool {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)
	if len(aStr) != len(bStr) {
		return false
	}
	counts := make(map[rune]int)
	for _, r := range aStr {
		counts[r]++
	}
	for _, r := range bStr {
		counts[r]--
		if counts[r] < 0 {
			return false
		}
	}
	return XthPrime(r, b)
}

func XthPrime(n, ix int) bool {
	primes := []int{}
	for i := 2; len(primes) <= n; i++ {
		if prime(i) {
			primes = append(primes, i)
		}
	}
	return primes[ix-1] == n
}

func sheldon(n, idx int) bool {

	if !prime(n) {
		return false
	}

	r := reverse(n)
	if !prime(r) {
		return false
	}

	s := strconv.Itoa(n)

	if len(s) < 2 {
		return false
	}

	product := int(s[0]-'0') * int(s[1]-'0')
	if product != idx {
		return false
	}
	idxReserve := reverse(idx)
	return isPermutation(idx, idxReserve, r)
}
