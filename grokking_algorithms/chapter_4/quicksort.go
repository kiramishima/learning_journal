package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	// Pick a pivot
	pivotIndex := 0

	// Move the pivot to the right
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	arr[left], arr[right] = arr[right], arr[left]

	quicksort(arr[:left])
	quicksort(arr[left+1:])

	return arr
}

func main() {
	fmt.Println(quicksort([]int{10, 5, 2, 3}))
}
