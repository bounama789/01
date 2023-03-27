package statistics

import (
	"math"
)

func Median(data []float64) float64 {
	sortIntegerTable(data)

	return data[int(len(data)/2)]
}

func Mean(data []float64) float64 {
	total:= tableSum(data)
	return float64(total) / float64(len(data))
}

func Variance(data []float64) int {
	mean := math.Round(Mean(data))
	var temp int
	for _, v := range data {
		temp += int(math.Ceil(math.Pow(v- mean,2)))
		
	}

	return int(float64(temp / len(data)-1)+1)
}

func Stdeviation(data []float64) int {
	return int(math.Round(math.Sqrt(float64(Variance(data)))))
}

func sortIntegerTable(table []float64) {
	for i := range table {
		for u := range table {
			if table[i] < table[u] {
				table[i], table[u] = table[u], table[i]
			}
		}
	}
}

func tableSum(data []float64) int{
	var total int
	for _, v := range data {
		total += int(v)
	}
	return total
}
