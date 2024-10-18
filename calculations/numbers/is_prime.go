package numbers

// Erwartet eine Zahl n und prüft, ob n eine Primzahl ist.
func IsPrime(n int) bool {
	// TODO
	teiler := 0

	for i := n; i > 0; i-- {
		rest := n % i
		if rest == 0 {
			teiler++
		}
	}

	if teiler == 2 {
		return true
	}

	return false

}
