package heap

import (
	"fmt"
	"heap/pkg/prettyprint"
	"time"

	"golang.org/x/exp/constraints"
)

const SLEEP_TIME = 2 * time.Second

func InteractivelyPushBack[T constraints.Ordered](heap []T, value T) []T {
	out := append(heap, value)
	fmt.Printf(">>> Pushing back %v...\n", value)
	prettyprint.PrintHeap(out)
	time.Sleep(SLEEP_TIME)
	fmt.Printf(">>> Going up with %v...\n", value)
	InteractivelyGoUp(&out, uint(len(heap)))
	return out
}

func InteractivelyPushFront[T constraints.Ordered](heap []T, value T) []T {
	out := append([]T{value}, heap...)
	fmt.Printf(">>> Pushing front %v...\n", value)
	prettyprint.PrintHeap(out)
	time.Sleep(SLEEP_TIME)
	fmt.Printf(">>> Going down with %v...\n", value)
	InteractivelyGoDown(&out, 0)
	return out
}

func InteractivelyGoDown[T constraints.Ordered](heap *[]T, rootIndex uint) {
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
		fmt.Printf(">>> Swapping elements %v (index %d) and %v (index %d)\n", (*heap)[rootIndex], rootIndex, (*heap)[largestValue], largestValue)
		prettyprint.PrintHeap(*heap)
		time.Sleep(SLEEP_TIME)
		InteractivelyGoDown(heap, largestValue)
	}
}

func InteractivelyGoUp[T constraints.Ordered](heap *[]T, rootIndex uint) {
	if rootIndex == 0 {
		return
	}
	parent := (rootIndex - 1) / 2
	if (*heap)[rootIndex] <= (*heap)[parent] {
		fmt.Printf(">>> Swapping elements %v (index %d) and %v (index %d)\n", (*heap)[rootIndex], rootIndex, (*heap)[parent], parent)
		(*heap)[rootIndex], (*heap)[parent] = (*heap)[parent], (*heap)[rootIndex]
		prettyprint.PrintHeap(*heap)
		time.Sleep(SLEEP_TIME)
		InteractivelyGoUp(heap, parent)
	}
}

func InteractivelyPushBackMax[T constraints.Ordered](heap []T, value T) []T {
	out := append(heap, value)
	fmt.Printf(">>> Pushing back %v...\n", value)
	prettyprint.PrintHeap(out)
	time.Sleep(SLEEP_TIME)
	fmt.Printf(">>> Going up with %v...\n", value)
	InteractivelyGoUpMax(&out, uint(len(heap)))
	return out
}

func InteractivelyPushFrontMax[T constraints.Ordered](heap []T, value T) []T {
	out := append([]T{value}, heap...)
	fmt.Printf(">>> Pushing front %v...\n", value)
	prettyprint.PrintHeap(out)
	time.Sleep(SLEEP_TIME)
	fmt.Printf(">>> Going down with %v...\n", value)
	InteractivelyGoDownMax(&out, 0)
	return out
}

func InteractivelyGoDownMax[T constraints.Ordered](heap *[]T, rootIndex uint) {
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
		fmt.Printf(">>> Swapping elements %v (index %d) and %v (index %d)\n", (*heap)[rootIndex], rootIndex, (*heap)[largestValue], largestValue)
		(*heap)[rootIndex], (*heap)[largestValue] = (*heap)[largestValue], (*heap)[rootIndex]
		prettyprint.PrintHeap(*heap)
		time.Sleep(SLEEP_TIME)
		InteractivelyGoDownMax(heap, largestValue)
	}
}

func InteractivelyGoUpMax[T constraints.Ordered](heap *[]T, rootIndex uint) {
	if rootIndex == 0 {
		return
	}
	parent := (rootIndex - 1) / 2
	if (*heap)[rootIndex] >= (*heap)[parent] {
		fmt.Printf(">>> Swapping elements %v (index %d) and %v (index %d)\n", (*heap)[rootIndex], rootIndex, (*heap)[parent], parent)
		(*heap)[rootIndex], (*heap)[parent] = (*heap)[parent], (*heap)[rootIndex]
		prettyprint.PrintHeap(*heap)
		time.Sleep(SLEEP_TIME)
		InteractivelyGoUpMax(heap, parent)
	}
}
