
// Copyright 2023 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openai_test

import (
	"context"
	"os"
	"strings"
	"testing"
	"unicode/utf8"

	"changkun.de/x/chat/internal/openai"
)

func countWords(s string) int {
	counter := 0
	for _, word := range strings.Fields(s) {
		runeCount := utf8.RuneCountInString(word)
		if len(word) == runeCount {
			counter++
		} else {
			counter += runeCount
		}
	}

	return counter
}

func TestEdit(t *testing.T) {