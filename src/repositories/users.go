package repositories

import (
	"api-TalkCond/src/models"
	"database/sql"
	"os/user"
)

type Users struct {
	db *sql.DB
}

func  NewUsersRepositories(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) CreateUser(user models.User) (uint64, error) {
	  statement, erro := repository.db.Prepare(
			"insert into usuarios (nome, nick, email, senha, apartamento) values(?,?, ?, ?, ?)",
		)
		if erro != nil {
			return 0, erro
		}
		defer statement.Close()

		result, erro := statement.Exec(user.Nome, user.Nick, user.Email, user.Senha, user.Apartamento)
		if erro != nil {
			return 0, erro
		}
		lastIdInserted, erro := result.LastInsertId()
		if erro != nil {
			return 0, erro
		}
		return uint64(lastIdInserted), nil
}