package subsequences_search

import (
	"github.com/oleksiivelychko/go-helper/matrix"
	"math"
)

/*
SubsequencesSearch

O(string1length*string2length)

Dynamic programming is a way to solve complex tasks by breaking them down into simpler subtasks.
*/
func SubsequencesSearch(word1, word2 string) int {
	Matrix := matrix.IntMatrix(len(word1), len(word2))

	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				if i >= 1 && j >= 1 {
					Matrix[i][j] = Matrix[i-1][j-1] + 1
				} else {
					Matrix[i][j] = 1
				}
			} else {
				if i >= 1 && j >= 1 {
					Matrix[i][j] = int(math.Max(float64(Matrix[i-1][j]), float64(Matrix[i][j-1])))
				} else {
					Matrix[i][j] = 0
				}
			}
		}
	}

	return Matrix[len(word1)-1][len(word2)-1]
}
