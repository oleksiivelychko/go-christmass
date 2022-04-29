package dijkstra_search

import (
	"github.com/oleksiivelychko/go-helper/in"
	"math"
)

/*
DijkstraSearch

Only for weighted graphs.
Weighted graph - a graph in which each edge is assigned some weight(value) e.g. number.
For negative weights must be used Bellman–Ford algorithm.

1. Each vertex has minimal-known distance (costs) to V(start).
2. V(start) assigned as 0, V(finish) assigned as infinity.
3. All vertices marked as non-processed.
4. Iterate graph and try to reduce cost to make it to V(finish).
5. Get one of non-processed vertex which has minimal cost.
7. Vertices from this V(E) named `neighbors`.
8. For each neighbor considers new distance (nodeCost+neighborCost)
9. If its value less than replace with new distance.
*/
func DijkstraSearch(graph map[string]map[string]float64, costs map[string]float64, parents map[string]string) float64 {
	var processed []string

	node := findLowestCostNode(costs, processed)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]
		for neighborNode := range neighbors {
			newCost := cost + neighbors[neighborNode]
			if costs[neighborNode] > newCost {
				costs[neighborNode] = newCost
				parents[neighborNode] = neighborNode
			}
		}

		processed = append(processed, node)
		node = findLowestCostNode(costs, processed)
	}

	return costs["finish"]
}

func findLowestCostNode(costs map[string]float64, processed []string) string {
	var lowestCost = math.Inf(1)
	var lowestCostNode = ""

	for node := range costs {
		cost := costs[node]
		_, found := in.In(processed, node)
		if cost < lowestCost && !found {
			lowestCost = cost
			lowestCostNode = node
		}
	}

	return lowestCostNode
}
