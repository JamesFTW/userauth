package main

import (
  "net/http"
  "log"
  "database/sql"
  "fmt"
_"github.com/lib/pq"
  "golang.org/x/crypto/bcrypt"
)

type Store interface {
  CreateUser(user *Credentials) error
  SignInUser(user *Credentials, w http.ResponseWriter) error
}

type dbStore struct {
  db *sql.DB
}

const cost int = bcrypt.DefaultCost

func hashedPassword(password string) string {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

  if err != nil {
    fmt.Println("GenerateFromPassword error: %s", err)
  }

  return string(bytes)

}

func (store *dbStore) CreateUser(creds *Credentials) error {
  _, err := store.db.Query("INSERT INTO USERS(username, password) VALUES ($1,$2)", creds.Username, hashedPassword(creds.Password))

  if err != nil {
    fmt.Println(err)
  }

  return err
}

func (store *dbStore) SignInUser(creds *Credentials, w http.ResponseWriter) error {
  var username string

  results, err := store.db.Query("Select password FROM users where username=$1", creds.Username)

  if err != nil {
    log.Fatal(err)
  }

  err = store.db.QueryRow("Select password FROM users where username=$1", creds.Username).Scan(&username)

  if err != nil {
    log.Printf("Error Code: %d",http.StatusUnauthorized)
    return err
  }

  defer results.Close()

  for results.Next() {
    var password string

    err := results.Scan(&password)

    if err != nil {
      w.WriteHeader(http.StatusUnauthorized)
    }

    if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(creds.Password)); err != nil {
      log.Printf("Error Code: %d",http.StatusUnauthorized)
      return err
    }
  }
  return err
}

var store Store

func InitStore(s Store) {
  store = s
}
