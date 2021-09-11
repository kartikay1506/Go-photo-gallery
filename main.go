package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	if r.URL.Path == "/" {
// 		fmt.Fprint(w, "<h1>Welcome to the Go Server!</h1>")
// 	} else if r.URL.Path == "/contact" {
// 		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:cooperalex512@gmail.com\">cooperalex512@gmail.com</a>")
// 	} else {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprintf(w, "404, Page Not Found")
// 	}
// }

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to the Go Server!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:cooperalex512@gmail.com\">cooperalex512@gmail.com</a>")
}

// func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	fmt.Fprint(w, "Welcome!\n")
// }

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

func main() {
	// router := httprouter.New()
	// router.GET("/", Index)
	// router.GET("/hello/:name", Hello)

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":4000", r)
}
