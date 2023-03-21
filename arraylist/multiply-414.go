package arraylist

import "fmt"

func multiply(num1 string, num2 string) string {
	var result string

	if num1 == "0" || num2 == "0" {
		return "0"
	}
	for num1 < num2 {

	}

	var resultInt int64 = 0

	for i := len(num1) - 1; i >= 0; i-- {
		temp1 := num1[i] - '0'

		var tempMulti = temp1
		for j := len(num2); j >= 0; j-- {
			temp2 := num2[j] - '0'

			tempMulti = tempMulti*10 + temp1*temp2
		}

		resultInt += int64(tempMulti)
	}
	result = fmt.Sprint(resultInt)
	return result
}
