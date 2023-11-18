package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"

	"github.com/aledeltoro/sitemap-builder/internal/service"
)

func main() {
	domain := flag.String("domain", "https://htmx.org/", "Domain used to build a sitemap")

	sitemap := service.BuildSitemap(*domain)

	output, err := xml.MarshalIndent(sitemap, " ", " ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(xml.Header + string(output))
}
