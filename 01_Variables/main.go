package main

import "fmt"

//var x int = 32
//var y bool = false
//var z float64 = 32.98
//var c string

//const name, age = "saeed", 44

type hotdog int

var t2 hotdog

func main() {

	t2 = 12
	fmt.Printf("%T \t %d \n", t2, t2)

	y := float32(t2)
	fmt.Printf("%T \t %f \n", y, y)

	//a := fmt.Sprintf("name is %s and age is %d", name, age)
	//s := fmt.Sprintf("%v \t %v \t %v \t %v", x, y, z, c)
	//fmt.Println(a, s)
}
