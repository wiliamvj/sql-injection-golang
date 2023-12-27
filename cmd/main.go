package main

import (
  "net/http"

  "github.com/go-chi/chi"
  _ "github.com/lib/pq"
  "github.com/wiliamvj/sql-injection-golang/internal/database"
  "github.com/wiliamvj/sql-injection-golang/internal/handler"
)

func main() {
  err := database.NewDBConnection()
  if err != nil {
    panic(err)
  }

  database.SeedUsers()
  r := chi.NewRouter()
  r.Get("/users", handler.GetUsersInjectHandler)
  r.Get("/users/correct", handler.GetUsersCorrectHandler)
  r.Delete("/users", handler.DeleteUserInjectHandler)
  r.Delete("/users/correct", handler.DeleteUserCorrectHandler)

  server := &http.Server{
    Addr:    ":8080",
    Handler: r,
  }
  server.ListenAndServe()
}
