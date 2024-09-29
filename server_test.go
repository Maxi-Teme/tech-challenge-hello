package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETHello(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	res := httptest.NewRecorder()

	helloHandler := HelloHandler{env: "test"}

	helloHandler.ServeHTTP(res, req)

	t.Run("returns status 200", func(t *testing.T) {
		result := res.Code
		expected := 200

		if result != expected {
			t.Errorf("got %q, expected %q", result, expected)
		}
	})
}
