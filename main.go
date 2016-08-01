package main

import (
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "fmt"
  "html/template"
)

type Stuff struct {
  Breed string
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler).
    Methods("GET")
  r.HandleFunc("/dogs", DogHandlerPost).
    Methods("POST")
  r.HandleFunc("/dogs", DogHandler).
    Methods("GET")
  log.Fatal(http.ListenAndServe(":8000", r))
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
  stf := &Stuff{Breed: r.FormValue("breed")}
  render(w, "templates/dogs.html", stf)
}

func render(w http.ResponseWriter, filename string, data interface{}) {
  tmpl, err := template.ParseFiles(filename);
  if err != nil {
    fmt.Println("Blahhh")
  }
  if err := tmpl.Execute(w, data); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
