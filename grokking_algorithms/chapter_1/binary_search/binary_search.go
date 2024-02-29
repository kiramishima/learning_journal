package main

import "fmt"

func binary_search(list []int, item int) *int {
	// low and high keep track of which
	// part of the list you'll search in
	low := 0
	high := len(list) - 1
	var find *int = nil // if the item don't exist, the default is nil

	// we don't have while in golang, instead we use for loop
	// While you haven¿t narrowed it down to one element
	for low <= high {
		// each time check the middle element
		mid := (low + high) / 2
		guess := list[mid]
		if guess == item { // Found the item
			find = &mid
			break
		} else if guess > item { // The guess was too high
			high = mid - 1
		} else { // The guess was too low
			low = mid + 1
		}
	}
	// return the value
	return find
}

func main() {
	var nums = []int{1, 3, 5, 7, 9}
	fmt.Printf("%v\n", *binary_search(nums, 3)) // returns the index 1, element is in the second slot
	fmt.Printf("%v", binary_search(nums, -1))   // return nil, it wasn't found
}
