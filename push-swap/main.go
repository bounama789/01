package main

import (
	"fmt"
	"push-swap/utils"
)

func main() {
	var a = utils.Stack{
		Alias: 'a',
		Data:  []int{2, 1, 3, 6, 5, 8},
	}
	var b = utils.Stack{
		Alias: 'b',
		Data:  nil,
	}

	if utils.Sort(a, b) {
		fmt.Println(*utils.Operations)
		fmt.Println(a)
	}
}
