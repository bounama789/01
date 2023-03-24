package utils

import "sort"

type Stack struct {
	Alias rune
	Data  []int
}

var Operations = new([]string)
var optList = new([]string)

func (stack Stack) PushToTopOf(otherStack Stack) (Stack, Stack) {
	var temp = make([]int, len(stack.Data)+1)
	temp[0] = otherStack.Data[0]
	otherStack.Data = otherStack.Data[1:]
	for i, v := range stack.Data {
		temp[i+1] = v
	}
	stack.Data = temp

	*Operations = append(*Operations, "p"+string(stack.Alias))

	return stack, otherStack
}

func (stack Stack) SwapfirstTwo() {
	stack.Data[0], stack.Data[1] = stack.Data[1], stack.Data[0]

	op := "s" + string(stack.Alias)
	temp := *Operations

	if len(*Operations) > 0 && isLastSame(op) {
		temp = temp[:len(temp)-1]
		*Operations = append(temp, "ss")
	} else {
		*Operations = append(temp, op)
	}

}

func (stack Stack) RotateStack() Stack {
	temp := stack.Data[0]
	stack.Data = stack.Data[1:]
	stack.Data = append(stack.Data, temp)

	op := "r" + string(stack.Alias)
	opTemp := *Operations

	if len(*Operations) > 0 && isLastSame(op) {
		opTemp = opTemp[:len(opTemp)-1]
		*Operations = append(opTemp, "rr")
	} else {
		*Operations = append(opTemp, op)
	}
	return stack
}

func (stack Stack) ReverseRotateStack() Stack {
	temp := make([]int, len(stack.Data))
	temp[0] = stack.Data[len(stack.Data)-1]
	for i, v := range stack.Data[:len(stack.Data)-1] {
		temp[i+1] = v
	}
	stack.Data = temp

	op := "rr" + string(stack.Alias)
	opTemp := *Operations

	if len(*Operations) > 0 && isLastSame(op) {
		opTemp = opTemp[:len(opTemp)-1]
		*Operations = append(opTemp, "rrr")
	} else {
		*Operations = append(opTemp, op)
	}
	return stack
}

func isLastSame(op string) bool {
	temp := *Operations
	lastOp := temp[len(temp)-1]
	if op[:len(op)-1] == lastOp[:len(lastOp)-1] && op[len(op)-1] != lastOp[len(lastOp)-1] {
		return true
	}
	return false
}

func Median(data []int) int {
	temp := make([]int, len(data))
	copy(temp, data)
	sort.Ints(temp)
	return temp[int(len(temp)/2)]
}

func sortIntegerTable(table []int) {
	for i := range table {
		for u := range table {
			if table[i] < table[u] {
				table[i], table[u] = table[u], table[i]
			}
		}
	}
}

func checker(a, b int) int {
	return a - b
}

func isSortedAsc(f func(a, b int) int, a []int) bool {
	asc := true
	for i, v := range a {
		if i < len(a)-1 {
			if checker(v, a[i+1]) > 0 {
				asc = false
			}
		}
	}
	return asc
}

func isSortedDesc(f func(a, b int) int, a []int) bool {
	for i, v := range a {
		if i < len(a)-1 {
			if checker(v, a[i+1]) < 0 {
				return false
			}
		}
	}
	return true
}

var i int

func Sort(a Stack, b Stack) (bool, Stack) {
	var chunck [][]int

	for !isSortedAsc(checker, a.Data) {
		median := Median(a.Data)
		if len(a.Data) > 2 {
			var c []int
			if len(a.Data) > 1 && a.Data[0] > a.Data[1] {
				a.SwapfirstTwo()
				if isSortedAsc(checker,a.Data) {
					break
				}
			}
			for valLessthanMedian(a.Data, median) {
				if a.Data[0] < median {
					c = append(c, a.Data[0])
					b, a = b.PushToTopOf(a)
					if isSortedAsc(checker,a.Data) {
						break
					}
					if len(b.Data) > 1 && b.Data[0] < b.Data[1] {
						b.SwapfirstTwo()
						if isSortedAsc(checker,a.Data) {
							break
						}
					}

				} else {
					if a.Data[len(a.Data)-1] < median {
						a = a.ReverseRotateStack()
						if isSortedAsc(checker,a.Data) {
							break
						}
					} else {
						if len(a.Data) == 3 && a.Data[1] < median {
							a.SwapfirstTwo()
						}else {
							a = a.RotateStack()

						}
						if isSortedAsc(checker,a.Data) {
							break
						}
					}
				}
			}

			sortIntegerTable(c)
			chunck = append(chunck, c)
		} else {
			if !isSortedAsc(checker, a.Data) {
				a.SwapfirstTwo()
			}
			i = len(chunck) - 1
			// for len(b.Data) > 0 {
			// 	median := Median(chunck[i])
			// 	if len(chunck[i]) == 1 {
			// 		a, b = a.PushToTopOf(b)
			// 		i--
			// 		continue
			// 	}
			// 	if !isSortedDesc(checker, b.Data) {
			// 		var nrotate int

			// 		for valGreatthanMedian(b.Data, median) {
			// 			if b.Data[0] > median {
			// 				a, b = a.PushToTopOf(b)

			// 				if len(chunck[i]) > 1 {
			// 					chunck[i] = chunck[i][:len(chunck[i])-1]
			// 				}
			// 			} else {
			// 				b = b.RotateStack()
			// 				nrotate++
			// 			}
			// 		}
			// 		if len(chunck[i]) == 2 {
			// 			b.SwapfirstTwo()
			// 			a, b = a.PushToTopOf(b)
			// 			a, b = a.PushToTopOf(b)
			// 			chunck[i] = nil
			// 		}
			// 		if nrotate != 0 {
			// 			for i := 0; i < len(chunck[i]); i++ {
			// 				b = b.ReverseRotateStack()
			// 				continue
			// 			}
			// 		}
			// 		if i > 0 {
			// 			i--
			// 		}
			// 	} else {
			// 		a, b = a.PushToTopOf(b)
			// 	}
			// }

		}

	}

	if len(b.Data) > 0 {
		for len(b.Data) > 0 {
			median := Median(chunck[i])
			if len(chunck[i]) == 1 {
				a, b = a.PushToTopOf(b)
				i--
				continue
			}
			if !isSortedDesc(checker, b.Data) {
				var nrotate int

				for valGreatthanMedian(b.Data, median) {
					if b.Data[0] > median {
						a, b = a.PushToTopOf(b)

						if len(chunck[i]) > 1 {
							chunck[i] = chunck[i][:len(chunck[i])-1]
						}
					} else {
						b = b.RotateStack()
						nrotate++
					}
				}
				if len(chunck[i]) == 2 {
					b.SwapfirstTwo()
					a, b = a.PushToTopOf(b)
					a, b = a.PushToTopOf(b)
					chunck[i] = nil
				}
				if nrotate != 0 {
					for i := 0; i < len(chunck[i]); i++ {
						b = b.ReverseRotateStack()
						continue
					}
				}
				if i > 0 {
					i--
				}
			} else {
				a, b = a.PushToTopOf(b)
			}
		}

	}
	return true, a
}

func valLessthanMedian(data []int, median int) bool {
	for _, v := range data {
		if v < median {
			return true
		}
	}
	return false
}

func valGreatthanMedian(data []int, median int) bool {
	for _, v := range data {
		if v > median {
			return true
		}
	}
	return false
}
