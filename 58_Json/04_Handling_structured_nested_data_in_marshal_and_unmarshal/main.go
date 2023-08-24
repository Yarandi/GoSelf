package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//------------------------
//Nested json data
//------------------------

type Car struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Speed  float64 `json:"speed"`
	Engine Engine  `json:"engine"`
}

type Engine struct {
	CC       int    `json:"cc"`
	Cylinder int    `json:"cylinder"`
	Oil      string `json:"oil"`
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/unmarshal", unmarshal)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h6>we are in /home, take a look at /marshal or /unmarshal</h6>")
}

func marshal(w http.ResponseWriter, req *http.Request) {
	c1 := Car{
		Id:    001,
		Name:  "Jeep",
		Speed: 210,
		Engine: Engine{
			CC:       2400,
			Cylinder: 4,
			Oil:      "2050",
		},
	}

	jsonData, err := json.Marshal(c1)
	if err != nil {
		log.Fatalln(err)
	}

	//cmd
	//fmt.Println(string(jsonData))

	//set json header
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func unmarshal(w http.ResponseWriter, req *http.Request) {

	var data Car

	str := `{"id":1,"name":"Jeep","speed":210,"engine":{"cc":2400,"cylinder":4,"oil":"2050"}}`

	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(w, data)
}
