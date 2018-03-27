package main

import "fmt"

// Requirement: Given a sorted array of integers, return the index of the given key.
//     Return -1 if not found.
func binarySearchIter(arr []int, key int) int {
	if len(arr) == 0 {
		return -1
	}

	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 10, 13, 24, 35, 43, 57, 67, 88, 91, 101}
	fmt.Println(binarySearchIter(arr, 13))
}
