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
// That's what the net/http/httptest package is for! It still enables us to
// fire HTTP requests, but we can send them to test servers running on
// localhost, eliminating the internet from our test runs.

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	// Function calls prefixed with `defer` will be executed at the end of their
	// containing function. This is useful for readability, since we can move
	// the `.Close()` calls next to when we're creating the servers.
	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	expected := fastURL
	got := Racer(slowURL, fastURL)

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
