package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Stuff struct {
	Breed string
	Other string
	Blue  string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).
		Methods("GET")
	r.HandleFunc("/dogs", DogHandlerPost).
		Methods("POST")
	r.HandleFunc("/dogs", DogHandler).
		Methods("GET")
	r.HandleFunc("/stinnette", StinnetteHandler).
		Methods("GET", "DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func StinnetteHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/david.html")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	http.ServeFile(w, r, "templates/index.html")
}

func DogHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/dogs.html")
}

func DogHandlerPost(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.New("dogbreed")
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", v)
	}
	fmt.Println(r.FormValue("breed"))
	stf := Stuff{Breed: r.FormValue("breed"), Other: "other", Blue: "true"}
	more := Stuff{Breed: r.FormValue("breed"), Other: "as;ldfja;slkjdfds;", Blue: "Do"}
	evenmore := Stuff{Breed: r.FormValue("breed"), Other: "AHH", Blue: "Clue"}

	stfArr := []Stuff{stf, more, evenmore}
	render(w, "templates/dogs.html", stfArr)
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		fmt.Println("Blahhh")
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
