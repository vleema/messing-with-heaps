package main

import (
	"fmt"
	"heap/pkg/heap"
)

func main() {
	slice := []int{6, 1, 8, 7, 3, 9, 5, 2, 4}
	fmt.Println(">>> Original Array: ", slice)
	// heapified := heap.Heapify(slice)
	// fmt.Println(">>> Heapified Array: ", heapified)
	sorted := heap.HeapSort(slice)
	fmt.Println(">>> Ordered Array: ", sorted)
	fmt.Println(">>> Desceding Ordered Array: ", heap.HeapSortDescending(slice))
}
