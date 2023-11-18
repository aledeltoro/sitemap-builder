package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSitemap(t *testing.T) {
	c := require.New(t)

	sitemap := NewSitemap()
	sitemap.AddURL("https://htmx.org")

	c.Len(sitemap.URL, 1)
}
