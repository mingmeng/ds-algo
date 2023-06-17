package backtrack

import "fmt"

var letterCombinationsMap = [10]string{
	"",     // 0
	"",     // 1
	"abc",  // 2
	"def",  // 3
	"ghi",  // 4
	"jkl",  // 5
	"mno",  // 6
	"pqrs", // 7
	"tuv",  // 8
	"wxyz", // 9
}

var letterCombinationsPath []byte
var letterCombinationsResults []string

func letterCombinationsDFS(digits string, start int) {
	if start == len(digits) {
		letterCombinationsResults = append(letterCombinationsResults, fmt.Sprint(letterCombinationsPath))
		return
	}

	letters := letterCombinationsMap[digits[start]-'0']

	for i := 0; i < len(letters); i++ {
		letterCombinationsPath = append(letterCombinationsPath, letters[i])
		letterCombinationsDFS(digits, start+1)
		letterCombinationsPath = letterCombinationsPath[:len(letterCombinationsPath)-1]
	}
}

func letterCombinations(digits string) []string {
	letterCombinationsResults = make([]string, 0)
	letterCombinationsPath = make([]byte, 0)
	letterCombinationsDFS(digits, 0)
	return letterCombinationsResults
}
