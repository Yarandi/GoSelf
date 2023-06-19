package main

import (
	"fmt"
	"net/http"
)

func main(){

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		n, err := fmt.Fprintf(w, "Hello web app worlds")
	
		if err != nil{
			fmt.Println(err)
		}
        //fmt.Println(fmt.Sprintf("Number of Bytes : %d", n))
        fmt.Printf("Number of Bytes : %d", n)
 
	})
	_ = http.ListenAndServe(":8080",nil)



}