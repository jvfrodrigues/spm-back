package domain

import "github.com/google/uuid"

type Password struct {
	Id       string `json:"id"`
	Url      string `json:"url"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewPassword(url string, name string, username string, password string) *Password {
	return &Password{
		Id:       uuid.New().String(),
		Url:      url,
		Name:     name,
		Username: username,
		Password: password,
	}
}
