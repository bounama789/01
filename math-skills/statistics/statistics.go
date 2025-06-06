package statistics

import (
	"math"
)

func Median(data []float64) float64 {
	temp := make([]float64, len(data))
	copy(temp,data)
	sortIntegerTable(temp)
	if len(temp)%2==0 {
		return (temp[int(len(temp)/2)] +  temp[int(len(temp)/2)-1]) /2
	}

	return data[int(len(data)/2)]
}

func Mean(data []float64) float64 {
	total := tableSum(data)
	return float64(total) / float64(len(data))
}

func Variance(data []float64) float64 {
	var temp float64
	for _, v := range data {
		mean := Mean(data)

		c := v - mean
		temp += c*c

	}

	return temp/float64(len(data))
}

func Stdeviation(data []float64) float64 {
	return math.Sqrt(Variance(data))
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

func tableSum(data []float64) int {
	var total int
	for _, v := range data {
		total += int(v)
	}
	return total
}

func GetRange(data []float64) float64 {
	temp := make([]float64, len(data))
		copy(temp,data)
	sortIntegerTable(temp)
	return temp[len(temp)-1] - temp[0]
}

func GuessRange(input int,r int) []int {
	return []int{input-r,input+r}
}
