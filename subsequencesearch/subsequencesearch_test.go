package subsequencesearch

import "testing"

/*
| m | i | s | h |			| f | o | s | h |
___________________		  ___________________
f | 0 | 0 | 0 | 0 |       f | 1 | 0 | 0 | 0 |
__|___|___|___|___|       __|___|___|___|___|
i | 0 | 1 | 1 | 1 |       i | 0 | 0 | 0 | 0 |
__|___|___|___|___|       __|___|___|___|___|
s | 0 | 1 | 2 | 2 |       s | 0 | 0 | 1 | 1 |
__|___|___|___|___|       __|___|___|___|___|
h | 0 | 1 | 2 | 3 |       h | 0 | 0 | 1 | 2 |
__|___|___|___|___|       __|___|___|___|___|
*/
func TestSubsequencesSearch(t *testing.T) {
	c := search("fish", "mish")
	if c != 3 {
		t.Errorf("mismatch count %d", c)
	}

	c = search("fish", "fosh")
	if c != 2 {
		t.Errorf("mismatch count %d", c)
	}
}

func BenchmarkSubsequencesSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		search("fish", "aosh")
	}
}
