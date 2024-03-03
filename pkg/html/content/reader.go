package content

import "golang.org/x/net/html"

// Reader defines an interface for reading HTML content.
type Reader interface {
	Read() (*html.Node, error)
}