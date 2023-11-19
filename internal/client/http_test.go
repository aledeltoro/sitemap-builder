package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	mockHTML = `<html>
		<body>
			<h1>Hello!</h1>
			<a href="/other-page">A link to another page</a>
			<a href="/dog">
				<span>Something in a span</span>
				Text not in a span
				<b>Bold text!</b>
			</a>
		</body>
		</html>`
)

func TestGetPageLinks(t *testing.T) {
	c := require.New(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, mockHTML)
	}))

	defer server.Close()

	links, err := GetPageLinks(server.URL)
	c.NoError(err)
	c.Len(links, 2)
}
