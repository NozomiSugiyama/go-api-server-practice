package handler

import (
	"encoding/json"
	"net/http"
	"github.com/NozomiSugiyama/wcdi-userstore/server/model"
	"github.com/gorilla/mux"
)

func UserList(w http.ResponseWriter, r *http.Request) {
	users, err := model.GetUsers()

	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

func UserShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := model.GetUser(vars["id"])

	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := model.DeleteUser(vars["id"]); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
