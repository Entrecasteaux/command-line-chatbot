
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
	g6 := ((uint16(g) * 5) / 255)
	b6 := ((uint16(b) * 5) / 255)
	i := 36*r6 + 6*g6 + b6
	if foreground {
		return fgcolors[i]
	}
	return bgcolors[i]
}

func grayscale(scale uint8, foreground bool) ([]byte, bool) {
	var source [256][]byte

	if foreground {
		source = fgTermRGB
	} else {
		source = bgTermRGB
	}

	switch scale {
	case 0x08:
		return source[232], true
	case 0x12:
		return source[233], true
	case 0x1c:
		return source[234], true
	case 0x26:
		return source[235], true
	case 0x30:
		return source[236], true
	case 0x3a:
		return source[237], true
	case 0x44:
		return source[238], true
	case 0x4e:
		return source[239], true
	case 0x58:
		return source[240], true
	case 0x62:
		return source[241], true
	case 0x6c:
		return source[242], true
	case 0x76:
		return source[243], true
	case 0x80:
		return source[244], true
	case 0x8a:
		return source[245], true
	case 0x94:
		return source[246], true
	case 0x9e:
		return source[247], true
	case 0xa8:
		return source[248], true
	case 0xb2:
		return source[249], true
	case 0xbc:
		return source[250], true
	case 0xc6:
		return source[251], true
	case 0xd0:
		return source[252], true
	case 0xda:
		return source[253], true
	case 0xe4:
		return source[254], true
	case 0xee:
		return source[255], true
	}
	return nil, false
}

var (
	yellow = fgString("", 252, 140, 3)
	red    = fgString("", 255, 0, 0)
	green  = fgString("", 0, 255, 0)
)

// \033[

var fgTermRGB = [...][]byte{
	[]byte("38;5;0"),
	[]byte("38;5;1"),
	[]byte("38;5;2"),
	[]byte("38;5;3"),
	[]byte("38;5;4"),
	[]byte("38;5;5"),
	[]byte("38;5;6"),
	[]byte("38;5;7"),
	[]byte("38;5;8"),
	[]byte("38;5;9"),
	[]byte("38;5;10"),
	[]byte("38;5;11"),
	[]byte("38;5;12"),
	[]byte("38;5;13"),
	[]byte("38;5;14"),
	[]byte("38;5;15"),
	[]byte("38;5;16"),
	[]byte("38;5;17"),
	[]byte("38;5;18"),
	[]byte("38;5;19"),
	[]byte("38;5;20"),
	[]byte("38;5;21"),
	[]byte("38;5;22"),
	[]byte("38;5;23"),
	[]byte("38;5;24"),
	[]byte("38;5;25"),
	[]byte("38;5;26"),
	[]byte("38;5;27"),
	[]byte("38;5;28"),
	[]byte("38;5;29"),
	[]byte("38;5;30"),
	[]byte("38;5;31"),
	[]byte("38;5;32"),
	[]byte("38;5;33"),
	[]byte("38;5;34"),
	[]byte("38;5;35"),
	[]byte("38;5;36"),
	[]byte("38;5;37"),
	[]byte("38;5;38"),
	[]byte("38;5;39"),
	[]byte("38;5;40"),
	[]byte("38;5;41"),
	[]byte("38;5;42"),
	[]byte("38;5;43"),
	[]byte("38;5;44"),
	[]byte("38;5;45"),