package service

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var urlPathToPage = map[string]string{
	"/":           "samples/index.html",
	"/other-page": "samples/second.html",
	"/dog":        "samples/third.html",
	"/new-page":   "samples/fourth.html",
	"/dog-cat":    "samples/fifth.html",
}

func TestBuildSitemap(t *testing.T) {
	c := require.New(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pagePath, ok := urlPathToPage[r.URL.Path]
		if !ok {
			msg := fmt.Sprintf("unexpected url path: %s", r.URL.Path)
			panic(msg)
		}

		fmt.Fprint(w, loadPage(pagePath))
	}))

	defer server.Close()

	sitemap := BuildSitemap(server.URL)
	c.Len(sitemap.URL, 5)
}

func loadPage(pagePath string) string {
	file, err := os.Open(pagePath)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return string(data)
}
