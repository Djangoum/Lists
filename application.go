package main

import (
	stdlog "log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"home.com/lists/backend/middlewares"
)

func main() {
	// db := repositories.CreateDatabaseConnection()
	// usersRepo := repositories.NewUsersMysqlRepository(db)
	// usersService := users.NewService(usersRepo)

	r := mux.NewRouter()
	// handlers.RegisterListsHandlers(r)
	// handlers.RegisterUserHandlers(r, usersService)

	httpPipeline := middlewares.RegisterMiddlewares(r)

	r.HandleFunc("/greet", GreeterHandler)

	srv := &http.Server{
		Handler:      httpPipeline,
		Addr:         ":5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	var err = srv.ListenAndServe()

	stdlog.Fatal(err)
}

func GreeterHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
	w.WriteHeader(http.StatusOK)
}
