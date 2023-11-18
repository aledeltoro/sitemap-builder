package client

import (
	"fmt"
	"net/http"
	"time"
)

var (
	client = http.Client{
		Timeout: time.Second * 30,
	}
)

// GetPage performs a HTTP GET request to fetch the HTML of the given URL
func GetPage(url string) (*http.Response, error) {
	res, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code: %d", res.StatusCode)
	}

	return res, nil
}
