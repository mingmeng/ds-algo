package stack

type CQueue struct {
	InStack  []int
	OutStack []int
}

func Constructor() CQueue {
	return CQueue{
		InStack:  make([]int, 0),
		OutStack: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.InStack = append(this.InStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.OutStack) == 0 {
		for i := len(this.InStack); i > 0; i-- {
			this.OutStack = append(this.OutStack, this.InStack[i])
		}
		result := this.OutStack[0]
		this.InStack = make([]int, 0)
		return result
	}
	length := len(this.OutStack)
	// 先拿到结果
	result := this.OutStack[length-1]
	// 将tail pop出去
	this.OutStack = this.OutStack[0 : length-1]
	return result
}
