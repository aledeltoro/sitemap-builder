package models

import "encoding/xml"

const (
	sitemapNamespace = "http://www.sitemaps.org/schemas/sitemap/0.9"
)

// Sitemap struct for the XML schema for the Sitemap protocol
type Sitemap struct {
	XMLName   xml.Name `xml:"urlset"`
	Namespace string   `xml:"xmlns,attr"`
	URL       []URL    `xml:"url"`
}

// URL struct to display an URL in a Sitemap
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

// AddURL appends an URL to the existing Sitemap
func (s *Sitemap) AddURL(url string) {
	s.URL = append(s.URL, URL{Locator: url})
}
