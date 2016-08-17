package strings

import "testing"

func TestHasOnlyUnique(t *testing.T) {
	cases := map[string]bool{
		"abcd": true,
	}

	for given, expected := range cases {
		got := HasOnlyUnique(given)

		if expected != got {
			t.Errorf("wrong result, expected %t but got %t", expected, got)
		}
	}
}
