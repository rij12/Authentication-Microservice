package models

type User struct {
	UserID   string `bson:"userid"   json:"userid"`
	Email    string `bson:"email"    json:"email"`
	Password string `bson:"password" json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message    string
	StatusCode int
}
