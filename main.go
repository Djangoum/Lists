package main

import (
	stdlog "log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"home.com/lists/backend/handlers"
	"home.com/lists/backend/middlewares"
	"home.com/lists/backend/repositories"
	"home.com/lists/backend/usecases/users"
)

func main() {
	db := repositories.CreateDatabaseConnection()
	usersRepo := repositories.NewUsersMysqlRepository(db)
	usersService := users.NewService(usersRepo)

	r := mux.NewRouter()
	handlers.RegisterListsHandlers(r)
	handlers.RegisterUserHandlers(r, usersService)

	httpPipeline := middlewares.RegisterMiddlewares(r)

	srv := &http.Server{
		Handler:      httpPipeline,
		Addr:         "127.0.0.1:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	var err = srv.ListenAndServe()

	stdlog.Fatal(err)
}
