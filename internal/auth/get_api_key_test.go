package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoHeader(t *testing.T) {
	header := make(http.Header)

	header.Set("Authorization", "")
	fmt.Printf("Testing blank authorization header.\n")
	result, err := GetAPIKey(header)
	if err == nil {
		t.Errorf("GetApiKey should return empty string, instead returns %v", result)
	}
}

func TestMalformedHeader(t *testing.T) {
	header := make(http.Header)

	header.Set("Authorization", "123456")
	fmt.Printf("Testing Authoriztion:123456.\n")
	result, err := GetAPIKey(header)
	if err == nil {
		t.Errorf("GetApiKey should return empty string, instead returns %v", result)
	}
}