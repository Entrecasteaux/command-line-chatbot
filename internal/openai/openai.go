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
	TopP             float64 `json:"top_p,omitempty"`
	N                int     `json:"n,omitempty"`
	Stream           bool    `json:"stream,omitempty"`
	LogProbs         int     `json:"logprobs,omitempty"`
	Echo             bool    `json:"echo,omitempty"`
	Stop             string  `json:"stop,omitempty"`
	PresencePenalty  float64 `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`
	BestOf           int     `json:"best_of,omitempty"`
}

type CompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`