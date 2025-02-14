package models

import (
	"testing"
	"time"
)

func TestRequest_AddHeader(t *testing.T) {
	req := &Request{}
	req.AddHeader("Content-Type", "application/json")
	req.AddHeader("Authorization", "Bearer token123")

	if len(req.Headers) != 2 {
		t.Errorf("Expected 2 headers, got %d", len(req.Headers))
	}

	if req.Headers[0].Key != "Content-Type" {
		t.Errorf("Expected first header key to be 'Content-Type', got '%s'", req.Headers[0].Key)
	}

	if req.Headers[1].Value != "Bearer token123" {
		t.Errorf("Expected second header value to be 'Bearer token123', got '%s'", req.Headers[1].Value)
	}
}

func TestRequest_ToJSON(t *testing.T) {
	req := &Request{
		ID:        "test-id",
		Name:      "Test Request",
		Method:    GET,
		URL:       "https://example.com",
		Headers:   []Header{{Key: "Accept", Value: "application/json"}},
		Body:      `{"test": true}`,
		CreatedAt: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
	}

	json, err := req.ToJSON()
	if err != nil {
		t.Errorf("ToJSON() returned error: %v", err)
	}

	if json == "" {
		t.Error("ToJSON() returned empty string")
	}

	expectedContent := []string{
		`"id": "test-id"`,
		`"method": "GET"`,
		`"url": "https://example.com"`,
	}

	for _, content := range expectedContent {
		if !contains(json, content) {
			t.Errorf("JSON output missing expected content: %s", content)
		}
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[len(s)-len(substr):] == substr ||
		len(s) > len(substr) && s[:len(substr)] == substr ||
		len(s) > len(substr) && findInString(s, substr)
}

func findInString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
