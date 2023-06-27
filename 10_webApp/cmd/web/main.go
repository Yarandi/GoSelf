package main

import (
	"fmt"
	"github.com/Yarandi/go-course/pkg/config"
	"github.com/Yarandi/go-course/pkg/render"
	"log"
	"net/http"

	handler "github.com/Yarandi/go-course/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {

	//config
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not create template cache ")
	}
	app.TemplateCache = tc

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Printf("starting Application on port %s ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
