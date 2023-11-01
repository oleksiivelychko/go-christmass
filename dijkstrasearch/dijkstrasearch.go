// Package dijkstrasearch
// For weighted graph only, where weighted graph is a graph in which each edge has assigned the weight (value)
// - each vertex has minimal-known distance (costs) to V(start)
// - V(start) is assigned as 0, V(finish) is assigned as infinity
package dijkstrasearch

import "math"

func search(g map[string]map[string]float64, costs map[string]float64) float64 {
	// all vertices are marked as unprocessed by default
	var processed []string

	// get first vertex with minimal cost
	vertex := findLowestCostVertex(costs, processed)
	// iterate over graph and try to reduce cost to achieve V(finish)
	for vertex != "" {
		neighbors := g[vertex]

		// vertices from current vertex V(E) are called neighbors
		for n := range neighbors {
			// for each neighbor has being calculated new distance (currentVertexCost+neighborCost)
			cost := costs[vertex] + neighbors[n]

			// if a new cost is lowest then existing one than replace it
			if costs[n] > cost {
				costs[n] = cost
			}
		}

		processed = append(processed, vertex)
		// get one of unprocessed vertices which has minimal cost
		vertex = findLowestCostVertex(costs, processed)
	}

	return costs["finish"]
}

func findLowestCostVertex(costs map[string]float64, processed []string) string {
	var lowest = math.Inf(1)
	var lowestVertex = ""

	for vertex, cost := range costs {
		// detect the lowest cost from unprocessed vertices
		if cost < lowest && !func(p []string, v string) bool {
			for _, i := range p {
				if i == v {
					return true
				}
			}
			return false
		}(processed, vertex) {
			lowest = cost
			lowestVertex = vertex
		}
	}

	return lowestVertex
}
