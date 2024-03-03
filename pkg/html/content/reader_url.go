package content

import (
	"net/http"

	"golang.org/x/net/html"
)

// readerWithUrl implements the Reader interface for reading HTML content from a URL.
type readerWithUrl struct {
	url string // url is the URL from which HTML content will be read.
}

// NewReaderWithUrl creates a new instance of ReaderWithUrl with the specified URL.
func NewReaderWithUrl(url string) Reader {
	return &readerWithUrl{
		url: url,
	}
}

// Read reads HTML content from the specified URL and returns the root node of the parsed HTML tree.
func (r *readerWithUrl) Read() (*html.Node, error) {
	// Fetch HTML content from the URL
	resp, err := http.Get(r.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the HTML content
	node, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return node, nil
}