package service

import (
	"log"
	"net/url"
	"strings"

	"github.com/aledeltoro/sitemap-builder/internal/client"
	"github.com/aledeltoro/sitemap-builder/internal/models"
)

// BuildSitemap searches a domain to build a sitemap according
// to the XML schema of the Sitemap protocol.
// The search is performed by doing a Breadth-First Search (BFS)
func BuildSitemap(domain string) *models.Sitemap {
	sitemap := models.NewSitemap()
	queue := models.NewQueue()
	visited := map[string]bool{}

	sitemap.AddURL(domain)
	queue.Enqueue(domain)
	visited[domain] = true

	for queue.Size() != 0 {
		currentDomain := queue.DeQueue()

		parsedDomain, err := url.Parse(currentDomain)
		if err != nil {
			log.Println(err)
			continue
		}

		links, err := client.GetPageLinks(parsedDomain.String())
		if err != nil {
			log.Println(err)
			continue
		}

		for _, rawLink := range links {
			link := processLink(parsedDomain, rawLink.Href)
			if link == "" {
				continue
			}

			if _, ok := visited[link]; !ok {
				sitemap.AddURL(link)
				queue.Enqueue(link)
				visited[link] = true
			}
		}
	}

	return sitemap
}

func processLink(parsedDomain *url.URL, rawLink string) string {
	link, err := parsedDomain.Parse(rawLink)
	if err != nil {
		return ""
	}

	doesDomainMatch := strings.Contains(link.Host, parsedDomain.Hostname())

	if !doesDomainMatch {
		return ""
	}

	return link.String()
}
