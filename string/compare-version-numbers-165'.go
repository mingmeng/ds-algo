package main

import (
	"strconv"
	"strings"
)

// new space
func compareVersion(version1 string, version2 string) int {
	v1Seg := strings.Split(version1, ".")
	v2Seg := strings.Split(version2, ".")

	longestLength := len(v1Seg)
	if len(v1Seg) < len(v2Seg) {
		longestLength = len(v2Seg)
	}

	for i := 0; i < longestLength; i++ {
		v1n := 0
		v2n := 0
		if len(v1Seg) >= i {
			v1n, _ = strconv.Atoi(v1Seg[i])
		}
		if len(v2Seg) >= i {
			v2n, _ = strconv.Atoi(v2Seg[i])
		}

		if v1n > v2n {
			return 1
		} else if v1n < v2n {
			return -1
		}
	}

	return 0
}

// two pointer
func compareVersion2(version1 string, version2 string) int {
	var i, j int
	for i < len(version1) || j < len(version2) {

		var number1, number2 int
		for ; i < len(version1) && version1[i]-'.' != 0; i++ {
			number1 = number1*10 + int(version1[i]-'0')
		}
		// 跳过.
		i++
		for ; j < len(version2) && version2[j]-'.' != 0; j++ {
			number2 = number2*10 + int(version2[j]-'0')
		}
		//跳过.
		j++
		if number1 > number2 {
			return 1
		} else if number2 > number1 {
			return -1
		}
	}

	return 0
}
