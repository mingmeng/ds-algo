package string

import "fmt"

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return fmt.Sprint(numerator)
	}

	if numerator%denominator == 0 {
		return fmt.Sprint(numerator / denominator)
	}

	result := make([]byte, 0)

	if numerator*denominator < 0 {
		result = append(result, '-')
	}
	numerator = abs(numerator)
	denominator = abs(denominator)

	integerPart := numerator / denominator
	result = append(result, []byte(fmt.Sprint(integerPart))...)

	numerator = numerator % denominator
	result = append(result, '.')

	m := make(map[int]int)

	for numerator != 0 && m[numerator] == 0 {
		m[numerator] = len(result)
		numerator = numerator * 10
		result = append(result, '0'+byte(numerator/denominator))
		numerator = numerator % denominator
	}

	if numerator > 0 {
		insertIndex := m[numerator]
		result = append(result[:insertIndex], '(')
		result = append(result, result[insertIndex:]...)
		result = append(result, ')')
	}

	return string(result)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
