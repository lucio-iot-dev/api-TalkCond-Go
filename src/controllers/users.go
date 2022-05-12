package controllers

import (
	"api-TalkCond/src/answers"
	"api-TalkCond/src/banco"
	"api-TalkCond/src/models"
	"api-TalkCond/src/repositories"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// CreateUser inserts a user into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		answers.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("cadastro"); erro != nil {
		answers.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		answers.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepositories(db)
	user.ID, erro = repository.CreateUser(user)
	if erro != nil {
		answers.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	answers.JSON(w, http.StatusCreated, user)
}

// SearchUsers search all users saved in the database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("usuarios"))
	db, erro := banco.Conectar()
	if erro != nil {
		answers.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepositories(db)
	users, erro := repository.Search(nameOrNick)
	if erro != nil {
		answers.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	answers.JSON(w, http.StatusOK, users)
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
