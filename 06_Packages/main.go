package main

import (
	"fmt"

	"github.com/Yarandi/myniceprogram/helpers"
)

func main() {

	var myvar helpers.SomeType
	myvar.TypeName = "Custom name"
	myvar.TypeNumber = 85

	fmt.Println(myvar.TypeName)
	helpers.Test()
}
