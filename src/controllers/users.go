package controllers

import (
	"api-TalkCond/src/banco"
	"api-TalkCond/src/models"
	"api-TalkCond/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateUser inserts a user into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
   bodyRequest, erro := ioutil.ReadAll(r.Body) {
     if erro != nil {
       log.Fatal(erro)
     }

     var user models.User
     if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
       log.Fatal(erro)
     }

     db, erro := banco.Conectar()
     if erro != nil {
       log.Fatal(erro)
     }

     repository := repositories.NewUsersRepositories(db)
     userID, erro := repository.CreateUser(user)
     if erro != nil {
       log.Fatal(erro)
     }
     w.Write([]byte(fmt.Sprintf("Id inserido: %d", userID)))
   }
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
func UpdateUser(w http.ResponseWriter, r *http.Requestgit status) {
  w.Write([]byte("Atualizando um usuário!"))
}

// DeleteUser deletes a user's information in the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Deletando um usuário!"))
}