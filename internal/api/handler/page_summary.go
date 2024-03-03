package handler

import (
	"html/template"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/krishan1988/home24-web-analyser/internal/core"
	"github.com/krishan1988/home24-web-analyser/pkg/html/content"
)

// GetPageSummaryHandler returns a Gin handler function for displaying the page summary.
func GetPageSummaryHandler(tmpl *template.Template) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Extract URL parameter from the query string
		url := ctx.Query("url")

		// Validate the URL using the IsValidURL function
		if !IsValidURL(url) {
			// If the URL is invalid, render the error template with an error message
			err := tmpl.ExecuteTemplate(ctx.Writer, "error.tmpl", gin.H{
				"ErrorMessage": "Invalid URL",
			})
			if err != nil {
				ctx.String(http.StatusInternalServerError, "Error rendering template")
				return
			}
			return
		}

		// // If the URL is valid, create a sample PageSummary for demonstration
		// pageSummary := models.PageSummary{
		// 	Version:             "HTML5",
		// 	Title:               "Sample HTML Document",
		// 	HeaderCount:         map[string]int{"h1": 3, "h2": 5, "h3": 2},
		// 	ExternalLinks:       []string{"https://example.com", "https://google.com"},
		// 	InternalLinks:       []string{"/about", "/contact"},
		// 	AccessibleLinkCount: 4,
		// 	HasLoginPage:        true,
		// }

		rootNode, err := content.NewReaderWithUrl(url).Read()
		if err != nil {
			err := tmpl.ExecuteTemplate(ctx.Writer, "error.tmpl", gin.H{
				"ErrorMessage": "Invalid URL",
			})
			if err != nil {
				ctx.String(http.StatusInternalServerError, "Error rendering template")
				return
			}
			return
		}

		selfUrl := ctx.Request.URL.String()
		pageSummary, err := core.GetPageSummary(content.NewExtractor(rootNode), selfUrl)
		if err != nil {
			err := tmpl.ExecuteTemplate(ctx.Writer, "error.tmpl", gin.H{
				"ErrorMessage": "Unable to get the page summary",
			})
			if err != nil {
				ctx.String(http.StatusInternalServerError, "Error rendering template")
				return
			}
			return
		}

		// Render the page_summary template with the PageSummary data
		err = tmpl.ExecuteTemplate(ctx.Writer, "page_summary.tmpl", pageSummary)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error rendering template")
			return
		}
	}
}

// IsValidURL checks if the given string is a valid URL using regex.
func IsValidURL(input string) bool {
	const urlPattern = `^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-zA-Z0-9]+([\-\.]{1}[a-zA-Z0-9]+)*\.[a-zA-Z]{2,5}(:[0-9]{1,5})?(\/.*)?$`

	match, err := regexp.MatchString(urlPattern, input)
	if err != nil {
		return false
	}
	return match
}