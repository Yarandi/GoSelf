package main

import (
	"fmt"
)

type Animals interface {
	whatSay() string
	whatEat() string
	HowManyLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Cow struct {
	Color    string
	HowTeeth int
}

func main() {

	dog := Dog{
		Name:  "Max",
		Breed: "German Shepered",
	}
	PrintInfo(&dog)
}

func (d *Dog) whatSay() string {
	return "WOOF"
}

func (d *Dog) whatEat() string {
	return "MEAT"
}

func (d *Dog) HowManyLegs() int {
	return 4
}

func PrintInfo(a Animals) {
	fmt.Println("this animal say", a.whatSay(), "and eat some", a.whatEat(), "and it has", a.HowManyLegs(), "legs")
}
