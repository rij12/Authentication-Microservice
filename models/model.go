package models

import "time"

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

type Config struct {
	Server struct {
		Host             string        `yaml:"host"`
		Port             string        `yaml:"port"`
		TimeoutInSeconds time.Duration `yaml:"timeout"`
	} `yaml:"server"`
	Database struct {
		Username string `envconfig:"MONGO_USERNAME"`
		Password string `envconfig:"MONGO_PASSWORD"`
		Host     string `yaml:"host", envconfig:"MONGO_HOST" env-default:"localhost"`
		Port     int    `yaml:"port", envconfig:"MONGO_PORT" env-default:"27017"`
	} `yaml:"database"`
	Crypto struct {
		JWTSecret string        `envconfig:"JWT_SECRET"`
		JWTIssuer string        `yaml:"issuer"`
		JWTExpire time.Duration `yaml:"tokenExpire"`
		SSLCert   string        `yaml:"SSLCert"`
		SSLKey    string        `yaml:"SSLKey"`
	} `yaml:"crypto"`
}
