package main

import (
	"fmt"

	"sort"
)

func main() {
	// Create a slice of integers
	var numbers []int = []int{4, 2, 9, 6,
		23, 12, 34, 0, 1}
	// Sort the slice in ascending order
	sort.Ints(numbers)
	fmt.Println("Sorted numbers: ", numbers)
	// Calculate the median of the sorted slice
	median := calculateMedian(numbers)
	fmt.Printf("Median: %.2f\n", median)
}
func calculateMedian(numbers []int) float64 {
	n := len(numbers)
	if n%2 == 0 {
		// If the length of the slice is even, the median is the average of the two
		// middle numbers
		middle1 := numbers[n/2-1]
		middle2 := numbers[n/2]
		return float64(middle1+middle2) / 2
	} else {
		// If the length of the slice is odd, the median is the middle number
		return float64(numbers[n/2])
	}
}
