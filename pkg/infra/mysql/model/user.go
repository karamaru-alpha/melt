package model

import (
	"github.com/karamaru-alpha/melt/pkg/domain/entity"
)

type User struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func NewUser(user *entity.User) *User {
	return &User{
		ID:   user.ID,
		Name: user.Name,
	}
}

func (u *User) ToEntity() *entity.User {
	return &entity.User{
		ID:   u.ID,
		Name: u.Name,
	}
}
