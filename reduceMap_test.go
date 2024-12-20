package godash

import (
	"testing"
)

func TestReduceMap(t *testing.T) {
	ageMap := map[string]int{
		"Kurt Cobain":     27,
		"Dave Grohl":      25,
		"Krist Novoselic": 24,
	}

	expected := 76

	var reducerFn ReducerMapFn[string, int, int] = func(acc int, _ string, value int, m map[string]int) int {
		return acc + value
	}

	reduced := ReduceMap(ageMap, reducerFn, 0)

	if reduced != expected {
		t.Fatalf("unexpected output for Reduce(): %d | expecting: %d", reduced, expected)
	}
}
