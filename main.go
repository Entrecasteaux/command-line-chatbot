
// Copyright 2023 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"changkun.de/x/chat/internal/openai"
	"changkun.de/x/chat/internal/term"
)

func main() {
	if os.Getenv("OPENAI_API_KEY") == "" {
		fmt.Fprint(os.Stderr, "Please set OPENAI_API_KEY environment variable.\n")
		return
	}

	stdin := os.Stdin
	stdout := os.Stdout
	fmt.Fprint(stdout, term.Orange("Hi, I'm a chatbot. How can I help you?\n"))

	session := []openai.ChatMessage{
		{
			Role:    "system",
			Content: "You are a helpful assistant.",
		},
	}

	for {