package service

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/aledeltoro/html-link-parser/link"
	"github.com/aledeltoro/sitemap-builder/internal/models"
)

func BuildSitemap(domain string) *models.Sitemap {
	sitemap := models.NewSitemap()

	queue := models.NewQueue()
	queue.Enqueue(domain)

	visited := map[string]bool{}

	for queue.Size() != 0 {
		domain := queue.DeQueue()
		visited[domain] = true

		url, err := url.Parse(domain)
		if err != nil {
			log.Fatalln(err)
		}

		req, err := http.NewRequest(http.MethodGet, url.String(), nil)
		if err != nil {
			log.Fatalln(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			log.Fatalf("invalid status code: %d", res.StatusCode)
		}

		sitemap.URL = append(sitemap.URL, models.URL{Locator: domain})

		links, err := link.Extract(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		hostname := getDomain(url.Host)

		for _, link := range links {
			if !strings.Contains(link.Href, hostname) {
				fmt.Printf("URL outside of domain or relative url: %s", link.Href)

				continue
			}

			_, ok := visited[link.Href]
			if !ok {
				queue.Enqueue(link.Href)
				continue
			}

			fmt.Printf("Already visited: %s \n", link.Href)
		}

	}

	return sitemap
}

func getDomain(url string) string {
	return strings.TrimPrefix(url, "www.")
}
