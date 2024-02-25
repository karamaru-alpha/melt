package model

import (
	domain "github.com/karamaru-alpha/melt/pkg/domain/entity"
)

type User struct {
	ID   string `gorm:"primary_key"`
	Name string
}

func NewUser(entity *domain.User) *User {
	if entity == nil {
		return nil
	}
	return &User{
		ID:   entity.ID,
		Name: entity.Name,
	}
}

func (u *User) ToEntity() *domain.User {
	return &domain.User{
		ID:   u.ID,
		Name: u.Name,
	}
}
