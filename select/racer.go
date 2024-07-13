// You have been asked to make a function called WebsiteRacer which takes two
// URLs and "races" them by hitting them with an HTTP GET and returning the URL
// which returned first. If none of them return within 10 seconds then it
// should return an error.

package racer

import (
	"net/http"
	"time"
)

func Racer(url1, url2 string) (winner string) {
	startUrl1 := time.Now()
	http.Get(url1)
	url1Duration := time.Since(startUrl1)

	startUrl2 := time.Now()
	http.Get(url2)
	url2Duration := time.Since(startUrl2)

	if url1Duration < url2Duration {
		return url1
	}

	return url2
}
