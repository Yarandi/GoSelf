package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Car struct {
	Id      int
	Name    string
	Details []string
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/marshal", marshal)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "you are in home!")
}

func marshal(w http.ResponseWriter, req *http.Request) {

	c1 := Car{
		Id:      01,
		Name:    "Nissan",
		Details: []string{"Off-road", "rough-road"},
	}
	// Marshaling with indentation
	jsonData, _ := json.MarshalIndent(c1, "", "    ")

	//MarshalIndent works on cms
	fmt.Println("Indented JSON:")
	fmt.Println(string(jsonData))

	//TODO:: dose not work on web we need more time to check
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(jsonData)
}
