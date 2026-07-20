package web

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func Run(addr string) error {

	router := gin.Default()

	tmpl := template.Must(template.ParseFiles(
		"web/templates/index.html",

		"web/templates/partials/uploads.html",
		"web/templates/partials/favorites.html",
		"web/templates/partials/dictionaries.html",
		"web/templates/partials/jobs.html",
	))

	router.SetHTMLTemplate(tmpl)

	registerRoutes(router)

	return router.Run(addr)
}
