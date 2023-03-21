package main

import "fmt"

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
}

func backspaceCompare(s string, t string) bool {
	fmt.Println(backspace(s))
	fmt.Println(backspace(t))
	return backspace(s) == backspace(t)
}

func backspace(s1 string) string {
	var slow = 0
	var fast = 0

	s := []byte(s1)

	for fast < len(s) {
		if s[fast] == '#' {
			if slow < 0 {
				slow = -1
			} else {
				slow -= 1
			}
		} else {
			s[slow+1] = s[fast]
			slow += 1
		}
		fast += 1
	}
	if slow == -1 {
		return ""
	}
	return string(s[:slow])
}
