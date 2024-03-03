package content

import (
	"golang.org/x/net/html"
)

// Extractor defines the interface for extracting content based on provided events.
type Extractor interface {
	Extract(events ...Event)
}

// extractor implements the Extractor interface.
type extractor struct {
	root *html.Node // root is the HTML node from which extraction starts.
}

// NewExtractor creates a new instance of the Extractor.
func NewExtractor(root *html.Node) Extractor {
	return &extractor{root: root}
}

// Extract performs content extraction based on the provided events.
func (e extractor) Extract(events ...Event) {
	// Map to store callback functions for each HTML node type.
	eventCallbacks := make(map[html.NodeType]CallbackFunc)

	// Populate the eventCallbacks map with provided events.
	for _, event := range events {
		eventCallbacks[event.nodeType] = event.callbackFunc
	}

	// Define a recursive crawler function to traverse HTML nodes and trigger callbacks.
	var crawl func(*html.Node)
	crawl = func(node *html.Node) {
		// Trigger callback function if registered for the current node type.
		if callbackFunc, ok := eventCallbacks[node.Type]; ok {
			callbackFunc(node)
		}

		// Recursively crawl through child nodes.
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawl(child)
		}
	}

	// Start crawling from the root node.
	crawl(e.root)
}