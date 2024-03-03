package core_test

import (
	"testing"

	"github.com/krishan1988/home24-web-analyser/internal/core"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestHTMLNode_Valid(t *testing.T) {
	
	node := core.HTMLNode{Data: "div"}
	assert.True(t, node.Valid("div"))

	node = core.HTMLNode{Data: "p"}
	assert.False(t, node.Valid("div"))
}

func TestHTMLNode_AttrValue(t *testing.T) {

	node := core.HTMLNode{Attr: []html.Attribute{{Key: "class", Val: "example"}}}
	val, exists := node.AttrValue("class")
	assert.True(t, exists)
	assert.Equal(t, "example", val)

	_, exists = node.AttrValue("id")
	assert.False(t, exists)
}

func TestHTMLNode_IsExternalLink(t *testing.T) {
	node := core.HTMLNode{Attr: []html.Attribute{{Key: "href", Val: "https://example.com"}}}
	isExternal, err := node.IsExternalLink("https://your-site.com")
	assert.NoError(t, err)
	assert.True(t, isExternal)

	node = core.HTMLNode{Attr: []html.Attribute{{Key: "href", Val: "/about"}}}
	isExternal, err = node.IsExternalLink("https://example.com")
	assert.NoError(t, err)
	assert.False(t, isExternal)

	node = core.HTMLNode{}
	_, err = node.IsExternalLink("https://example.com")
	assert.Error(t, err)

	node = core.HTMLNode{Attr: []html.Attribute{{Key: "href", Val: ":"}}}
	_, err = node.IsExternalLink("https://example.com")
	assert.Error(t, err)

	node = core.HTMLNode{Attr: []html.Attribute{{Key: "href", Val: "/about"}}}
	_, err = node.IsExternalLink(":")
	assert.Error(t, err)
}

func TestHTMLNode_IsPasswordInput(t *testing.T) {
	node := core.HTMLNode{Attr: []html.Attribute{{Key: "type", Val: "password"}}}
	assert.True(t, node.IsPasswordInput())

	node = core.HTMLNode{Attr: []html.Attribute{{Key: "type", Val: "text"}}}
	assert.False(t, node.IsPasswordInput())

	node = core.HTMLNode{}
	assert.False(t, node.IsPasswordInput())
}

func TestHTMLNode_GetVersion(t *testing.T) {
	// Test for valid HTML version
	node := core.HTMLNode{Data: "html 5"}
	assert.Equal(t, "5", node.GetVersion())

	// Test for HTML version with unexpected format
	node = core.HTMLNode{Data: "html "}
	assert.Equal(t, "", node.GetVersion())
}