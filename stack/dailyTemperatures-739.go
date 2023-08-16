package stack

import "container/list"

func dailyTemperatures(temperatures []int) []int {
	stack := list.New()
	stack.PushBack(0)

	result := make([]int, len(temperatures))
	for i := 1; i < len(temperatures); i++ {
		top := stack.Back().Value.(int)
		if temperatures[i] > temperatures[top] {

			for stack.Len() > 0 && temperatures[i] > temperatures[stack.Back().Value.(int)] {
				stTop := stack.Back()
				result[stTop.Value.(int)] = i - stTop.Value.(int)
				stack.Remove(stTop)
			}
		}
		stack.PushBack(i)
	}

	return result
}
