package core

import (
	"github.com/krishan1988/home24-web-analyser/internal/models"
	"github.com/krishan1988/home24-web-analyser/pkg/html/content"
	"golang.org/x/net/html"
)

func GetPageSummary(extr content.Extractor, selfURL string) (*models.PageSummary, error) {
	summary := models.NewPageSummary()
	hasForm := false

	doctypeEventHandler := func(n *html.Node) {
		if n == nil {
			return
		}
		summary.WithVersion(HTMLNode(*n).GetVersion())
	}

	elementEventHandler := func(n *html.Node) {
		if n == nil {
			return
		}

		switch tagName := n.Data; tagName {
		case "title":
			summary.WithTitle(n.FirstChild.Data)

		case "h1", "h2", "h3", "h4", "h5", "h6":
			summary.CountHeader(tagName)

		case "a":
			isExternal, err := HTMLNode(*n).IsExternalLink(selfURL)
			if err != nil {
				return
			}
			link, _ := HTMLNode(*n).AttrValue("href")
			summary.AddExternalLink(link)
			if isExternal {
				summary.CountLinkAsAccessible()
			}

		case "form":
			hasForm = true

		case "input":
			if hasForm && HTMLNode(*n).IsPasswordInput() {
				summary.HasLoginPage = true
			}
		}
	}

	doctypeEvent := content.NewEvent(html.DoctypeNode, doctypeEventHandler)
	elementEvent := content.NewEvent(html.ElementNode, elementEventHandler)

	extr.Extract(doctypeEvent, elementEvent)

	return summary, nil
}