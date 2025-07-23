package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     map[string]string
		expectedKey string
		expectError bool
	}{
		{
			name:        "valid API key",
			headers:     map[string]string{"Authorization": "ApiKey my-secret-key"},
			expectedKey: "my-secret-key",
			expectError: false,
		},
		{
			name:        "missing Authorization header",
			headers:     map[string]string{},
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "malformed Authorization header",
			headers:     map[string]string{"Authorization": "ApiKey"},
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "incorrect scheme in Authorization header",
			headers:     map[string]string{"Authorization": "Bearer my-secret-key"},
			expectedKey: "",
			expectError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := make(http.Header)
			for k, v := range tt.headers {
				headers.Set(k, v)
			}

			key, err := GetAPIKey(headers)

			if (err != nil) != tt.expectError {
				t.Errorf("GetAPIKey() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if key != tt.expectedKey {
				t.Errorf("GetAPIKey() = %v, expected %v", key, tt.expectedKey)
			}
		})
	}
}
