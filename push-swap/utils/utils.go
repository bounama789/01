package utils

type Stack struct {
	Alias rune
	Data []int
}

var Operations = new([]string)
var optList = new([]string)

func (stack Stack) PushToTopOf(otherStack Stack) (Stack,Stack) {
	var temp = make([]int, len(stack.Data)+1)
	temp[0] = otherStack.Data[0]
	otherStack.Data = otherStack.Data[1:]
	for i, v := range stack.Data {
		temp[i+1] = v
	}
	stack.Data = temp
	
	*Operations = append(*Operations, "p"+string(stack.Alias))

	return stack,otherStack
}

func (stack Stack) SwapfirstTwo() {
	stack.Data[0],stack.Data[1] = stack.Data[1],stack.Data[0]

	op:="s"+string(stack.Alias)
	temp := *Operations

	if len(*Operations )> 0 && isLastSame(op){
		temp = temp[:len(temp)-1]
		*Operations = append(temp, "ss")
	} else {
		*Operations = append(temp, op)
	}
	
}

func (stack Stack) RotateStack() Stack {
	temp := stack.Data[0]
	stack.Data = stack.Data[1:]
	stack.Data = append(stack.Data,temp)

	op:="r"+string(stack.Alias)
	opTemp := *Operations

	if len(*Operations )> 0 && isLastSame(op){
		opTemp = opTemp[:len(opTemp)-1]
		*Operations = append(opTemp, "rr")
	} else {
		*Operations = append(opTemp, op)
	}
	return stack
}

func isLastSame(op string) bool {
	temp := *Operations
	lastOp := temp[len(temp)-1]
	if op[:len(op)-1] ==  lastOp[:len(lastOp)-1] && op[len(op)-1] != lastOp[len(lastOp)-1] {
		return true
	}
	return false
}