package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	root "github.com/krishan1988/home24-web-analyser"
	"github.com/krishan1988/home24-web-analyser/internal/api/handler"
)

func Configure(port int) {
	
	// Initialize Gin
	router := gin.Default()

	tmpl := root.GetTemplates()

	// Define a route to render the template
	router.GET("/", handler.GetIndexHandler(tmpl))

	router.GET("/summary", handler.GetPageSummaryHandler(tmpl))

	// Run the server
	router.Run(fmt.Sprintf(":%v", port))
}