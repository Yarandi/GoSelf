package main

import "fmt"

var i interface{}

func found(i interface{}) {
	fmt.Printf("Type %T, value %v \n", i, i)
}

func main() {
	s := "simpleLearn"
	found(s)
	d := 12.3
	found(d)
	g := 12
	found(g)
	q := false
	found(q)
}
