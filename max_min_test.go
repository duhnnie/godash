package godash

import (
	"testing"
)

func TestMaxInt(t *testing.T) {
	t.Parallel()
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 2},
		{-1, -2, -1},
		{0, 0, 0},
	}

	for _, test := range tests {
		result := Max(test.a, test.b)
		if result != test.expected {
			t.Errorf("Max(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestMaxFloat64(t *testing.T) {
	t.Parallel()
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{3.5, 2.5, 3.5},
		{-1.1, -2.2, -1.1},
		{0.0, 0.0, 0.0},
	}

	for _, test := range tests {
		result := Max(test.a, test.b)
		if result != test.expected {
			t.Errorf("Max(%f, %f) = %f; expected %f", test.a, test.b, result, test.expected)
		}
	}
}

func TestMaxUint(t *testing.T) {
	t.Parallel()
	tests := []struct {
		a, b     uint
		expected uint
	}{
		{5, 10, 10},
		{0, 0, 0},
	}

	for _, test := range tests {
		result := Max(test.a, test.b)
		if result != test.expected {
			t.Errorf("Max(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestMinInt(t *testing.T) {
	t.Parallel()
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 1},
		{-1, -2, -2},
		{0, 0, 0},
	}

	for _, test := range tests {
		result := Min(test.a, test.b)
		if result != test.expected {
			t.Errorf("Min(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestMinFloat64(t *testing.T) {
	t.Parallel()
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{3.5, 2.5, 2.5},
		{-1.1, -2.2, -2.2},
		{0.0, 0.0, 0.0},
	}

	for _, test := range tests {
		result := Min(test.a, test.b)
		if result != test.expected {
			t.Errorf("Min(%f, %f) = %f; expected %f", test.a, test.b, result, test.expected)
		}
	}
}

func TestMinUint(t *testing.T) {
	t.Parallel()
	tests := []struct {
		a, b     uint
		expected uint
	}{
		{5, 10, 5},
		{0, 0, 0},
	}

	for _, test := range tests {
		result := Min(test.a, test.b)
		if result != test.expected {
			t.Errorf("Min(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}
