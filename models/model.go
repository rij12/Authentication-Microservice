package models

type User struct {
	UserID   string `bson:"userId"   json:"userId"`
	Email    string `bson:"email"    json:"email"`
	Password string `bson:"password" json:"password"`
}

type UserResult struct {
	UserID string `bson:"userId"   json:"userId"`
	Email  string `bson:"email"    json:"email"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
