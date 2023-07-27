package data

import (
	"github.com/jvfrodrigues/simple-password-manager-back/domain"
	dtos "github.com/jvfrodrigues/simple-password-manager-back/domain/dtos"
	"golang.org/x/exp/slices"
)

type PasswordRepository struct {
	passwords []domain.Password
}

func NewPasswordRepository() *PasswordRepository {
	return &PasswordRepository{
		passwords: make([]domain.Password, 0),
	}
}

func (pr *PasswordRepository) FindAll() ([]domain.Password, error) {
	return pr.passwords, nil
}

func (pr *PasswordRepository) Create(password dtos.CreatePasswordRequest) error {
	newPassword := domain.NewPassword(password.Url, password.Name, password.Username, password.Password)
	pr.passwords = append(pr.passwords, *newPassword)
	return nil
}

func (pr *PasswordRepository) Update(password domain.Password) error {
	idx := slices.IndexFunc(pr.passwords, func(p domain.Password) bool { return p.Id == password.Id })
	pr.passwords[idx] = password
	return nil
}

func (pr *PasswordRepository) Delete(id string) error {
	updatedPasswords := slices.DeleteFunc(pr.passwords, func(p domain.Password) bool { return p.Id == id })
	pr.passwords = updatedPasswords
	return nil
}
