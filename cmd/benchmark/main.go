package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/wcharczuk/go-chart/v2"
)

func main() {
	file, err := os.Open("benchmark_results.txt")
	if err != nil {
		fmt.Println("Error while opening the file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Map to store results for each benchmark
	results := map[string]map[int]float64{
		"PushBack":    {},
		"PushBackMax": {},
		"PushFront":   {},
	}

	// Regex to parse benchmark results
	re := regexp.MustCompile(`^(Benchmark(PushBack|PushBackMax|PushFront)(\d+))-\d+\s+\d+\s+(\d+) ns/op`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			benchmarkType := matches[2]
			size, _ := strconv.Atoi(matches[3]) // Input size
			timePerOp, _ := strconv.ParseFloat(matches[4], 64)

			// Store results
			results[benchmarkType][size] = timePerOp
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading the benchmark file:", err)
		os.Exit(1)
	}

	styles := map[string]chart.Style{
		"PushBack": {
			StrokeColor: chart.ColorBlue,
			StrokeWidth: 2,
		},
		"PushBackMax": {
			StrokeColor:     chart.ColorRed,
			StrokeWidth:     2,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		"PushFront": {
			StrokeColor: chart.ColorYellow,
			StrokeWidth: 2,
			Hidden:      true,
		},
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "Input Size",
		},
		YAxis: chart.YAxis{
			Name: "Time per Operation (ns/op)",
		},
		Series: []chart.Series{},
	}

	for benchmarkType, data := range results {
		xValues := []float64{}
		yValues := []float64{}
		for size, time := range data {
			xValues = append(xValues, float64(size))
			yValues = append(yValues, time)
		}
		graph.Series = append(graph.Series, chart.ContinuousSeries{
			Name:    benchmarkType,
			XValues: xValues,
			YValues: yValues,
			Style:   styles[benchmarkType],
		})
	}

	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}

	// Render the chart
	outputFile, err := os.Create("benchmark_results.png")
	if err != nil {
		fmt.Println("Error while creating output file:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	err = graph.Render(chart.PNG, outputFile)
	if err != nil {
		fmt.Println("Error while rendering chart:", err)
		os.Exit(1)
	}

	fmt.Println("Chart saved as benchmark_results.png")
}
