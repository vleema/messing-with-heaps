i:
				go build -o interactive-heap cmd/interactive/main.go

run: i 
				./interactive-heap

bench:
				go test ./cmd/benchmark/tests -bench=. -benchmem  > benchmark_results.txt
				@echo "Saved benchmarks in benchmark_results.txt"

graphs:
				go run cmd/benchmark/main.go
				@echo "Charts generated with sucess benchmark_results.png"

all: i bench graphs
				@echo "Everything done! Binary and graphics ready"
				
