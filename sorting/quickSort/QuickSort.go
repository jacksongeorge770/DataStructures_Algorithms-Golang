package main

import (
	"fmt"
)

// QuickSort function
func quickSort(arr []int, low, high int) {
	if low < high {
		// Partition the array
		pi := partition(arr, low, high)

		// Recursively apply QuickSort to the left and right partitions
		quickSort(arr, low, pi-1)  // Sort left subarray
		quickSort(arr, pi+1, high) // Sort right subarray
	}
}

// Partition function (Lomuto partition scheme)
func partition(arr []int, low, high int) int {
	pivot := arr[high] // Choosing the last element as the pivot
	i := low - 1       // i tracks the position for swapping

	for j := low; j < high; j++ {
		// If current element is smaller than or equal to the pivot
		if arr[j] < pivot {
			i++
			// Swap arr[i] and arr[j]
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Swap pivot with element at i+1 to place it in the correct position
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1 // Returning the partition index
}

func main() {
	arr := []int{8, 1, 3, 9, 5, 4, 7, 6, 2}
	fmt.Println("Original array:", arr)

	quickSort(arr, 0, len(arr)-1)

	fmt.Println("Sorted array:", arr)
}

