package models

import "encoding/xml"

const (
	sitemapNamespace = "http://www.sitemaps.org/schemas/sitemap/0.9"
)

type Sitemap struct {
	XMLName   xml.Name `xml:"urlset"`
	Namespace string   `xml:"xmlns,attr"`
	URL       []URL    `xml:"url"`
}

type URL struct {
	XMLName xml.Name `xml:"url"`
	Locator string   `xml:"loc"`
}

// NewSitemap returns a new Sitemap instance
func NewSitemap() *Sitemap {
	return &Sitemap{
		Namespace: sitemapNamespace,
		URL:       make([]URL, 0),
	}
}
