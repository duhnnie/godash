package godash

import (
	"errors"
	"testing"
)

func TestAnyTrue(t *testing.T) {
	result, err := Any([]int{1, 2, 3, 4}, func(item int) (bool, error) {
		return item > 2, nil
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !result {
		t.Errorf("expected true, got false")
	}
}

func TestAnyFalse(t *testing.T) {
	result, err := Any([]int{1, 2, 3, 4}, func(item int) (bool, error) {
		return item > 10, nil
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result {
		t.Errorf("expected false, got true")
	}
}

func TestAnyEmptySlice(t *testing.T) {
	result, err := Any([]int{}, func(item int) (bool, error) {
		return true, nil
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result {
		t.Errorf("expected false, got true")
	}
}

func TestAnyWithError(t *testing.T) {
	_, err := Any([]int{1, 2, 3}, func(item int) (bool, error) {
		if item == 2 {
			return false, errors.New("some error")
		}
		return item > 5, nil
	})

	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestAnyDifferentTypes(t *testing.T) {
	result, err := Any([]string{"foo", "bar", "baz"}, func(item string) (bool, error) {
		return item == "bar", nil
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !result {
		t.Errorf("expected true, got false")
	}
}

func TestAnyFirstElement(t *testing.T) {
	result, err := Any([]int{5, 1, 2, 3}, func(item int) (bool, error) {
		return item > 4, nil
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !result {
		t.Errorf("expected true, got false")
	}
}
