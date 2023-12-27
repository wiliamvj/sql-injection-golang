package database

import (
  "database/sql"
  "errors"
  "os"

  "github.com/joho/godotenv"
)

var DBConnection *sql.DB

func NewDBConnection() error {
  err := godotenv.Load(".env")
  if err != nil {
    return errors.New("error loading .env file")
  }

  databaseURL := os.Getenv("DATABASE_URL")
  db, err := sql.Open("postgres", databaseURL)
  if err != nil {
    return err
  }
  DBConnection = db

  return nil
}
