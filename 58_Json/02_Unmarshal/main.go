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
	http.HandleFunc("/unmarshal", unmarshal)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h6>we are in /home, take a look at /unmarshal </h6>")
}

func unmarshal(w http.ResponseWriter, r *http.Request) {
	var data product
	str := `{"Id":14432,"ProductId":456425,"ProductName":"MacBook Pro M768bb","WarehousesName":["DGK","WAR","AMZ","ASD"]}`

	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		log.Fatalln("error in unmarshal", err)
	}
	fmt.Println(data)
	for i, v := range data.WarehousesName {
		fmt.Println(i, v)
	}
}
