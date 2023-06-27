package render

import (
	"bytes"
	"github.com/Yarandi/go-course/pkg/helpers"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//get the template cache from the app config

	//create a template cache
	tc, err := CreateTemplateCache()
	helpers.DD(tc)

	if err != nil {
		log.Fatal(err)
	}

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	//get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl") // glob return file names
	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.page.tmpl
	for _, page := range pages {

		//get the actual filename
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//looking for layouts
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			tm, err := ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
			helpers.DD(tm)
			//fmt.Println(ts)
		}

		myCache[name] = ts
		//fmt.Println(myCache)
	}

	return myCache, nil
}
