package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetData(path string) []float64 {
	var data []float64

	file,err := os.Open(path)

	if err != nil {
		fmt.Println("Cannot open file. Possibly file is missing or not in correct format")
		os.Exit(1)
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		value,err1 := strconv.ParseFloat(scanner.Text(),64)
		if err1 != nil {
			continue
		}
		data = append(data, value)
	}
	return data
}