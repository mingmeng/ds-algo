package stack

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	// result:=0
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			lengthOfStack := len(stack)
			num1 := stack[lengthOfStack-1]
			num2 := stack[lengthOfStack-2]

			var tmp int
			switch tokens[i] {
			case "+":
				tmp = num1 + num2
			case "-":
				tmp = num1 - num2
			case "*":
				tmp = num1 * num2
			case "/":
				tmp = num1 / num2
			}
			stack = append(stack, tmp)
		} else {
			factor := 0
			isNega := false
			for j := 0; j < len(tokens[i]); j++ {
				ch := tokens[i][j]
				if ch == '-' {
					isNega = true
				} else {
					factor = factor*10 + int(ch-'0')
				}
			}

			if isNega {
				factor = 0 - factor
			}

			stack = append(stack, factor)
		}
	}

	return stack[0]
}
