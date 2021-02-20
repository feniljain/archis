package entities

//User struct represents user db structure
type User struct {
	ID     string `json:"_id" db:"_id"`
	Email  string `json:"email" db:"email"`
	Amount int    `json:"amount" db:"amount"`
}
