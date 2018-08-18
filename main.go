package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
}

func main() {
  router := httprouter.New()

  router.GET("/", Index)
  router.GET("/hello/:name", Hello)

//   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
//   })

  log.Fatal(http.ListenAndServe(":8080", router))
}
