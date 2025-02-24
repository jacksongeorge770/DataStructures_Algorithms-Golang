package main

import "fmt"

// MergeSort function
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Find the middle point and split the array into two halves
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])   // Sort the left half
	right := mergeSort(arr[mid:])  // Sort the right half

	// Merge the sorted halves
	return merge(left, right)
}

// Merge function to merge two sorted arrays
func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0

	// Merge the two sorted arrays into a single sorted array
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements if any
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	arr := []int{8, 1, 3, 9, 5, 4, 7, 6, 2}
	fmt.Println("Original array:", arr)

	// Call mergeSort
	sortedArr := mergeSort(arr)

	fmt.Println("Sorted array:", sortedArr)
}

