package main

import "fmt"

// mergeSort sorts an array using iterative merge sort
func mergeSort(arr []int) []int {
    n := len(arr)
    
    // Iteratively merge subarrays of increasing size
    for size := 1; size < n; size *= 2 {
        for start := 0; start < n; start += 2 * size {
            mid := start + size - 1
            end := min(start + 2*size - 1, n-1)

            // Merge the subarrays
            merge(arr, start, mid, end)
        }
    }
    return arr
}

// merge merges two sorted subarrays into one sorted subarray
func merge(arr []int, start, mid, end int) {
    left, right := start, mid+1
    temp := make([]int, end-start+1)
    i := 0
    
    // Merge the two subarrays into the temporary array
    for left <= mid && right <= end {
        if arr[left] <= arr[right] {
            temp[i] = arr[left]
            left++
        } else {
            temp[i] = arr[right]
            right++
        }
        i++
    }

    // Copy remaining elements from the left subarray
    for left <= mid {
        temp[i] = arr[left]
        left++
        i++
    }

    // Copy remaining elements from the right subarray
    for right <= end {
        temp[i] = arr[right]
        right++
        i++
    }

    // Copy the sorted subarray back to the original array
    for i := 0; i < len(temp); i++ {
        arr[start+i] = temp[i]
    }
}

// min returns the smaller of two integers
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    // Test the mergeSort function
    nums := []int{5, 2, 8, 1, 3, 7, 6, 4}
    fmt.Println("Original Array:", nums)

    sorted := mergeSort(nums)
    fmt.Println("Sorted Array:", sorted)
}
