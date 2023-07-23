package main

import (
	"log"
	"math"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
	"fdateDMY": dayMonthYear,
	"fdbl":     double,
	"fsq":      square,
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func dayMonthYear(t time.Time) string {
	return t.Format("02-01-2006")
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3) //time.Now()
	if err != nil {
		log.Fatalln(err)
	}
}
