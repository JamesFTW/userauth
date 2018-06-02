package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "database/sql"
  "github.com/gorilla/mux"
  "io/ioutil"
  _"github.com/lib/pq"
)

func newRouter() *mux.Router {
  r := mux.NewRouter()

  r.HandleFunc("/signup", SignUp).Methods("POST")
  r.HandleFunc("/hello", parseGhPost).Methods("POST")


  return r
}

func main() {
  fmt.Println("Start server...")

  connString := "postgres://iggdsjbhoglwwe:9e1031e4498edce6b02933786cdeacf123dd5860335414ad1022223cc7dd4e32@ec2-54-235-132-202.compute-1.amazonaws.com:5432/d5v4uga3qon34i"
  db, err := sql.Open("postgres", connString)

  if err != nil {
    panic(err)
  }

  err = db.Ping()

  if err != nil {
    panic(err)
  }

  InitStore(&dbStore{db: db})

  r := newRouter()
  http.ListenAndServe(":8080", r)
}


func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello World")
}
