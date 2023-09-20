package main

import (
	"github.com/Yarandi/go-course/60_MVC/03_Controllers/controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	uc := controllers.NewUserController()
	r := httprouter.New()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
