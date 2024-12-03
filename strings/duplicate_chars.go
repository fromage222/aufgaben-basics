package strings

// Erwartet einen String s und liefert einen neuen String,
// in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
func DuplicateChars(s string) string {
	ergebnis := ""
	for i := 0; i < len(s); i++ {
		ergebnis += string(s[i]) + string(s[i])
	}
	return ergebnis
}
