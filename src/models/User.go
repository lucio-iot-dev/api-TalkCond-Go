package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a user using the social network
type User struct {
	ID          uint64    `json:"id,omitempty"`
	Nome        string    `json:"nome,omitempty"`
	Nick        string    `json:"nick,omitempty"`
	Email       string    `json:"email,omitempty"`
	Senha       string    `json:"senha,omitempty"`
	Apartamento uint64    `json:"apartamento,omitempty"`
	CriadoEm    time.Time `json:"CriadoEm,omitempty"`
}

func (user *User) Prepare() error {
	if erro := user.validate(); erro != nil {
		return erro
	}
	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if user.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}
	if user.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	// if stage == "cadastro" && user.Senha == "" {
	// 	return errors.New("A senha é obrigatória e não pode estar em branco")
	// }

	if user.Senha == "" {
		return errors.New("A senha é obrigatório e não pode estar em branco")
	}
	if user.Apartamento == 0 {
		return errors.New("O numero do apartamento não pode estar em branco")
	}
	return nil

}

func (user *User) format() error {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	return nil
}
