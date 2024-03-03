package handler

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIndexHandler returns a Gin handler function for serving the index page.
func GetIndexHandler(tmpl *template.Template) gin.HandlerFunc {
	
	return func(ctx *gin.Context) {
		// Render the index template
		err := tmpl.ExecuteTemplate(ctx.Writer, "index.tmpl", nil)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error rendering template")
			return
		}
	}
}