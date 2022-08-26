// Copyright 2023 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openai

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func httpRequest[Req, Resp any](ctx context.Context, url string, in *Req) (*Resp, error) {
	b, _ := json.Marshal(in)

	r, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("failed to create the request: %w", err)
	}
	r.Header.Set("Content-Type", "application/json; charset=UTF-8")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	resp, err := http.DefaultClient.Do(r.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to do the request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response code %v: %v", resp.StatusCode, string(data))
	}

	var out Resp
	err = json.Unmarshal(data, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to parse edit response: %w", err)