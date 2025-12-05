package godash

import (
	"errors"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("transforms integers to strings", func(t *testing.T) {
		t.Parallel()
		input := []int{1, 2, 3}
		result, err := Map(input, func(element int, index int, collection []int) (string, error) {
			return fmt.Sprintf("%d", element), nil
		})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		expected := []string{"1", "2", "3"}
		if !slicesEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("returns error on transformation failure", func(t *testing.T) {
		t.Parallel()
		input := []int{1, 2, 3}
		result, err := Map(input, func(element int, index int, collection []int) (string, error) {
			if element == 2 {
				return "", errors.New("transformation failed")
			}
			return fmt.Sprintf("%d", element), nil
		})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if len(result) != 3 {
			t.Errorf("expected partial result length 3, got %d", len(result))
		}
	})

	t.Run("handles empty collection", func(t *testing.T) {
		t.Parallel()
		input := []int{}
		result, err := Map(input, func(element int, index int, collection []int) (string, error) {
			return fmt.Sprintf("%d", element), nil
		})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if len(result) != 0 {
			t.Errorf("expected empty result, got %v", result)
		}
	})

	t.Run("passes correct index and collection to function", func(t *testing.T) {
		t.Parallel()
		input := []int{10, 20, 30}
		var indices []int
		Map(input, func(element int, index int, collection []int) (int, error) {
			indices = append(indices, index)
			return element * 2, nil
		})
		if !slicesEqual(indices, []int{0, 1, 2}) {
			t.Errorf("expected indices [0 1 2], got %v", indices)
		}
	})
}

func slicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
