package main

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StatusResult struct {
	Message string `json:"message"`
}

func TestStatus(t *testing.T) {
	req, _ := http.Get("http://localhost:8080/status")
	defer req.Body.Close()

	reqBytes, _ := io.ReadAll(req.Body)

	var status StatusResult

	err := json.Unmarshal(reqBytes, &status)
	if err != nil {
		panic("Failed to deserialize json")
	}

	assert.Equal(t, status.Message, "ok")
}
