package middlewares

import (
	"net/http"

	"github.com/urfave/negroni"
)

// RegisterMiddlewares adds all necesary middlewares to build the http pipeline
func RegisterMiddlewares(handler http.Handler) http.Handler {
	n := negroni.Classic()

	n.UseHandler(handler)

	return n
}
