package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type Words struct {
	Id          int
	Name        string
	Family      string
	Age         int
	FirstScore  float64
	SecondScore float64
}

func main() {

	//xs := []string{"zero", "one", "two", "three"}
	//data := struct {
	//	Words []string
	//	Lname string
	//}{xs, "Hamed"}

	//defined new struct for checking if and and predefined function in templates

	w1 := Words{
		Id:          1,
		Name:        "Hamed",
		Family:      "Yarandi",
		Age:         35,
		FirstScore:  12.3,
		SecondScore: 14.2,
	}
	w2 := Words{
		Id:          2,
		Name:        "Mina",
		Family:      "Fam",
		Age:         32,
		FirstScore:  11.8,
		SecondScore: 8.9,
	}
	data := []Words{w1, w2}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
