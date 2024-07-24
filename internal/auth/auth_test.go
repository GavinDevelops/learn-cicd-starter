package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey testing")
	_, err := GetAPIKey(header)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestFailGetAPIKey(t *testing.T) {
	_, err := GetAPIKey(http.Header{})
	if err == nil {
		t.Fatal("Expected fail on no Authorization Header")
	}
}

func TestFailImproperAuthHeader(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "testing")
	_, err := GetAPIKey(header)
	if err == nil {
		t.Fatal("Expectiong fail on improper Authorizaion header")
	}
	t.Fatal("Testin ci fail")
}
