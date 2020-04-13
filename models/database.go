package models

type Database interface {
	delete(user User) User
	update(user User) User
	create(user User) User
	getUserByEmail(email string) User
	getUserbyID(id string) User
	ConnectDB(username string, password string, url string)
	PingDb() error
}
