package service

import (
  "fmt"

  "github.com/wiliamvj/sql-injection-golang/internal/database"
)

type User struct {
  ID    int    `json:"id"`
  Name  string `json:"name"`
  Email string `json:"email"`
}

func GetUserInject(id string) ([]User, error) {
  query := fmt.Sprintf("SELECT id, name, email FROM users WHERE id = %s", id)
  rows, err := database.DBConnection.Query(query)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var users []User
  for rows.Next() {
    var user User
    if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
      return nil, err
    }
    users = append(users, user)
  }

  return users, nil
}

func GetUserCorrect(id string) ([]User, error) {
  query := "SELECT id, name, email FROM users WHERE id = $1"
  rows, err := database.DBConnection.Query(query, id)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var users []User
  for rows.Next() {
    var user User
    if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
      return nil, err
    }
    users = append(users, user)
  }

  return users, nil
}

func DeleteUserInject(id string) error {
  query := fmt.Sprintf("DELETE FROM users WHERE id = %s", id)
  _, err := database.DBConnection.Exec(query)
  if err != nil {
    return err
  }

  return nil
}

func DeleteUserCorrect(id string) error {
  query := "DELETE FROM users WHERE id = $1"
  _, err := database.DBConnection.Exec(query, id)
  if err != nil {
    return err
  }

  return nil
}
