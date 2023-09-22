package main

import (
    "database/sql"
    _ "github.com/lib/pq"
)

type Storage interface {
    CreateAccount(*Account) error
    DeteteAccount(int) error
    UpdateAccount(*Account) error
    GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
    db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
    connStr := "user=root dbname=go-bank password=password sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &PostgresStore{
        db:db,
    }, nil
}
