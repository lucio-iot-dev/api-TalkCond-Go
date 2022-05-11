package models

import (
	"errors"
	"strings"
	"time"
)

// User represents a user using the social network
type User struct {
	ID uint64 `json:"id,omitempty"`
	Nome string `json:"nome,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
	Apartamento uint64 `json:"apartamento,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (User *User) validate() error {
	if User.Nome == ""{
    return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if User.Nick == ""{
    return errors.New("O nick é obrigatório e não pode estar em branco")
	}
	if User.Email == ""{
    return errors.New("O email é obrigatório e não pode estar em branco")
	}
	if User.Senha == ""{
    return errors.New("A senha é obrigatório e não pode estar em branco")
	}
	if User.Apartamento == 0  {
    return errors.New("O numero do apartamento não pode estar em branco")
	}
	return nil

}

func (User *User) format() error{
	User.Nome = strings.TrimSpace(User.Nome)
	User.Nick = strings.TrimSpace(User.Nick)
	User.Email = strings.TrimSpace(User.Email)
}
