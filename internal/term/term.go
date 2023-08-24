
// Copyright 2023 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package term

// Red turns a given string to red
func Red(in string) string {
	return fgString(in, 255, 0, 0)
}

// Orange turns a given string to orange
func Orange(in string) string {
	return fgString(in, 252, 140, 3)
}

// Green turns a given string to green
func Green(in string) string {
	return fgString(in, 0, 255, 0)
}

// Gray turns a given string to gray
func Gray(in string) string {
	return fgString(in, 125, 125, 125)
}

var (
	before   = []byte("\033[")
	after    = []byte("m")
	reset    = []byte("\033[0;00m")
	fgcolors = fgTermRGB[16:232]
	bgcolors = bgTermRGB[16:232]
)

func fgString(in string, r, g, b uint8) string {
	return string(fgBytes([]byte(in), r, g, b))
}

// Bytes colorizes the foreground with the terminal color that matches
// the closest the RGB color.
func fgBytes(in []byte, r, g, b uint8) []byte {
	return colorize(color(r, g, b, true), in)
}

func colorize(color, in []byte) []byte {
	return append(append(append(append(before, color...), after...), in...), reset...)
}

func color(r, g, b uint8, foreground bool) []byte {
	// if all colors are equal, it might be in the grayscale range
	if r == g && g == b {
		color, ok := grayscale(r, foreground)
		if ok {
			return color
		}
	}

	// the general case approximates RGB by using the closest color.
	r6 := ((uint16(r) * 5) / 255)