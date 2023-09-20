package controllers

import (
	"encoding/json"
	"fmt"
	models2 "github.com/Yarandi/go-course/60_MVC/03_Controllers/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController struct{}

// NewUserController is a constructor for creating instances of the UserController. It returns a pointer to a newly created UserController instance.
func NewUserController() *UserController { //returns a pointer to a UserController instance
	return &UserController{}
}

// GetUser have to capitalized to be exported
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	u := models2.User{
		Id:     params.ByName("id"),
		Name:   "Hamed Yarandi",
		Age:    36,
		Gender: "Male",
	}

	//marshalling the values
	uj, _ := json.Marshal(u)

	//set json header for response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //200

	//print
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models2.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = "1247"

	uj, _ := json.Marshal(u)

	//set json header for response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //201

	//print
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "write code to delete")
}
