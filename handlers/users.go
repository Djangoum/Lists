package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"home.com/lists/backend/presenters"
	"home.com/lists/backend/usecases/users"
)

func getUsers(users_service *users.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		users, err := users_service.ListUsers()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		toJ := []*presenters.User{}
		for _, user := range users {
			toJ = append(toJ, &presenters.User{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
			})
		}

		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error encoding users to json"))
		}

		return
	})
}

func createUser(users_service *users.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var addUserModel AddUser

		err := json.NewDecoder(r.Body).Decode(&addUserModel)

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error deserializing input data" + err.Error()))
			return
		}

		user, err := users_service.CreateUser(addUserModel.Email, addUserModel.Password, addUserModel.FirstName, addUserModel.LastName)

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error trying to create new user" + err.Error()))
			return
		}

		toJ := presenters.User{}

		toJ.Email = user.Email
		toJ.ID = user.ID
		toJ.FirstName = user.FirstName
		toJ.LastName = user.LastName

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error returning new user"))
			return
		}
	})
}

func RegisterUserHandlers(r *mux.Router, service *users.Service) {
	r.Handle("/v1/users", getUsers(service))

	r.Handle("/v1/user", createUser(service)).Methods("POST").Name("createUser")
}
