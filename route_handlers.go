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

const errNum int = 500

func InternalServerError(err error, w http.ResponseWriter) {
  if err != nil {
    http.Error(w, err.Error(), errNum)
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

func SignIn(w http.ResponseWriter, r *http.Request) {
  b, err := ioutil.ReadAll(r.Body)
  defer r.Body.Close()

  var userCheck Credentials

  InternalServerError(err, w)

  user := Credentials{}
  err = json.Unmarshal(b, &userCheck)

  InternalServerError(err, w)

  Username, err := userCheck.Username, nil
  Password, err := userCheck.Password, nil

  InternalServerError(err, w)

  user.Username = string(Username)
  user.Password = string(Password)

  err = store.SignInUser(&user, w)

  if err != nil {
    w.WriteHeader(http.StatusUnauthorized)
  }

}
