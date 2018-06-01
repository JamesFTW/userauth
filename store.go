package main

import (
  "database/sql"
  "fmt"
_"github.com/lib/pq"
  "golang.org/x/crypto/bcrypt"
)

type Store interface {
  CreateUser(user *Credentials) error
}

type dbStore struct {
  db *sql.DB
}

func hashedPassword(password string) string {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)

  if err != nil {
    fmt.Println("GenerateFromPassword error: %s", err)
  }

  return string(bytes)

}

func (store *dbStore) CreateUser(creds *Credentials) error {
  _, err := store.db.Query("INSERT INTRO USERS(username, password) VALUES ($1,$2)", creds.Username, hashedPassword(creds.Password))

  if err != nil {
    fmt.Println(err)
  }
  return err
}

var store Store

func InitStore(s Store) {
  store = s
}
