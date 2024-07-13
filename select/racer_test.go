package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Just directly testing a function that makes HTTP calls is not a good idea.
// It would be mixing the logic we actually want to test with the network
// conditions, which can be flaky. Also, could be pretty slow.
// That's what the net/http/httptest package is for!

func TestRacer(t *testing.T) {

	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	expected := fastURL
	got := Racer(slowURL, fastURL)

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}

	slowServer.Close()
	fastServer.Close()
}
