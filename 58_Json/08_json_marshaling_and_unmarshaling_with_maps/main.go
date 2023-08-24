package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Car struct {
	Id      int
	Name    string
	Details map[string]string
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/unmarshal", unmarshal)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "you are in home!")
}

func marshal(w http.ResponseWriter, req *http.Request) {

	c1 := Car{
		Id:      01,
		Name:    "Nissan",
		Details: map[string]string{"Brand": "Nissan", "Model": "Roniz"},
	}
	jsonData, _ := json.Marshal(c1)
	//cmd
	//fmt.Println(string(jsonData))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
func unmarshal(w http.ResponseWriter, req *http.Request) {

	var data Car
	str := `{"Id":1,"Name":"Nissan","Details":{"Brand":"Nissan","Model":"Roniz"}}`

	json.Unmarshal([]byte(str), &data)
	fmt.Fprintln(w, data)

}
