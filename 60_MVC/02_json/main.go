package main

import (
	"GoSelfStudy/60_MVC/02_json/models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	r := httprouter.New()

	r.GET("/", index)

	//add route for creating user
	r.POST("/user", createUser)

	//add route plus parameters for getting user
	r.GET("/user/:id", getUser)

	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

// getting user
func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	u := models.User{
		Id:     p.ByName("id"),
		Name:   "Hamed Yarandi",
		Age:    35,
		Gender: "Male",
	}

	//marshal into json
	jsonData, err := json.Marshal(u)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	// Write JSON data directly to the response
	w.Write(jsonData)
}

// create user
func createUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}

	//decode json from curl to save on u
	//curl -X POST -H "Content-Type: application/json" -d '{"Name":"Hamed", "Age":35, "Gender":"male"}' http://localhost:8080/user
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		log.Fatalln(err)
	}

	//change Id value
	u.Id = "3225"

	//marshal to show in json format
	jsonData, _ := json.Marshal(u)

	//write content type / status code / and Payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	fmt.Fprintf(w, "%s \n", jsonData)
}
