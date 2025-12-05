package godash

import (
	"fmt"
	"testing"
)

func TestSome_IntSlice_HasEvenNumbers(t *testing.T) {
	t.Parallel()
	input := []int{1, 3, 5, 4, 7}
	result, err := Some(input, func(item int, index int, slice []int) (bool, error) {
		return item%2 == 0, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result {
		t.Errorf("expected true, got false")
	}
}

func TestSome_IntSlice_NoEvenNumbers(t *testing.T) {
	t.Parallel()
	input := []int{1, 3, 5, 7, 9}
	result, err := Some(input, func(item int, index int, slice []int) (bool, error) {
		return item%2 == 0, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result {
		t.Errorf("expected false, got true")
	}
}

func TestSome_StringSlice_LengthGreaterThanThree(t *testing.T) {
	t.Parallel()
	input := []string{"go", "a", "ok", "dashboard"}
	result, err := Some(input, func(item string, index int, slice []string) (bool, error) {
		return len(item) > 3, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result {
		t.Errorf("expected true, got false")
	}
}

func TestSome_EmptySlice(t *testing.T) {
	t.Parallel()
	input := []int{}
	result, err := Some(input, func(item int, index int, slice []int) (bool, error) {
		return item > 0, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result {
		t.Errorf("expected false, got true")
	}
}

func TestSome_ErrorFromIterator(t *testing.T) {
	t.Parallel()
	input := []int{1, 2, 3}
	testErr := fmt.Errorf("iterator error")

	_, err := Some(input, func(item int, index int, slice []int) (bool, error) {
		if item == 2 {
			return false, testErr
		}
		return false, nil
	})

	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err != testErr {
		t.Errorf("expected error %v, got %v", testErr, err)
	}
}

func TestSome_EarlyTermination(t *testing.T) {
	t.Parallel()
	input := []int{1, 2, 3, 4, 5}
	callCount := 0

	result, err := Some(input, func(item int, index int, slice []int) (bool, error) {
		callCount++
		return item == 2, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result {
		t.Errorf("expected true, got false")
	}
	if callCount != 2 {
		t.Errorf("expected 2 calls, got %d", callCount)
	}
}
