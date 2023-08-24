package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Car struct {
	Id    int     `json:"Id"`
	Name  string  `json:"NameOfCar"`
	Speed float64 `json:"-"` //Omit this field from result of marshal and unmarshal
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
		Name:  "Nissan",
		Speed: 210,
	}

	json, err := json.Marshal(c1)
	if err != nil {
		log.Fatalln(err)
	}

	//set json header
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func unmarshal(w http.ResponseWriter, req *http.Request) {

	var data Car
	str := `{"Id":1,"Name":"Nissan","Speed":210}`

	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprint(w, data)

}
