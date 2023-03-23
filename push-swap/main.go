package main

import (
	"fmt"
	"push-swap/utils"
)



func main()  {
	var a = utils.Stack {
		Alias: 'a',
		Data: []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14},
	}
	var b = utils.Stack{
		Alias: 'b',
		Data: nil,
	}

	a.SwapfirstTwo()
	
	b,a = b.PushToTopOf(a)
	
	b,a = b.PushToTopOf(a)
	b,a = b.PushToTopOf(a)

a.SwapfirstTwo()
b.SwapfirstTwo()
a=a.RotateStack()
b=b.RotateStack()


	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(*utils.Operations)

}