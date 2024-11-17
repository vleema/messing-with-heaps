package heap_test

import (
	"testing"

	"heap/pkg/heap"
)

func benchmarkPushBack(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, 0, i)
		for j := 0; j < size; j++ {
			slice = heap.PushBack(slice, i)
		}
	}
}

func BenchmarkPushBack10(b *testing.B)    { benchmarkPushBack(b, 10) }
func BenchmarkPushBack100(b *testing.B)   { benchmarkPushBack(b, 100) }
func BenchmarkPushBack1000(b *testing.B)  { benchmarkPushBack(b, 1000) }
func BenchmarkPushBack10000(b *testing.B) { benchmarkPushBack(b, 10000) }

func benchmarkPushBackMax(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, 0, i)
		for j := 0; j < size; j++ {
			slice = heap.PushBackMax(slice, i)
		}
	}
}

func BenchmarkPushBackMax10(b *testing.B)    { benchmarkPushBackMax(b, 10) }
func BenchmarkPushBackMax100(b *testing.B)   { benchmarkPushBackMax(b, 100) }
func BenchmarkPushBackMax1000(b *testing.B)  { benchmarkPushBackMax(b, 1000) }
func BenchmarkPushBackMax10000(b *testing.B) { benchmarkPushBackMax(b, 10000) }

func benchmarkPushFront(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, 0, i)
		for j := 0; j < size; j++ {
			slice = heap.PushFront(slice, i)
		}
	}
}

func BenchmarkPushFront10(b *testing.B)    { benchmarkPushFront(b, 10) }
func BenchmarkPushFront100(b *testing.B)   { benchmarkPushFront(b, 100) }
func BenchmarkPushFront1000(b *testing.B)  { benchmarkPushFront(b, 1000) }
func BenchmarkPushFront10000(b *testing.B) { benchmarkPushFront(b, 10000) }
