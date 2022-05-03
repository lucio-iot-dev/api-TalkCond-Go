package controllers

import "net/http"

// CreateUser inserts a user into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Usuário criado com sucesso!"))
}

// SearchUsers search all users saved in the database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Buscando todos os usuários!"))
}

// SearchUser searches for a user saved in the database
func SearchUser(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Buscando um usuário!"))
}

// Update User changes a user's information in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Atualizando um usuário!"))
}

// DeleteUser deletes a user's information in the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Deletando um usuário!"))
}