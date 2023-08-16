package graph

func canFinish(numCourses int, prerequisites [][]int) bool {

	graph := buildGraph(prerequisites)

	hasCircle := false

	visited := make([]bool, numCourses)
	path := make([]bool, numCourses)
	var dfs func(graph map[int][]int, to int)
	dfs = func(graph map[int][]int, to int) {

		if visited[to] {
			return
		}

		if path[to] {
			hasCircle = true
			return
		}

		visited[to] = true
		path[to] = true
		for _, next := range graph[to] {
			dfs(graph, next)
		}
		path[to] = false

	}

	for from := range graph {
		dfs(graph, from)
	}

	return !hasCircle
}

func buildGraph(prerequisites [][]int) map[int][]int {
	result := make(map[int][]int)
	for _, pair := range prerequisites {
		from := pair[0]
		to := pair[1]

		result[from] = append(result[from], to)
	}

	return result
}
