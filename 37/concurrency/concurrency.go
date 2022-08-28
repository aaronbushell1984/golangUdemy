// Package concurrency demonstrates/simulates potentially improving speed of multiple sequential web calls
//
// Adding a concurrency pattern using goroutines and channels allows the program to run in parallel if hardware and conditions allow
package concurrency

// WebsiteChecker simulates (in this example a function which would check the status of a website
type WebsiteChecker func(string) bool

// result struct defines the type to hold website url string and if the site was accessible when received from a channel
type result struct {
	website      string
	isAccessible bool
}

// CheckWebsites was refactored during this exercise to create a goroutine for each WebsiteChecker call on url
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultsChannel <- result{u, wc(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultsChannel
		results[r.website] = r.isAccessible
	}
	return results
}
