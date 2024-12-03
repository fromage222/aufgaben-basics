package strings

// Erwartet zwei Strings s und sep sowie eine Zahl n.
// Liefert einen String, der aus n Kopien von s besteht, die duch sep getrennt werden.
func ConcatN(s, sep string, n int) string {
	ergebnis := ""

	if sep == "" {
		for j := 0; j < n-1; j++ {
			ergebnis += s
		}
	}
	if sep != "" {
		for i := 0; i < n-1; i++ {
			ergebnis += s + sep

		}

	}
	ergebnis += s
	return ergebnis
}
