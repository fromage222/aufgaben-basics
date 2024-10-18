package numbers

// Erwartet eine Zahl n >= 1 und liefert die Anzahl der Teiler dieser Zahl zurück.
func CountDivisors(n int) int {

	teiler := 0

	for i := n; i > 0; i-- {
		rest := n % i
		if rest == 0 {
			teiler++
		}
	}

	return teiler
}
