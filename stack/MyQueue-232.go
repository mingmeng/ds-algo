package stack

type MyQueue struct {
	InStack  []int
	OutStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		InStack:  make([]int, 0),
		OutStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.InStack = append(this.InStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.OutStack) > 0 {
		newLen := len(this.OutStack) - 1
		result := this.OutStack[newLen]
		this.OutStack = this.OutStack[:newLen]
		return result
	} else {
		for i := len(this.InStack) - 1; i > 0; i-- {
			this.OutStack = append(this.OutStack, this.InStack[i])
		}
		result := this.InStack[0]
		this.InStack = make([]int, 0)
		return result
	}
}

func (this *MyQueue) Peek() int {
	result := this.Pop()
	this.OutStack = append(this.OutStack, result)
	return result
}

func (this *MyQueue) Empty() bool {
	return len(this.InStack) == 0 && len(this.OutStack) == 0
}
