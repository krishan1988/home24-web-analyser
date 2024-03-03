package content

import (
	"bytes"
	"reflect"
	"testing"

	"golang.org/x/net/html"
)

func TestExtractor_Extract(t *testing.T) {
	// Create a sample HTML document for testing.
	htmlContent := `
	<html>
		<head>
			<title>Title</title>
		</head>
		<body>
			<h1>Heading 1</h1>
			<p>Paragraph 1</p>
			<h2>Heading 2</h2>
			<p>Paragraph 2</p>
		</body>
	</html>`

	// Parse the sample HTML content.
	rootNode, err := html.Parse(bytes.NewReader([]byte(htmlContent)))
	if err != nil {
		t.Fatal("failed to parse HTML:", err)
	}

	// Define mock callback functions for testing.
	var extractedTags []html.NodeType
	callbackFunc := func(node *html.Node) {
		extractedTags = append(extractedTags, node.Type)
	}

	// Create an instance of the extractor.
	extractor := NewExtractor(rootNode)

	// Define events to extract h1 and h2 tags.
	events := []Event{
		{html.ElementNode, callbackFunc}, // h1 tag
		{html.ElementNode, callbackFunc}, // h2 tag
	}

	// Perform the extraction.
	extractor.Extract(events...)

	// Define the expected result (order matters).
	expectedTags := []html.NodeType{html.ElementNode, html.ElementNode}

	// Check if the extracted tags match the expected result.
	if !reflect.DeepEqual(extractedTags, expectedTags) {
		t.Errorf("extracted tags mismatch: expected %v, got %v", expectedTags, extractedTags)
	}
}