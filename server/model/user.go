package model

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID                   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	DisplayName          string        `json:"display_name" bson:"display_name"`
	FirstName            string        `json:"first_name" bson:"first_name"`
	LastName             string        `json:"last_name" bson:"last_name"`
	FirstNameEnglish     string        `json:"first_name_english" bson:"first_name_english"`
	LastNameEnglish      string        `json:"last_name_english" bson:"last_name_english"`
	Birthday             string        `json:"birthday" bson:"birthday"`
	Gender               string        `json:"gender" bson:"gender"`
	Email                string        `json:"email" bson:"email"`
	EmailForShared       string        `json:"email_for_shared" bson:"email_for_shared"`
	CountryPrefecture    string        `json:"country_prefecture" bson:"country_prefecture"`
	PhoneNumber          string        `json:"phone_number" bson:"phone_number"`
	SlackID              string        `json:"slack_id" bson:"slack_id"`
	GithubID             string        `json:"github_id" bson:"github_id"`
	License              []string      `json:"license" bson:"license"`
	SubscriptionReligion struct {
		ID string `json:"_id" bson:"_id,omitempty"`
	} `json:"subscription_religion" bson:"subscription_religion"`
}

func GetUsers() ([]User, error) {
	var session = GetSession("mongodb")
	defer session.Close()

	db := session.DB("test")

	var users []User
	db.C("users").Find(bson.M{}).All(&users)
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
	db.C("users").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user)
	fmt.Println("Results: ", user, id)

	return user, nil
}

func UpdateUser(user User) error {
	return nil
}

func DeleteUser(id string) error {
	return nil
}
