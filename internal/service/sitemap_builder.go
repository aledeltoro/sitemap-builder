package service

import (
	"log"
	"net/url"
	"strings"

	"github.com/aledeltoro/html-link-parser/link"
	"github.com/aledeltoro/sitemap-builder/internal/client"
	"github.com/aledeltoro/sitemap-builder/internal/models"
)

// BuildSitemap searches a domain to build a sitemap according
// to the XML schema of the Sitemap protocal
func BuildSitemap(domain string) *models.Sitemap {
	sitemap := models.NewSitemap()

	queue := models.NewQueue()
	queue.Enqueue(domain)

	visited := map[string]bool{}

	for queue.Size() != 0 {
		domain := queue.DeQueue()

		parsedDomain, err := url.Parse(domain)
		if err != nil {
			log.Fatalln(err)
		}

		response, err := client.GetPage(parsedDomain.String())
		if err != nil {
			log.Fatalln(err)
		}

		defer func() {
			_ = response.Body.Close()
		}()

		visited[domain] = true
		sitemap.AddURL(parsedDomain.String())

		links, err := link.Extract(response.Body)
		if err != nil {
			log.Fatalln(err)
		}

		for _, link := range links {
			parsedLink, err := parsedDomain.Parse(link.Href)
			if err != nil {
				log.Println("error parsing link: ", err)
				continue
			}

			if !strings.Contains(parsedLink.Host, parsedDomain.Hostname()) {
				continue
			}

			_, ok := visited[parsedLink.String()]
			if !ok {
				queue.Enqueue(parsedLink.String())
				continue
			}
		}
	}

	return sitemap
}
