package client

import (
	"fmt"
	"io"
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

func TestGetPage(t *testing.T) {
	c := require.New(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, mockHTML)
	}))

	defer server.Close()

	res, err := GetPage(server.URL)
	c.NoError(err)

	defer func() {
		_ = res.Body.Close()
	}()

	page, err := io.ReadAll(res.Body)
	c.NoError(err)
	c.Equal(mockHTML, string(page))
}
