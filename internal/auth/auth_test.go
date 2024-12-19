package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type GetKeyTests struct {
		KeyHeader string
		Expected  string
	}

	tests := []GetKeyTests{
		{KeyHeader: "ApiKey Baloney", Expected: "Baloney"},
		{KeyHeader: "Bearer Scoop", Expected: ""},
		{KeyHeader: "Api Rub-A-Dub", Expected: ""},
		{KeyHeader: "ApiKey 1234ABCd", Expected: "1234ABCD"},
		{KeyHeader: "", Expected: ""},
	}

	for _, tt := range tests {
		req, _ := http.NewRequest("GET", "https://tony.tone", nil)
		if tt.KeyHeader != "" {
			req.Header.Add("Authorization", tt.KeyHeader)
		}
		key, _ := GetAPIKey(req.Header)
		if key != tt.Expected {
			t.Fatalf("Oh no! Header=%s shouuld have resulted in key=%s, but got=%s",
				tt.KeyHeader, tt.Expected, key)
		}
	}
}
