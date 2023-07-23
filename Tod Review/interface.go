package main

import "fmt"

type MotorVehicle interface {
	Mileage() float64
}

type BMW struct {
	distance     float64
	averageSpeed float64
	fuel         float64
}
type Audi struct {
	distance float64
	fuel     float64
}
type Nissan struct {
	distance     float64
	averageSpeed float64
	fuel         float64
}

func (b BMW) Mileage() float64 {
	return b.distance / b.fuel
}
func (a Audi) Mileage() float64 {
	return a.distance / a.fuel
}
func (n Nissan) Mileage() float64 {
	return n.distance / n.fuel
}

func totalMileage(m []MotorVehicle) {

	tm := 0.0
	for _, v := range m {
		//fmt.Printf("%T \t", v)
		tm += v.Mileage()
	}
	fmt.Printf("Total mileage in month is %f", tm)
}

func main() {
	b1 := BMW{
		distance:     480.5,
		averageSpeed: 120,
		fuel:         50,
	}

	a1 := Audi{
		distance: 420.8,
		fuel:     47.0,
	}

	n1 := Nissan{
		distance: 420,
		fuel:     60.0,
	}

	person := []MotorVehicle{a1, b1, n1}
	totalMileage(person)
	//fmt.Printf("%T", person) //[]main.MotorVehicle
}
