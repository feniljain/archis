package user

import (
	"archis/pkg/entities"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/twinj/uuid"
)

type repo struct {
	db *sqlx.DB
}

//AuthRequest depicts the login request structure
type AuthRequest struct {
	Email string `json:"email"`
	//Password string `json:"password"`
	Amount int `json:"amount"`
}

//UpdateRequest depicts the login request structure
type UpdateRequest struct {
	ID     string  `json:"_id"`
	Email  *string `json:"email,omitempty"`
	Amount *int    `json:"amount,omitempty"`
}

//NewPostgresRepo returns a new user repository
func NewPostgresRepo(db *sqlx.DB) Repository {
	return &repo{db: db}
}

//CreateUser creates a user and return corresponding result
func (r *repo) CreateUser(req AuthRequest) error {
	userID := uuid.NewV4()

	q := fmt.Sprintf("insert into users values('%v', '%v', '%v');", userID, req.Email, req.Amount)

	_, err := r.db.Exec(q)
	if err != nil {
		return err
	}

	return nil
}

//UpdateUser creates a user and return corresponding result
func (r *repo) UpdateUser(req UpdateRequest) (entities.User, error) {

	_, err := r.db.NamedExec(`update users set email = :email, amount = :amount where _id = :id`, &req)
	if err != nil {
		return entities.User{}, err
	}

	return entities.User{
		Email:  *req.Email,
		Amount: *req.Amount,
		ID:     req.ID,
	}, nil
}

//GetUser gets a user by id
func (r *repo) GetUser(id string) (entities.User, error) {
	user := entities.User{}
	err := r.db.Get(&user, "select * from users where _id=$1", id)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

//DeleteUser deletes a user by id
func (r *repo) DeleteUser(id string) error {
	_, err := r.db.NamedExec(`delete from users where _id=:id`, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return err
	}

	return nil
}
