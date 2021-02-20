package user

import (
	"archis/pkg/entities"
)

//Repository represents the user repository
type Repository interface {
	CreateUser(AuthRequest) error
	UpdateUser(UpdateRequest) (entities.User, error)
	GetUser(id string) (entities.User, error)
	DeleteUser(id string) error
}
