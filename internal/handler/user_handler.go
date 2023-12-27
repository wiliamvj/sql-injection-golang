package handler

import (
  "encoding/json"
  "fmt"
  "net/http"

  "github.com/wiliamvj/sql-injection-golang/internal/service"
)

func GetUsersInjectHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
    http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
    return
  }
  id := r.URL.Query().Get("id")
  if id == "" {
    http.Error(w, "id não informado", http.StatusBadRequest)
    return
  }

  users, err := service.GetUserInject(id)
  if err != nil {
    fmt.Println(err)
    http.Error(w, "Erro ao buscar os usuários", http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(users); err != nil {
    http.Error(w, "Erro ao codificar os usuários para JSON", http.StatusInternalServerError)
    return
  }
}

func GetUsersCorrectHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
    http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
    return
  }
  id := r.URL.Query().Get("id")
  if id == "" {
    http.Error(w, "id não informado", http.StatusBadRequest)
    return
  }

  users, err := service.GetUserCorrect(id)
  if err != nil {
    fmt.Println(err)
    http.Error(w, "Erro ao buscar os usuários", http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(users); err != nil {
    http.Error(w, "Erro ao codificar os usuários para JSON", http.StatusInternalServerError)
    return
  }
}

func DeleteUserInjectHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodDelete {
    http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
    return
  }
  id := r.URL.Query().Get("id")
  if id == "" {
    http.Error(w, "id não informado", http.StatusBadRequest)
    return
  }

  err := service.DeleteUserInject(id)
  if err != nil {
    fmt.Println(err)
    http.Error(w, "Erro ao deletar os usuário", http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode("Usuário deletado com sucesso"); err != nil {
    http.Error(w, "Erro ao codificar os usuários para JSON", http.StatusInternalServerError)
    return
  }
}

func DeleteUserCorrectHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodDelete {
    http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
    return
  }
  id := r.URL.Query().Get("id")
  if id == "" {
    http.Error(w, "id não informado", http.StatusBadRequest)
    return
  }

  err := service.DeleteUserCorrect(id)
  if err != nil {
    fmt.Println(err)
    http.Error(w, "Erro ao deletar os usuário", http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode("Usuário deletado com sucesso"); err != nil {
    http.Error(w, "Erro ao codificar os usuários para JSON", http.StatusInternalServerError)
    return
  }
}
