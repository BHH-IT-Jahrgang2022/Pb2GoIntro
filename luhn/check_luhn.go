package main

import (
	"fmt"
	"strconv"
)

func check_luhn(num string) bool {
	sumNum := 0
	isSecond := false

	for i := 0; i < len(num); i++ {

		digit, err := strconv.Atoi(string(num[i]))
		if err != nil {
			return false
		}

		if isSecond {
			digit = digit * 2
		}

		sumNum += digit / 10
		sumNum += digit % 10

		isSecond = !isSecond
	}

	return (sumNum%10 == 0)
}

func main() {
	fmt.Println(check_luhn("79927398713"))
}
