package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

type SomeType struct{
	TypeName string
	TypeNumber int
}

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}

func Test() {
	fmt.Println("Hey Im fetched Helpers package")
}

