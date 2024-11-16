package prettyprint

import (
	"fmt"
	"math"
	"strings"
)

func PrintHeap[T any](heap []T) {
	n := len(heap)
	if n == 0 {
		return
	}

	height := computeHeight(n)
	lines := []string{}
	queue := []int{0}
	level := 0

	for len(queue) > 0 && level < height {
		levelSize := len(queue)
		nextQueue := []int{}

		spacesBefore, spacesBetween := computeSpacing(height, level)

		nodeLine, nextQueue := buildNodeLine(heap, queue, levelSize, spacesBefore, spacesBetween, nextQueue)
		lines = append(lines, nodeLine)

		if level < height-1 {
			slashLine := buildSlashLine(heap, queue, levelSize, spacesBefore, spacesBetween)
			lines = append(lines, slashLine)
		}

		queue = prepareNextQueue(nextQueue)
		level++
	}

	printLines(lines)
}

func computeHeight(n int) int {
	return int(math.Floor(math.Log2(float64(n)))) + 1
}

func computeSpacing(height, level int) (int, int) {
	spacesBefore := int(math.Pow(2, float64(height-level-1))) - 1
	spacesBetween := int(math.Pow(2, float64(height-level))) - 1
	return spacesBefore, spacesBetween
}

func buildNodeLine[T any](heap []T, queue []int, levelSize, spacesBefore, spacesBetween int, nextQueue []int) (string, []int) {
	nodeLine := strings.Repeat(" ", spacesBefore)
	for i := 0; i < levelSize; i++ {
		index := queue[i]
		if index < len(heap) {
			nodeLine += fmt.Sprintf("%v", heap[index])
			leftChild := 2*index + 1
			rightChild := 2*index + 2
			nextQueue = append(nextQueue, leftChild)
			nextQueue = append(nextQueue, rightChild)
		} else {
			nodeLine += "  "
			nextQueue = append(nextQueue, -1, -1)
		}
		if i < levelSize-1 {
			nodeLine += strings.Repeat(" ", spacesBetween)
		}
	}
	return nodeLine, nextQueue
}

func buildSlashLine[T any](heap []T, queue []int, levelSize, spacesBefore, spacesBetween int) string {
	slashLine := strings.Repeat(" ", spacesBefore)
	for i := 0; i < levelSize; i++ {
		index := queue[i]
		if index < len(heap) {
			leftChild := 2*index + 1
			rightChild := 2*index + 2

			if leftChild < len(heap) {
				slashLine += "/ "
			} else {
				slashLine += " "
			}

			if rightChild < len(heap) {
				slashLine += "\\"
			} else {
				slashLine += " "
			}
		} else {
			slashLine += "  "
		}
		if i < levelSize-1 {
			slashLine += strings.Repeat(" ", spacesBetween)
		}
	}
	return slashLine
}

func prepareNextQueue(nextQueue []int) []int {
	queue := []int{}
	for _, idx := range nextQueue {
		if idx != -1 {
			queue = append(queue, idx)
		}
	}
	return queue
}

func printLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
