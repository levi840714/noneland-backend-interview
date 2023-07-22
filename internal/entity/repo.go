package entity

type Repository interface {
	// User
	GetUsers() (users []User, err error)
}
