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
		ra := int(std/mean)
		input := data[98]
		println(int(input))
		gr := statistics.GuessRange(int(input),ra)

		fmt.Printf("Average: %v\nMedian: %v\nVariance: %v\nStandard Deviation: %v\nRange: %v\nGuessed: %v\n\n", mean, median, variance, std,ra,gr)
		println(int(std/mean))
	}
}
