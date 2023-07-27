package data

import (
	"simple-password-manager/domain"

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

func (pr *PasswordRepository) Create(password domain.Password) error {
	pr.passwords = append(pr.passwords, password)
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
