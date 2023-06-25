package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateOld(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}


//m := make(map[string]int)
var tc = make(map[string]*template.Template)
//make a new render functtion for increasing efficency of this function with using cache
func RenderTemplate(w http.ResponseWriter, t string){
	var tmpl *template.Template
	var err error
	
	
	//checking in the map if the value exist or not
	//val, ok := mapName["foo"]

	//check to see if we already have the template in our cache
	_, ok := tc[t] 

	if !ok { //means there is no template in tc slice
		
		log.Println("creating template and adding to cache")
		
		//so we need to creat the template 
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err) 
		}
	
	}else{
		//we have the template in the cache
		log.Println("using cache template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err) 
	}

}

func createTemplateCache(t string) error {
	templates :=  []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	} 

	//parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil{
		return err
	}

	//add template to cache (map)
	tc[t] = tmpl
	return nil
}

