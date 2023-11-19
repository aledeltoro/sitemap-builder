package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aledeltoro/html-link-parser/link"
)

var (
	client = http.Client{
		Timeout: time.Second * 30,
	}
)

// GetPage performs a HTTP GET request to fetch the HTML of the given URL
func GetPageLinks(url string) ([]link.Link, error) {
	res, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code: %d", res.StatusCode)
	}

	return link.Extract(res.Body)
}
