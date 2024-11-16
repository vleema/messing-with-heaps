package heap

import (
	"golang.org/x/exp/constraints"
)

// Heapify takes a generic a slice and puts it in heap structure
func Heapify[T constraints.Ordered](slice []T) []T {
	out := slice
	for i := len(slice)/2 - 1; i >= 0; i-- {
		GoDown(&out, uint(i))
	}
	return out
}

// HeapifyMax takes a generic a slice and puts it in max-heap structure
func HeapifyMax[T constraints.Ordered](slice []T) []T {
	out := slice
	for i := len(slice)/2 - 1; i >= 0; i-- {
		GoDownMax(&out, uint(i))
	}
	return out
}

// PushBack puts a `value` in the end of `heap`, maintaining the heap property
func PushBack[T constraints.Ordered](heap []T, value T) []T {
	out := append(heap, value)
	GoUp(&out, uint(len(heap)))
	return out
}

// PushBackMax puts a `value` in the end of `heap`, maintaining the max-heap property
func PushBackMax[T constraints.Ordered](heap []T, value T) []T {
	out := append(heap, value)
	GoUpMax(&out, uint(len(heap)))
	return out
}

// PushFront puts a `value` in the end of `heap`, maintaining the heap property
func PushFront[T constraints.Ordered](heap []T, value T) []T {
	out := append([]T{value}, heap...)
	GoDown(&out, 0)
	return out
}

// PushFrontMax puts a `value` in the end of `heap`, maintaining the max-heap property
func PushFrontMax[T constraints.Ordered](heap []T, value T) []T {
	out := append([]T{value}, heap...)
	GoDownMax(&out, 0)
	return out
}

// GoUp sifts an element up the heap to maintain the heap property.
func GoUp[T constraints.Ordered](heap *[]T, rootIndex uint) {
	if rootIndex == 0 {
		return
	}
	parent := (rootIndex - 1) / 2
	if (*heap)[rootIndex] < (*heap)[parent] {
		(*heap)[rootIndex], (*heap)[parent] = (*heap)[parent], (*heap)[rootIndex]
		GoUp(heap, parent)
	}
}

// GoUpMax sifts an element up the heap to maintain the max-heap property
func GoUpMax[T constraints.Ordered](heap *[]T, rootIndex uint) {
	if rootIndex == 0 {
		return
	}
	parent := (rootIndex - 1) / 2
	if (*heap)[rootIndex] > (*heap)[parent] {
		(*heap)[rootIndex], (*heap)[parent] = (*heap)[parent], (*heap)[rootIndex]
		GoUpMax(heap, parent)
	}
}

// GoDown sifts an element down the heap to maintain the heap property
func GoDown[T constraints.Ordered](heap *[]T, rootIndex uint) {
	largestValue := rootIndex
	leftChild := 2*rootIndex + 1
	rightChild := 2*rootIndex + 2
	heapSize := uint(len(*heap))
	if leftChild < heapSize && (*heap)[leftChild] < (*heap)[largestValue] {
		largestValue = leftChild
	}
	if rightChild < heapSize && (*heap)[rightChild] < (*heap)[largestValue] {
		largestValue = rightChild
	}
	if largestValue != rootIndex {
		(*heap)[rootIndex], (*heap)[largestValue] = (*heap)[largestValue], (*heap)[rootIndex]
		GoDown(heap, largestValue)
	}
}

// GoDownMax sifts an element down the heap to maintain the max-heap property
func GoDownMax[T constraints.Ordered](heap *[]T, rootIndex uint) {
	largestValue := rootIndex
	leftChild := 2*rootIndex + 1
	rightChild := 2*rootIndex + 2
	heapSize := uint(len(*heap))
	if leftChild < heapSize && (*heap)[leftChild] > (*heap)[largestValue] {
		largestValue = leftChild
	}
	if rightChild < heapSize && (*heap)[rightChild] > (*heap)[largestValue] {
		largestValue = rightChild
	}
	if largestValue != rootIndex {
		(*heap)[rootIndex], (*heap)[largestValue] = (*heap)[largestValue], (*heap)[rootIndex]
		GoDown(heap, largestValue)
	}
}
