package greddy

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	result := 0
	for i, j := 0, 0; i < len(g); i++ {
		for ; j < len(s); j++ {
			if g[i] <= s[j] {
				result += 1
				j++
				break
			}
		}
	}

	return result
}
