package utils

type Stack struct {
	alias rune
	data []int
}

var operations = new([]string)
var optList = new([]string)

func (stack Stack) PushToTopOf(otherStack Stack) {
	var temp = make([]int, len(otherStack.data)+1)
	temp[0] = stack.data[0]
	stack.data = stack.data[1:]
	for i, v := range otherStack.data {
		temp[i+1] = v
	}
	otherStack.data = temp
	
	*operations = append(*operations, "p"+string(otherStack.alias))
}

func (stack Stack) SwapfirstTwo() {
	stack.data[0],stack.data[1] = stack.data[1],stack.data[0]

	op:="s"+string(stack.alias)
	temp := *operations

	if len(*operations )> 0 && isLastSame(op){
		temp = temp[:len(temp)-1]
		*operations = append(temp, op)
	}
	*operations = append(temp, op)
}

func isLastSame(op string) bool {
	temp := *operations
	lastOp := temp[len(temp)-1]
	if op[:len(op)-1] ==  lastOp[:len(lastOp)-1] && op[len(op)-1] != lastOp[len(lastOp)-1] {
		return true
	}
	return false
}