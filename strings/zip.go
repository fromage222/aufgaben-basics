package strings

// Erwartet zwei strings s1 und s2.
// Liefert einen String, der abwechselnd aus den Buchstaben des einen und des anderen
// Strings besteht.
func Zip(s1, s2 string) string {
	ergebnis := ""
	min := len(s1)
	if len(s2) < min {
		min = len(s2)
	}
	for i := 0; i < min; i++ {
		ergebnis += string(s1[i]) + string(s2[i])

	}
	ergebnis += s1[min:]
	ergebnis += s2[min:]
	return ergebnis
}
