// Package selectexample demonstrates using a select statement to receive values from a channel
package selectexample

import (
	"fmt"
	"net/http"
	"time"
)

// timeoutSeconds allows an error to be thrown if no request is received in specified time
const timeoutSeconds = 10

// RacerSequential compares response time of two requests using responseTime
//
// The second request waits for the first response -> compare with ConfigurableRacerSelect
func RacerSequential(url1, url2 string) (winner string) {
	if responseTime(url1) < responseTime(url2) {
		return url1
	}
	return url2
}

// Racer allows users to call a function and provide only urls and abstract away the timeout setting
//
// This in turn calls ConfigurableRacerSelect which handles the timeout constant set. This also allows a faster timeout setting to be used in testing
func Racer(url1, url2 string) (winner string, err error) {
	return ConfigurableRacerSelect(url1, url2, timeoutSeconds)
}

// ConfigurableRacerSelect waits for the result from ping() in a select case statement
//
// The first channel to receive the result will return and thus be the winner:
//
//	case <-ping(url1):
//		return url1, nil
//
// Case time.After prevents the channel blocking forever and allows a reasonable timeout to be set to return an error:
//
//	case <-time.After(timeout * time.Second):
//		return "", fmt.Errorf("timed out after %v seconds waiting for %s and %s", timeoutSeconds, url1, url2)
func ConfigurableRacerSelect(url1, url2 string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout * time.Second):
		return "", fmt.Errorf("timed out after %v seconds waiting for %s and %s", timeoutSeconds, url1, url2)
	}
}

// ping fans out a channel for each url. The channel is closed and returned on completion
func ping(url string) <-chan struct{} {
	c := make(chan struct{})
	go func() {
		http.Get(url)
		close(c)
	}()
	return c
}

// responseTime records a start time, places a get request to url and returns the time between start and end of the request
func responseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
