package subsequence_search

import (
	"github.com/oleksiivelychko/go-code-helpers/structure"
	"math"
)

/*
SubsequenceSearch O(len(Word1)*len(Word2))
Dynamic programming is a way to solve complex tasks by breaking them down into simpler subtasks.
*/
func SubsequenceSearch(word1, word2 string) int {
	matrix := structure.Matrix(len(word1), len(word2))

	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				if i >= 1 && j >= 1 {
					matrix[i][j] = matrix[i-1][j-1] + 1
				} else {
					matrix[i][j] = 1
				}
			} else {
				if i >= 1 && j >= 1 {
					matrix[i][j] = int(math.Max(float64(matrix[i-1][j]), float64(matrix[i][j-1])))
				} else {
					matrix[i][j] = 0
				}
			}
		}
	}

	return matrix[len(word1)-1][len(word2)-1]
}
