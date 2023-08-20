
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