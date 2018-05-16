package model

import (
	"errors"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	DisplayName string        `json:"display_name" bson:"display_name"`
	FirstName   string        `json:"first_name" bson:"first_name"`
	LastName    string        `json:"last_name" bson:"last_name"`
	Birthday    string        `json:"birthday" bson:"birthday"`
	Gender      string        `json:"gender" bson:"gender"`
	Email       string        `json:"email" bson:"email"`
	PhoneNumber string        `json:"phone_number" bson:"phone_number"`
}

func GetUsers() ([]User, error) {
	var session = GetSession("mongodb")
	defer session.Close()

	db := session.DB("test")

	var users []User
	err := db.C("users").Find(bson.M{}).All(&users)
	if err != nil {
		fmt.Println("err: ", err)
		return users, err
	}
	fmt.Println("Results All: ", users)

	return users, nil
}

func CreateUser(user User) error {
	return nil
}

func GetUser(id string) (User, error) {
	var session = GetSession("mongodb")
	defer session.Close()

	db := session.DB("test")

	user := User{}
	if !bson.IsObjectIdHex(id) {
		fmt.Println("err: ", "invalid id")
		return user, errors.New("invalid id")
	}

	err := db.C("users").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user)
	if err != nil {
		fmt.Println("err: ", err)
		return user, err
	}
	fmt.Println("Results: ", user, id)

	return user, nil
}

func UpdateUser(user User) error {
	return nil
}

func DeleteUser(id string) error {
	return nil
}
