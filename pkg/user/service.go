package user

import (
	"archis/pkg/entities"
)

//Service is the outside module facing interface providing various implementations
type Service interface {
	CreateUser(AuthRequest) error
	UpdateUser(UpdateRequest) (entities.User, error)
	GetUser(id string) (entities.User, error)
	DeleteUser(id string) error
}

type userSvc struct {
	repo Repository
}

//NewUserService returns a new user service
func NewUserService(r Repository) Service {
	return &userSvc{repo: r}
}

func (uS *userSvc) CreateUser(req AuthRequest) error {
	return uS.repo.CreateUser(req)
}

func (uS *userSvc) UpdateUser(req UpdateRequest) (entities.User, error) {
	return uS.repo.UpdateUser(req)
}

func (uS *userSvc) GetUser(id string) (entities.User, error) {
	return uS.repo.GetUser(id)
}

func (uS *userSvc) DeleteUser(id string) error {
	return uS.repo.DeleteUser(id)
}
