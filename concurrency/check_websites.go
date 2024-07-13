package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	// This way of using the goroutime below will cause unexpected results in Go
	// versions < 1.22.0. The issue would be that the url variable would be reused
	// in all goroutines, causing the last value of the urls slice to be used in
	// all of the goroutines.
	// But since 1.22, this will work as expected.
	for _, url := range urls {
		go func() {
			results[url] = wc(url)
		}()
	}

	return results
}
