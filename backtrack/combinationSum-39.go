package backtrack

import "sort"

var combinationSumResult [][]int
var combinationSumCollection []int

func combinationSum(candidates []int, target int) [][]int {
	combinationSumResult = make([][]int, 0)
	combinationSumCollection = make([]int, 0)
	sort.Ints(candidates)
	combinationSumDFS(candidates, target, 0, 0)

	return combinationSumResult
}

func combinationSumDFS(candidates []int, target int, sum int, start int) {
	if sum == target {
		tmp := make([]int, len(combinationSumCollection))
		copy(tmp, combinationSumCollection)
		combinationSumResult = append(combinationSumResult, tmp)
		return
	}

	for i := start; i < len(candidates) && sum+candidates[i] < target; i++ {
		sum += candidates[i]
		combinationSumCollection = append(combinationSumCollection, candidates[i])
		combinationSumDFS(candidates, target, sum, i)
		sum -= candidates[i]
		combinationSumCollection = combinationSumCollection[:len(combinationSumCollection)-1]
	}

}
