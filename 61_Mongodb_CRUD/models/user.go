package models

type User struct {
	Name   string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Age    int    `json:"age" bson:"age"`
	Id     string `json:"id" bson:"_id,omitempty"`
}
