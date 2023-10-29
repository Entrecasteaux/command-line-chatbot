
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
		fmt.Fprint(stdout, term.Orange("User: "))
		buf := bytes.NewBuffer(nil)
		_, err := io.Copy(buf, stdin)
		if err != nil {
			fmt.Fprintf(stdout, "Error: %v", err)
			return
		}

		userMsg := openai.ChatMessage{
			Role:    "user",
			Content: buf.String(),
		}
		session = append(session, userMsg)

		respCh, errCh := openai.Chat(context.Background(), &openai.ChatRequest{
			Model:   "gpt-4",
			Stream:  true,
			Message: session,
		})

		response := openai.ChatMessage{
			Role:    "assistant",
			Content: "",
		}
		fmt.Fprint(stdout, term.Orange("Assistant: "))

	streamLoop:
		for {
			select {
			case r, ok := <-respCh: