package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Driver struct {
	Name    string
	Age     int
	Formula string
}

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Name  string
	Model int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	Buddha := sage{
		Name:  "Buddha",
		Motto: "The beleif of no beliefs",
	}
	Gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	car1 := car{
		Name:  "Roniz",
		Model: 2007,
	}

	car2 := car{
		Name:  "Hilux",
		Model: 2013,
	}

	cars := []car{car1, car2}
	sages := []sage{Buddha, Gandhi}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	//m := Driver{
	//	Name:    "Chomakher",
	//	Age:     32,
	//	Formula: "F1",
	//}

	//m := map[string]string{
	//	"first":  "one",
	//	"second": "two",
	//	"third":  "three",
	//	"fourth": "four",
	//}
	//m := []string{"map1", "map2", "map3"}
	//m := 32
	//m := "Hello every one"

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}

	P1 := struct {
		fname string
		lname string
	}{
		"james",
		"Bond",
	}
	fmt.Println(P1)
}
