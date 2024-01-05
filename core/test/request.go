package test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type RequestArgs struct {
	Method      string
	Path        string
	Body        []byte
	JWTToken    string
	ContentType string
}

func Request(t *testing.T, ts *httptest.Server, args *RequestArgs) (*http.Response, string) {
	ctx := context.Background()
	body := bytes.NewReader(args.Body)
	req, err := http.NewRequestWithContext(ctx, args.Method, ts.URL+args.Path, body)
	require.NoError(t, err)

	if args.JWTToken != "" {
		req.Header.Set("Authorization", "Bearer "+args.JWTToken)
	}

	if args.ContentType != "" {
		req.Header.Set("Content-Type", args.ContentType)
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}
