package strings

// Erwartet einen string s und zählt, wie oft der Buchstabe 'A' in s vorkommt.
func CountA(s string) int {
	ergebnis := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			ergebnis++
		}
	}
	return ergebnis
}

// Erwartet einen string s und einen Buchstaben c.
// Zählt, wie oft c in s vorkommt.
func CountChar(s string, c rune) int {
	ergebnis := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			ergebnis++
		}
	}
	return ergebnis
}
