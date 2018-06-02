package main

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "fmt"
)

type Credentials struct {
  Username string `json: username`
  Password string `json: "password"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {

  b, err := ioutil.ReadAll(r.Body)
  defer r.Body.Close()

  var info Credentials

  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }

  user := Credentials{}
  err = json.Unmarshal(b, &info)

  if err != nil {
    http.Error(w, err.Error(), 500)
    return
	}

  Username, err := info.Username, nil
  Password, err := info.Password, nil

  if err != nil {
    http.Error(w, err.Error(), 500)
    return
	}

  user.Username = string(Username)
  user.Password = string(Password)

  err = store.CreateUser(&user)

  if(err != nil) {
    fmt.Println(err)
  }
}
