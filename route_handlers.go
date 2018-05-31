package main

import (
  "net/http"
  "fmt"
  "encoding/json"
)

type Credentials struct {
  Username string `json: "username"`
  Password string `json: "password"`
}

var users []Credentials

func SignUp(w http.ResponseWriter, r *http.Request) {
  user := Credentials{}

  err := r.ParseForm()

  if err != nil {
    fmt.Println(fmt.Errorf("Error: %v", err))
    w.WriteHeader(http.StatusInternalServerError)

    return
  }

  user.Username = r.Form.Get("username")
  user.Password = r.Form.Get("password")

  err = store.CreateUser(&user)

  if(err != nil) {
    fmt.Println(err)
  }
}
