package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type product struct {
	Id             int
	ProductId      int
	ProductName    string
	WarehousesName []string
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encode)
	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h6>we are in /home, take a look at /marshal or /encode</h6>")
}

func marshal(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p1 := product{
		Id:             15432,
		ProductId:      1487990,
		ProductName:    "MacBook Pro M768bb",
		WarehousesName: []string{"DGK", "Amazoon", "AIRV"},
	}

	json, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func encode(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p1 := product{
		Id:             14432,
		ProductId:      456425,
		ProductName:    "MacBook Pro M768bb",
		WarehousesName: []string{"DGK", "WAR", "AMZ", "ASD"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
