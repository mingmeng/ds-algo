package stack

func isValid(s string) bool {
	leftStack := make([]byte, 0)

	pair := map[byte]byte{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[', '{', '(':
			leftStack = append(leftStack, s[i])
		case ']', '}', ')':
			leftLength := len(leftStack)
			if leftLength == 0 {
				return false
			}
			if pair[leftStack[leftLength-1]] != s[i] {
				return false
			}
			leftStack = leftStack[:leftLength-1]
		}
	}
	if len(leftStack) > 0 {
		return false
	}

	return true
}
