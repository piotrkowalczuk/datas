package strings

// HasOnlyUnique returns true if string has only unique characters.
func HasOnlyUnique(s string) bool {
	chars := map[rune]struct{}{}

	for _, c := range s {
		if _, ok := chars[c]; ok {
			return false
		}

		chars[c] = struct{}{}
	}
	return true
}
