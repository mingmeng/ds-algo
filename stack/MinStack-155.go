package stack

import "math"

type MinStack struct {
	stack     []int
	supStack  []int
	MinFactor int
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{
		stack:     make([]int, 0),
		supStack:  make([]int, 0),
		MinFactor: math.MaxInt64,
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	pushValue := x
	if pushValue > this.Min() {
		pushValue = this.Min()
	}
	this.supStack = append(this.supStack, pushValue)
}

func (this *MinStack) Pop() {
	stackLength := len(this.stack)
	supStackLength := len(this.supStack)
	this.stack = this.stack[0 : stackLength-1]
	this.supStack = this.supStack[0 : supStackLength-1]
}

func (this *MinStack) Top() int {
	length := len(this.stack)
	if length == 0 {
		return 0
	}
	return this.stack[length-1]
}

func (this *MinStack) Min() int {
	length := len(this.supStack)
	if length == 0 {
		return math.MaxInt64
	}
	return this.supStack[length-1]
}
