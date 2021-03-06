package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
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

// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "<h1>Welcome to the Go Server!</h1>")
// }

// func contact(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:cooperalex512@gmail.com\">cooperalex512@gmail.com</a>")
// }

// func faq(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "<h1>FAQ Page!</h1><p>Here is a list of some of the frequently asked questions.</p>")
// }

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404, Page not found</h1>")
}

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to the Go Server!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:cooperalex512@gmail.com\">cooperalex512@gmail.com</a>")
}

func faq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>FAQ Page!</h1><p>Here is a list of some of the frequently asked questions.</p>")
}

func main() {
	r := httprouter.New()
	r.NotFound = http.HandlerFunc(notFound)
	r.GET("/", home)
	r.GET("/contact", contact)
	r.GET("/faq", faq)

	// r := mux.NewRouter()
	// r.NotFoundHandler = http.HandlerFunc(notFound)
	// r.HandleFunc("/", home)
	// r.HandleFunc("/contact", contact)
	// r.HandleFunc("/faq", faq)
	http.ListenAndServe(":4000", r)
}
