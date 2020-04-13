package models

type database interface {
	delete(user User) User
	update(user User) User
	create(user User) User
	getUserByEmail(email string) User
	getUserbyID(id string) User
}
