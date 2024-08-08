package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nome     string `json:"nome"`
	Senha    string `json:"senha"`
	Endereco string `json:"endereco"`
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"isAdmin"`
}
