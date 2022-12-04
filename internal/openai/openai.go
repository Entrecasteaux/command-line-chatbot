// Copyright 2023 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openai

import (
	"context"
	"os"
	"time"
)

var (
	apiKey = os.Getenv("OPENAI_API_KEY")
)

type CompletionRequest struct {
	Model            string  `json:"model,omitempty"`
	Prompt           string  `json:"prompt,omitempty"`
	Suffix           string  `json:"suffix,omitempty"`
	MaxTokens        int     `json:"max_tokens,omitempty"`
	Temperature      float64 `json:"temperature,omitempty"`
	TopP             float64 `json:"top_p,omitempt