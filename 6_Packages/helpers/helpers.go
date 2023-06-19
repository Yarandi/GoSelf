package helpers

import (
	"fmt"
	"math/rand"
)

type SomeType struct{
	TypeName string
	TypeNumber int
}

func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}

func Test() {
	fmt.Println("Hey Im fetched Helpers package")
}

