package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getLists() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		return
	})
}

func createList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var newList AddList
		err := json.NewDecoder(r.Body).Decode(&newList)

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error deserializing create list command"))
			return
		}

		return
	})
}

func RegisterListsHandlers(r *mux.Router) {
	r.Handle("/v1/lists", getLists())
}
