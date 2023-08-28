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
	router := httprouter.New()

	router.GET("/", index)

	//add route for creating user
	router.POST("/user", createUser)

	//add route plus parameters for getting user
	router.GET("/user/:id", getUser)

	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

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

func createUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{}

	//decode to save
	err := json.NewDecoder(req.Body).Decode(&u)

	if err != nil {
		log.Fatalln(err)
	}

	//change Id
	u.Id = "879054"

	//marshal to show again
	jsonData, _ := json.Marshal(u)

	//write content type / status code / and Payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	fmt.Fprintf(w, "%s \n", jsonData)
}
