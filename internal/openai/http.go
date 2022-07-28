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
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s",