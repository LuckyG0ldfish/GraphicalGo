package context

import (
	"image/color"
)

var (
	Black = color.Gray16{0}
	White = color.Gray16{0xffff}
	Gray  = color.Alpha16{0b1110111100100010}

	// Transparent = color.Alpha16{0}
	// Opaque      = color.Alpha16{0xffff}
)
