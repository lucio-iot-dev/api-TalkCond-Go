package repositories

import (
	"api-TalkCond/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewUsersRepositories(db *sql.DB) *Users {
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

func (repository Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, erro := repository.db.Query(
		"select id, nome, nick, email, apartamento, criadaEm from usuarios where nome LIKE ?",
		nameOrNick, nameOrNick,
	)
	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var users[]models.User

	for lines.Next() {
		var user models.User
		
		if erro = lines.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			user.Email,
      user.Apartamento,
			user.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		users = append(users,user)

	}
	return users, nil
}
