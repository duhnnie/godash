package godash

import (
	"fmt"
	"testing"
)

func TestReduceMapTo_IntMapSum(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 3}
	result, err := ReduceMapTo(input, func(acc int, key string, value int) (int, error) {
		return acc + value, nil
	}, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 6 {
		t.Errorf("expected 6, got %d", result)
	}
}

func TestReduceMapTo_StringConcatenation(t *testing.T) {
	input := map[int]string{1: "a", 2: "b", 3: "c"}
	result, err := ReduceMapTo(input, func(acc string, key int, value string) (string, error) {
		return acc + value, nil
	}, "")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "abc" && result != "acb" && result != "bac" && result != "bca" && result != "cab" && result != "cba" {
		t.Errorf("expected concatenation of a, b, c, got %s", result)
	}
}

func TestReduceMapTo_EmptyMap(t *testing.T) {
	input := map[string]int{}
	result, err := ReduceMapTo(input, func(acc int, key string, value int) (int, error) {
		return acc + value, nil
	}, 10)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 10 {
		t.Errorf("expected 10, got %d", result)
	}
}

func TestReduceMapTo_ErrorFromReducer(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2}
	testErr := fmt.Errorf("reducer error")

	_, err := ReduceMapTo(input, func(acc int, key string, value int) (int, error) {
		if key == "b" {
			return 0, testErr
		}
		return acc + value, nil
	}, 0)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err != testErr {
		t.Errorf("expected error %v, got %v", testErr, err)
	}
}

func TestReduceMapTo_ComplexType(t *testing.T) {
	input := map[string]int{"x": 5, "y": 3, "z": 2}
	result, err := ReduceMapTo(input, func(acc int, key string, value int) (int, error) {
		if value > acc {
			return value, nil
		}
		return acc, nil
	}, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}
