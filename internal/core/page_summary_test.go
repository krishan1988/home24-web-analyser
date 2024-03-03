package core

import (
	"strings"
	"testing"

	"github.com/krishan1988/home24-web-analyser/pkg/html/content"
	"golang.org/x/net/html"
)

func TestGetPageSummary(t *testing.T) {
	const htmlStr = `<!DOCTYPE html>
	<html>
	<head>
		<title>Sample</title>
	</head>
	<body>
		<p>more content</p>
		<h1>more content</h1>
		<a href="www.google.com">more content</p>
	</body>
	</html>`

	// Parse HTML string
	root, _ := html.Parse(strings.NewReader(htmlStr))

	// Create content extractor
	extr := content.NewExtractor(root)

	// Get page summary
	summary, err := GetPageSummary(extr, "")
	if err != nil {
		t.Error(err)
	}

	t.Log(summary)
}