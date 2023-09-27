package dijkstra_search

import "math"

/*
dijkstraSearch
For weighted graph only, where weighted graph is a graph in which each edge has assigned the weight (value)

- each vertex has minimal-known distance (costs) to V(start)
- V(start) is assigned as 0, V(finish) is assigned as infinity
*/
func dijkstraSearch(graph map[string]map[string]float64, costs map[string]float64) float64 {
	// all vertices are marked as unprocessed by default
	var processedVertices []string

	// get first vertex with minimal cost
	vertex := findLowestCostVertex(costs, processedVertices)
	// iterate over graph and try to reduce cost to achieve V(finish)
	for vertex != "" {
		neighbors := graph[vertex]
		currentVertexCost := costs[vertex]

		// vertices from current vertex V(E) are called neighbors
		for neighbor := range neighbors {
			// for each neighbor has being calculated new distance (vertexCost+neighborCost)
			newCost := currentVertexCost + neighbors[neighbor]

			// if new cost is lowest then existing one than replace it
			if costs[neighbor] > newCost {
				costs[neighbor] = newCost
			}
		}

		processedVertices = append(processedVertices, vertex)
		// get one of unprocessed vertices which has minimal cost
		vertex = findLowestCostVertex(costs, processedVertices)
	}

	return costs["finish"]
}

func findLowestCostVertex(costs map[string]float64, processedVertices []string) string {
	var lowestCost = math.Inf(1)
	var lowestCostVertex = ""

	for vertex, cost := range costs {
		// detect the lowest cost from unprocessed vertices
		if cost < lowestCost && !func(slice []string, needle string) bool {
			for _, item := range slice {
				if item == needle {
					return true
				}
			}
			return false
		}(processedVertices, vertex) {
			lowestCost = cost
			lowestCostVertex = vertex
		}
	}

	return lowestCostVertex
}
