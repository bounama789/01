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
		mean := math.Round(statistics.Mean(data))
		variance := int(math.Round(statistics.Variance(data)))
		std := math.Round(statistics.Stdeviation(data))
		median := math.Round(statistics.Median(data))

		fmt.Printf("Average: %v\nMedian: %v\nVariance: %v\nStandard Deviation: %v \n", mean, median, variance, std)
	}
}
