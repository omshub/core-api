package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndex(t *testing.T) {
	cleanup := testServer(t)
	defer cleanup()

	resp, err := http.Get("http://localhost:1927")
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, 200, resp.StatusCode)
}

func TestPing(t *testing.T) {
	cleanup := testServer(t)
	defer cleanup()

	resp, err := http.Get("http://localhost:1927/ping")
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, 200, resp.StatusCode)

	var response struct {
		Message string `json:"message"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(t, err)
	require.Equal(t, "pong", response.Message)
}

func testServer(t *testing.T) func() {
	server := NewServer(Config{
		Port: "1927",
	}, Dependencies{})

	go func() {
		_ = server.Serve()
	}()

	return func() {
		require.NoError(t, server.Shutdown())
	}
}
