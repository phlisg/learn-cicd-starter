package auth

import (
	"net/http/httptest"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Create a new HTTP request with the API key in the header
    req := httptest.NewRequest("GET", "/", nil)
    req.Header.Set("Authorization", "ApiKey test-key")

    // Call the function
    apiKey, err := GetAPIKey(req.Header)

    // Assert the result
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    if apiKey != "test-key" {
        t.Errorf("expected apiKey to be 'test-key', got '%s'", apiKey)
    }
}

func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	// Create a new HTTP request without the Authorization header
	req := httptest.NewRequest("GET", "/", nil)	
	// Call the function
	apiKey, err := GetAPIKey(req.Header)
	
	// Assert the result
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
	}
	if apiKey != "" {
		t.Errorf("expected apiKey to be empty, got '%s'", apiKey)
	}
}
