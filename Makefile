i:
				go build -o interactive-heap cmd/interactive/main.go

p:
				go build -o performance-heap cmd/performance/main.go

run-i: i 
				./interactive-heap

run-p: p
				./performance-heap
