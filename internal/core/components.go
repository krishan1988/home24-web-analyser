package core

import (
	"errors"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

type HTMLNode html.Node

func (n HTMLNode) Valid(tagName string) bool {
	return n.Data == tagName
}

func (n HTMLNode) AttrValue(key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func (n HTMLNode) IsExternalLink(selfURL string) (bool, error) {
	href, found := n.AttrValue("href")
	if !found {
		return false, errors.New("anchor without href")
	}

	externalURLParsed, err := url.Parse(href)
	if err != nil || externalURLParsed.Host == ""{
		return false, nil
	}

	internalURLParsed, err := url.Parse(selfURL)
	if err != nil || internalURLParsed.Host == "" {
		return false, nil
	}

	return internalURLParsed.Host != externalURLParsed.Host, nil
}

func (n HTMLNode) IsPasswordInput() bool {
	for _, attr := range n.Attr {
		if attr.Key == "type" && attr.Val == "password" {
			return true
		}
	}
	return false
}

func (n HTMLNode) GetVersion() string {
	return strings.TrimPrefix(n.Data, "html ")
}