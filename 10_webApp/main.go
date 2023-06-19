package main

import (
	"errors"
	"fmt"
	"net/http"
)
 
const portNumber = ":8080"

//Home is the about home handler
func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the home page")
}

//About is the about page handler
func About(w http.ResponseWriter, r *http.Request){
	sum := addValues(2,3)
	_ , _ = fmt.Fprintf(w, fmt.Sprint("this is the about page and 2 + 3 is %d is ", sum))
}

//add two integfers and return the sum
func addValues(x, y int) (int){
    return x + y 
}


func Divide(w http.ResponseWriter, r *http.Request){
	 f, err := divideValue(100.0, 0.0)
	 if err != nil {
		fmt.Fprintf(w, "can not divide by zero ")
		return
	 }
	 fmt.Fprintf(w , "%f divided by %f is %f", 100.0, 0.0, f)
}

func divideValue(x , y float32) (float32, error){
	if y <= 0{
		err := errors.New("can not divide by zero at all")
		return 0, err
	}
	result := x / y
	return result, nil
}

//main is the main application function
func main(){

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)




	fmt.Printf("starting Application on port %s ", portNumber)
	_ = http.ListenAndServe(portNumber,nil)

}