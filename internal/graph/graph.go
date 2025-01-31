package graph

func DFS(startNode int, endNode int, graph map[int][]int, path []int, allPaths *[][]int) {
	path = append(path, startNode)
	if startNode == endNode {
		newPath := make([]int, len(path))
		copy(newPath, path)
		*allPaths = append(*allPaths, newPath)
		return
	}
	for _, neighbour := range graph[startNode] {
		DFS(neighbour, endNode, graph, path, allPaths)
	}
}

func getAdjacencyList(edges [][]int) map[int][]int {
	adjList := make(map[int][]int)
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		adjList[u] = append(adjList[u], v)
	}
	return adjList
}
