package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 page not implemented", http.StatusNotImplemented)
}

func NotImplementedHandler() http.Handler {
	return http.HandlerFunc(NotImplemented)
}

func main() {
	port := ":3000"
	webRoot := os.Getenv("WEBROOT")

	router := mux.NewRouter()

	// Handle form submitted by user
	router.HandleFunc("/api/hello",
		func(res http.ResponseWriter, req *http.Request) {
			name := req.FormValue("name")
			age := req.FormValue("age")
			fmt.Println("Hello from", name, "[", age, "]")
			formatter := render.New(render.Options{
				IndentJSON: true,
			})
			formatter.JSON(res, http.StatusOK, struct {
				Name string `json:"name"`
				Age  string `json:"age"`
			}{Name: name, Age: age})
		})

	// Handle unknown, send 501 not implemented
	router.Handle("/unknown", NotImplementedHandler())

	// Handle static files(my blog)
	router.PathPrefix("/").Handler(
		http.StripPrefix("/", http.FileServer(http.Dir(webRoot+"./static/"))))

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(router)

	fmt.Printf("Listening to port %v\n", port)
	http.ListenAndServe(port, n)
}
