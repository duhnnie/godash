package godash

import (
	"errors"
	"testing"
)

func TestReduceSum(t *testing.T) {
	result, err := Reduce([]int{1, 2, 3, 4}, func(acc int, cur int, idx int, slice []int) (int, error) {
		return acc + cur, nil
	}, 0)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 10 {
		t.Errorf("expected 10, got %d", result)
	}
}

func TestReduceEmptySlice(t *testing.T) {
	result, err := Reduce([]int{}, func(acc int, cur int, idx int, slice []int) (int, error) {
		return acc + cur, nil
	}, 42)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

func TestReduceDifferentTypes(t *testing.T) {
	result, err := Reduce([]string{"a", "b", "c"}, func(acc string, cur string, idx int, slice []string) (string, error) {
		return acc + cur, nil
	}, "")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != "abc" {
		t.Errorf("expected 'abc', got %q", result)
	}
}

func TestReduceWithError(t *testing.T) {
	_, err := Reduce([]int{1, 2, 3}, func(acc int, cur int, idx int, slice []int) (int, error) {
		if cur == 2 {
			return 0, errors.New("some error")
		}
		return acc + cur, nil
	}, 0)

	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestReduceWithIndex(t *testing.T) {
	result, err := Reduce([]int{10, 20, 30}, func(acc int, cur int, idx int, slice []int) (int, error) {
		return acc + idx, nil
	}, 0)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 3 {
		t.Errorf("expected 3, got %d", result)
	}
}
