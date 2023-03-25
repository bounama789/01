package main

import (
	"fmt"
	"os"
	"push-swap/utils"
)

func main() {
	if len(os.Args) == 2 {
		var a = utils.CheckEntry(os.Args[1])
		var b = utils.Stack{
			Alias: 'b',
			Data:  nil,
		}

		sorted, _ := utils.Sort(a, b)
		opt := *utils.Operations
		if sorted {
			for _, v := range opt {
				fmt.Println(v)
			}
		}
	}
}
