package godash

import (
	"errors"
	"testing"
)

func TestReduceMap_SumValues(t *testing.T) {
	t.Parallel()
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	reducer := func(acc int, key string, value int, dict map[string]int) (int, error) {
		return acc + value, nil
	}
	result, err := ReduceMap(m, reducer, 0)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if result != 6 {
		t.Errorf("expected 6, got %d", result)
	}
}

func TestReduceMap_ConcatenateStrings(t *testing.T) {
	t.Parallel()
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	reducer := func(acc string, key int, value string, dict map[int]string) (string, error) {
		return acc + value, nil
	}
	result, err := ReduceMap(m, reducer, "")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(result) != 3 || !contains(result, "a") || !contains(result, "b") || !contains(result, "c") {
		t.Errorf("expected concatenated string with a, b, c, got %s", result)
	}
}

func TestReduceMap_EmptyMap(t *testing.T) {
	t.Parallel()
	m := map[string]int{}
	reducer := func(acc int, key string, value int, dict map[string]int) (int, error) {
		return acc + value, nil
	}
	result, err := ReduceMap(m, reducer, 42)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

func TestReduceMap_ReducerError(t *testing.T) {
	t.Parallel()
	m := map[string]int{"a": 1, "b": 2}
	reducer := func(acc int, key string, value int, dict map[string]int) (int, error) {
		if key == "b" {
			return acc, errTest
		}
		return acc + value, nil
	}
	_, err := ReduceMap(m, reducer, 0)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestReduceMap_CountElements(t *testing.T) {
	t.Parallel()
	m := map[string]bool{"x": true, "y": false, "z": true}
	reducer := func(acc int, key string, value bool, dict map[string]bool) (int, error) {
		return acc + 1, nil
	}
	result, err := ReduceMap(m, reducer, 0)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if result != 3 {
		t.Errorf("expected 3, got %d", result)
	}
}

func contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == substr {
			return true
		}
	}
	return false
}

var errTest = errors.New("test error")
