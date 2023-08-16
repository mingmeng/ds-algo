package maths

import "fmt"

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	if numerator%denominator == 0 {
		return fmt.Sprint(numerator % denominator)
	}

	var result string
	// 先判断符号
	if numerator*denominator < 0 {
		result = "-"
	}
	numerator = abs(numerator)
	denominator = abs(denominator)
	// 处理整数部分
	result += fmt.Sprint(numerator/denominator, ".")

	numerator = numerator % denominator
	m := map[int]int{}

	for numerator != 0 && m[numerator] == 0 {
		m[numerator] = len(result)
		numerator = numerator * 10
		result = result + fmt.Sprintf("%d", numerator/denominator)
		numerator = numerator % denominator
	}

	// 说明有循环小数
	if numerator > 0 {
		insertIndex := m[numerator]

		return result[:insertIndex] + "(" + result[insertIndex:] + ")"
	}

	return result

}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
