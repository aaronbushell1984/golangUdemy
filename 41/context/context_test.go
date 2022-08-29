package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const configuredSleep = 100

type SpyStore struct {
	response string
	t        *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(configuredSleep * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (s *SpyStore) Cancel() {
}

func TestServer(t *testing.T) {
	t.Run("check response body", func(t *testing.T) {
		data := "Hello World"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		assertResponseBody(t, response, data)
	})
	t.Run("tells store to cancel if request is cancelled", func(t *testing.T) {
		data := "Hello World"
		store := &SpyStore{response: data}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := &SpyResponseWriter{}
		svr.ServeHTTP(response, request)

		assertResponseNotWritten(t, response)
	})

}

func assertResponseNotWritten(t *testing.T, response *SpyResponseWriter) {
	if response.written {
		t.Error("a response should not have been written")
	}
}

func assertResponseBody(t testing.TB, response *httptest.ResponseRecorder, data string) {
	t.Helper()
	if response.Body.String() != data {
		t.Errorf(`got: "%s" want "%s"`, response.Body.String(), data)
	}
}
