package salsa_festival

func canAllVeggiesCome(vegNum int, requirements [][]int) bool {
	vegGraph := make([][]int, vegNum)
	for _, req := range requirements {
		toVeg := req[0]
		fromVeg := req[1]
		vegGraph[fromVeg] = append(vegGraph[fromVeg], toVeg)
	}

	visited := make(map[int]bool)
	for veg := 0; veg < vegNum; veg++ {
		if visited[veg] {
			continue
		}
		pathVertices := make(map[int]bool)
		if dfs(vegGraph, veg, pathVertices, visited) {
			return false
		}
	}
	return true
}

func dfs(graph [][]int, vertex int, pathVertices map[int]bool, visited map[int]bool) bool {
	if pathVertices[vertex] {
		return true
	}
	if visited[vertex] {
		return false
	}
	pathVertices[vertex] = true
	visited[vertex] = true
	for _, childV := range graph[vertex] {
		if dfs(graph, childV, pathVertices, visited) {
			return true
		}
	}
	pathVertices[vertex] = false
	return false
}
