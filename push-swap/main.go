package main

import (
	"fmt"
	"push-swap/utils"
)

func main() {
	var a = utils.Stack{
		Alias: 'a',
		Data:  []int{2, 3, 5, 12, 10, 6,9,7,4,1,8,11},
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
