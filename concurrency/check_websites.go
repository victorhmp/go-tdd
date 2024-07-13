package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	// This way of using the goroutime below will cause unexpected results in Go
	// versions < 1.22.0. The issue would be that the url variable would be reused
	// in all goroutines, causing the last value of the urls slice to be used in
	// all of the goroutines.
	// But since 1.22, this will work as expected.
	for _, url := range urls {
		go func() {
			// This is a send statement. Sending a value to a channel.
			resultsChannel <- result{url, wc(url)}
		}()
	}

	// Now we're controling the timing of reads and writes to the results map,
	// as we're updating it in order, just reading a value from the shared channel.
	for i := 0; i < len(urls); i++ {
		// This is a receive expression. Receiving a value from a channel.
		result := <-resultsChannel
		results[result.string] = result.bool
	}

	return results
}
