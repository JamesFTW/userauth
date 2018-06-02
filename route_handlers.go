package main

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "fmt"
)

type Credentials struct {
  Username string `json: "username"`
  Password string `json: "password"`
}

func InternalServerError(err error, w http.ResponseWriter) {
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
}

func SignUp(w http.ResponseWriter, r *http.Request) {

  b, err := ioutil.ReadAll(r.Body)
  defer r.Body.Close()

  var info Credentials

  InternalServerError(err, w)

  user := Credentials{}
  err = json.Unmarshal(b, &info)

  InternalServerError(err, w)

  Username, err := info.Username, nil
  Password, err := info.Password, nil

  InternalServerError(err, w)

  user.Username = string(Username)
  user.Password = string(Password)

  err = store.CreateUser(&user)

  if(err != nil) {
    fmt.Println(err)
  }
}
