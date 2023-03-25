package checker

import "push-swap/utils"

func ApplyOperations(operations []string, a utils.Stack) (utils.Stack, utils.Stack) {
	b := utils.Stack{
		Alias: 'b',
		Data:  nil,
	}
	for _, v := range operations {
		switch v {
		case "pa":
			a, b = a.PushToTopOf(b)
		case "pb":
			b, a = b.PushToTopOf(a)
		case "sa":
			a.SwapfirstTwo()
		case "sb":
			b.SwapfirstTwo()
		case "ss":
			a.SwapfirstTwo()
			b.SwapfirstTwo()
		case "ra":
			a = a.RotateStack()
		case "rb":
			b = b.RotateStack()
		case "rr":
			a.RotateStack()
			b.RotateStack()
		case "rra":
			a = a.ReverseRotateStack()
		case "rrb":
			b = b.ReverseRotateStack()
		case "rrr":
			a = a.ReverseRotateStack()
			b = b.ReverseRotateStack()
		}
	}
	return a, b
}
