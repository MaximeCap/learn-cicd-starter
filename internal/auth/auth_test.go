package auth

import (
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	t.Run("Should have an error if not auth header", func(t *testing.T) {
		headers := http.Header{}
		headers.Add("content-type", "application/json")

		_, err := GetAPIKey(headers)
		if err == nil {
			t.Error("Expected error but got nothing")
		}
	})

	t.Run("should have only the token part", func(t *testing.T) {
		headers := http.Header{}
		headers.Add("Authorization", "ApiKey test")

		expect := "test"

		if got, _ := GetAPIKey(headers); got != expect {
			t.Errorf("Expect %s, but got %s", expect, got)
		}
	})

}
