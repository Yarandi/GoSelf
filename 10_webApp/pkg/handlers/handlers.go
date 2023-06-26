package handler

import (
	"net/http"

	"github.com/Yarandi/go-course/pkg/render"
)

// Home is the about home handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}


