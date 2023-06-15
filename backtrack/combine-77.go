package backtrack

var results [][]int

func combineBacktrack(n, k, starts int, path []int) {
	if k == len(path) {
		// 注意 slice append的是引用，后续的修改可能会修改到这份引用，因此记得cp一份内存
		tmp := make([]int, k)
		copy(tmp, path)
		results = append(results, tmp)
		return
	}

	for i := starts; i <= n; i++ {
		path = append(path, i)
		combineBacktrack(n, k, i+1, path)
		path = path[:len(path)-1]
	}
}

func combine(n int, k int) [][]int {
	results = make([][]int, 0)
	combineBacktrack(n, k, 1, make([]int, 0))
	return results
}
