// Package breadthfirstsearch
// Algorithm for searching a tree data structure for a node that satisfies a given property.
// O(V+E), V is number of vertices, E is number of edges.
// For unweighted graphs only.
package breadthfirstsearch

func search(g map[string][]string, search, start string) string {
	var (
		deque     []string
		processed []string
	)

	// put successors of first node into deque
	deque = append(deque, g[start]...)

	for len(deque) > 0 {
		// pop the first node from deque
		node := deque[0]
		deque = append(deque[:0], deque[1:]...)

		// check node among processed to avoid of processing of duplicated nested nodes
		if func(p []string, n string) bool {
			for _, i := range p {
				if i == n {
					return true
				}
			}

			return false
		}(processed, node) {
			continue
		}

		// success, search node has found
		if found := node == search; found {
			return node
		}

		// put into deque successors of node
		deque = append(deque, g[node]...)

		// mark node as processed
		processed = append(processed, node)
	}

	return ""
}
