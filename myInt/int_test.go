package myInt

import "testing"

func TestAdder(t *testing.T) {
	assestReturn := func(t *testing.T, sum, result int) {
		t.Helper()
		if sum != result {
			t.Errorf("result: '%d' sum: '%d'", result, sum)
		}
	}
	t.Run("2 + 2 return 4", func(t *testing.T) {
		sum := Add(2, 2)
		result := 4
		assestReturn(t, sum, result)
	})
	t.Run("-1 + 1 return 0", func(t *testing.T) {
		sum := Add(-1, 1)
		result := 0
		assestReturn(t, sum, result)
	})
}
