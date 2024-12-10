package godash

import (
	"fmt"
	"testing"
)

type person struct {
	FirstName string
	LastName  string
}

func TestMap(t *testing.T) {
	persons := []person{
		{"Kurt", "Cobain"},
		{"Dave", "Grohl"},
		{"Krist", "Novoselic"},
	}

	expected := []string{"Kurt Cobain", "Dave Grohl", "Krist Novoselic"}

	var mapFunc MapFn[person, string] = func(element person, index int, collection []person) string {
		return fmt.Sprintf("%s %s", element.FirstName, element.LastName)
	}

	for index, fullname := range Map(persons, mapFunc) {
		expectedFullName := expected[index]

		if expectedFullName != fullname {
			t.Fatalf("unexpected output for Map(): %s | expecting: %s ", fullname, expectedFullName)
		}
	}

}
