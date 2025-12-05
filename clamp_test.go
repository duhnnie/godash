package godash

import (
	"testing"
)

func TestClamp(t *testing.T) {
	t.Parallel()
	tests := []struct {
		value    int
		min      int
		max      int
		expected int
	}{
		{5, 1, 10, 5},
		{0, 1, 10, 1},
		{15, 1, 10, 10},
		{10, 10, 20, 10},
		{25, 10, 20, 20},
	}

	for _, test := range tests {
		result := Clamp(test.value, test.min, test.max)
		if result != test.expected {
			t.Errorf("Clamp(%d, %d, %d) = %d; expected %d", test.value, test.min, test.max, result, test.expected)
		}
	}
}
