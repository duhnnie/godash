package godash

import (
	"fmt"
	"testing"
)

func TestFilter_IntSlice_EvenNumbers(t *testing.T) {
	t.Parallel()
	input := []int{1, 2, 3, 4, 5, 6}
	result, err := Filter(input, func(item int) (bool, error) {
		return item%2 == 0, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 3 || result[0] != 2 || result[1] != 4 || result[2] != 6 {
		t.Errorf("expected [2 4 6], got %v", result)
	}
}

func TestFilter_StringSlice_LengthGreaterThanThree(t *testing.T) {
	t.Parallel()
	input := []string{"go", "a", "ok", "dashboard", "test"}
	result, err := Filter(input, func(item string) (bool, error) {
		return len(item) > 3, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 2 || result[0] != "dashboard" || result[1] != "test" {
		t.Errorf("expected [dashboard test], got %v", result)
	}
}

func TestFilter_EmptySlice(t *testing.T) {
	t.Parallel()
	input := []int{}
	result, err := Filter(input, func(item int) (bool, error) {
		return item > 0, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}

func TestFilter_NoMatches(t *testing.T) {
	t.Parallel()
	input := []int{1, 3, 5, 7}
	result, err := Filter(input, func(item int) (bool, error) {
		return item%2 == 0, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}

func TestFilter_ErrorFromIterator(t *testing.T) {
	t.Parallel()
	input := []int{1, 2, 3}
	testErr := fmt.Errorf("filter error")

	_, err := Filter(input, func(item int) (bool, error) {
		if item == 2 {
			return false, testErr
		}
		return true, nil
	})

	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err != testErr {
		t.Errorf("expected error %v, got %v", testErr, err)
	}
}

func TestFilter_AllMatch(t *testing.T) {
	t.Parallel()
	input := []int{2, 4, 6, 8}
	result, err := Filter(input, func(item int) (bool, error) {
		return item%2 == 0, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 4 {
		t.Errorf("expected 4 elements, got %d", len(result))
	}
}
