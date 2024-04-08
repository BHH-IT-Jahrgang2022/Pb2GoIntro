package main

import "fmt"

func main() {
	unsorted := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	fmt.Println(mergeSort(unsorted))
}

func mergeSort(unsorted []int) []int {
	if len(unsorted) <= 1 {
		return unsorted
	}

	middle := len(unsorted) / 2
	left := make([]int, middle)
	right := make([]int, len(unsorted)-middle)

	for i := 0; i < len(unsorted); i++ {
		if i < middle {
			left[i] = unsorted[i]
		} else {
			right[i-middle] = unsorted[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}
