package main

import (
	"fmt"
	"push-swap/utils"
)

func main() {
	var a = utils.Stack{
		Alias: 'a',
		Data:  []int{3 ,2 ,4 ,9, 5},
	}
	var b = utils.Stack{
		Alias: 'b',
		Data:  nil,
	}

	sorted,sortedStack := utils.Sort(a, b)
	if sorted {
		fmt.Println(*utils.Operations)
		fmt.Println(sortedStack)}
	
}
