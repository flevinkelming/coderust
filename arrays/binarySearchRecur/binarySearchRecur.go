package main

import "fmt"

// Requirement: Given a sorted array of integers, return the index of the given key.
//     Return -1 if not found.
func binarySearch(arr []int, key, low, high int) int {
	if low > high {
		return -1
	}

	mid := low + (high-low)/2
	if arr[mid] == key {
		return mid
	} else if arr[mid] > key {
		return binarySearch(arr, key, low, mid-1)
	} else {
		return binarySearch(arr, key, mid+1, high)
	}
}

func binarySearchRecur(arr []int, key int) int {
	return binarySearch(arr, key, 0, len(arr)-1)
}

func main() {
	arr := []int{1, 10, 13, 24, 35, 43, 57, 67, 88, 91, 101}
	fmt.Println(binarySearchRecur(arr, 13))
}
