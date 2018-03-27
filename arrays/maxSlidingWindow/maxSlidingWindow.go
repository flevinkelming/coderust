package main

import "fmt"

// Requirement: Given a large array of integers and a window of size 'w',
//     find the current maximum in the window as the window slides through the entire array.
func maxSlidingWindow(arr []int, windowSize int) []int {
	// Assume 0 <= windowSize <= arr
	var result []int
	if windowSize > len(arr) {
		return result
	}

	heap := make([]int, windowSize)
	for i := 0; i < windowSize; i++ {
		for len(heap) > 0 && arr[heap[len(heap)-1]] <= arr[i] {
			heap = heap[:len(heap)-1]
		}

		heap = append(heap, i)
	}

	result = append(result, arr[heap[0]])

	for i := 3; i < len(arr); i++ {
		for len(heap) > 0 && arr[heap[len(heap)-1]] <= arr[i] {
			heap = heap[:len(heap)-1]
		}

		if len(heap) > 0 && heap[0] <= i-windowSize {
			heap = heap[1:]
		}

		heap = append(heap, i)
		result = append(result, arr[heap[0]])
	}

	return result
}

func main() {
	arr := []int{-4, 2, -5, 1, -1, 6}
	fmt.Println(maxSlidingWindow(arr, 3))
}
