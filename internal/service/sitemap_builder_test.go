package service

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildSitemap(t *testing.T) {
	c := require.New(t)

	sitemap := BuildSitemap("https://www.y.uno/")

	output, err := xml.MarshalIndent(sitemap, " ", " ")
	c.NoError(err)
	fmt.Println(xml.Header + string(output))
}
