package main

import "fmt"

func selectionSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
				// fmt.Printf("%v\n", arr)
			}
		}
	}
	return arr
}

func main() {
	fmt.Println(selectionSort([]int{5, 3, 6, 2, 10}))
}
