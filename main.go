package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/julienschmidt/httprouter"
)

func Index(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
  fmt.Fprintln(response, "Welcome!")
}

func TodoIndex(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
  fmt.Fprintln(response, "TodoIndex!")
}

func TodoShow(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
  todoId := params.ByName("todoId")

  fmt.Fprintln(response, "TodoShow: ", todoId)
}

func main() {
  router := httprouter.New()

  router.GET("/", Index)
  router.GET("/todos", TodoIndex)
  router.GET("/todos/:todoId", TodoShow)

  log.Fatal(http.ListenAndServe(":8080", router))
}
