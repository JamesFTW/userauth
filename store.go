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

var store Store

func InitStore(s Store) {
  store = s
}
