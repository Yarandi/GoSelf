package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	//for {
	fmt.Fprintf(w, "Hello from deployed app on macos!!")
	//time.Sleep(time.Second * 5)
	//}
}

//-----------------------------------------------------------------------------------------
//sudo launchctl load /Library/LaunchDaemons/com.example.mygoapponmacos.plist
//sudo log stream --predicate 'senderImagePath CONTAINS "com.example.mygoapponmacos.plist"'
//sudo launchctl unload /Library/LaunchDaemons/com.example.mygoapponmacos.plist
//-----------------------------------------------------------------------------------------
