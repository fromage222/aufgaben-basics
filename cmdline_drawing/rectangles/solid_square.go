package rectangles

import (
	"fmt"
	"strings"
)

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein Quadrat mit dieser Seitenlänge auf der Konsole.
// Das Quadrat soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidSquare(length int) {

	for i := 0; i < length; i++ {
		fmt.Println(strings.Repeat("#", length)) //Package Strings. von nöten
	}
}
