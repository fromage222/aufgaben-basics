package rectangles

import (
	"fmt"
	"strings"
)

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Das Rechteck soll komplett mit `#`-Zeichen gefüllt sein.

func DrawSolidRectangle(height, width int) {

	for i := 0; i < height; i++ {
		fmt.Println(strings.Repeat("#", width)) //Package Strings. von nöten
	}
}
