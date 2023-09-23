package breadth_first_search

/*
BreadthFirstSearch (BFS)
Algorithm for searching a tree data structure for a node that satisfies a given property.

O(V+E), V is number of vertices, E is number of edges.
For unweighted graphs only.
*/
func BreadthFirstSearch(graph map[string][]string, searchValue, firstNode string) string {
	var deque []string
	var processedNodes []string

	// put successors of first node into deque
	deque = append(deque, graph[firstNode]...)

	for len(deque) > 0 {
		// pop the first node from deque
		node := deque[0]
		deque = append(deque[:0], deque[1:]...)

		// check node among processed to avoid of processing of duplicated nested nodes
		if func(slice []string, needle string) bool {
			for _, i := range slice {
				if i == needle {
					return true
				}
			}
			return false
		}(processedNodes, node) {
			continue
		}

		// success, search node has found
		if found := node == searchValue; found {
			return node
		}

		// put into deque successors of node
		deque = append(deque, graph[node]...)

		// mark node as processed
		processedNodes = append(processedNodes, node)
	}

	return ""
}
