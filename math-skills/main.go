package main

import (
	"fmt"
	"math"
	"math-skills/statistics"
	"math-skills/utils"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		data := utils.GetData(os.Args[1])
		mean := int(math.Round(statistics.Mean(data)))
		variance := int(statistics.Variance(data))
		std := int(statistics.Stdeviation(data))
		median := int(statistics.Median(data))

		fmt.Printf("Average: %v\nMedian: %v\nVariance: %v\nStandard Deviation: %v", mean, median, variance, std)
	}
}
