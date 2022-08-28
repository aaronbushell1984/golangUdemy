package selectexample

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const testTimeoutSeconds = 1

func TestRacerSequential(t *testing.T) {

	slowServer := makeDelayedServer(20 * time.Millisecond)
	defer slowServer.Close()
	fastServer := makeDelayedServer(0 * time.Millisecond)
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := RacerSequential(slowUrl, fastUrl)

	assertUrl(t, got, want)
}

func TestRacerSelect(t *testing.T) {
	t.Run("test fastest", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertUrl(t, got, want)
	})
	t.Run("returns error if server takes over defined time", func(t *testing.T) {
		server1TooSlow := makeDelayedServer(testTimeoutSeconds * time.Second)
		defer server1TooSlow.Close()
		server2TooSlow := makeDelayedServer(testTimeoutSeconds * time.Second)
		defer server2TooSlow.Close()

		_, err := ConfigurableRacerSelect(server1TooSlow.URL, server2TooSlow.URL, testTimeoutSeconds)
		if err == nil {
			t.Fatalf("exepected error and did not get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return slowServer
}

func assertUrl(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q want: %q", got, want)
	}
}
