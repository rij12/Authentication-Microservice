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

type Config struct {
	Server struct {
		Host string `yaml:"host", envconfig:"SERVER_HOST" env-default:"localhost"`
		Port string `yaml:"port", envconfig:"SERVER_PORT" env-default:"8000"`
	} `yaml:"server"`
	Database struct {
		Username string `envconfig:"MONGO_USERNAME"`
		Password string `envconfig:"MONGO_PASSWORD"`
		Host     string `yaml:"host", envconfig:"MONGO_HOST" env-default:"localhost"`
		Port     int    `yaml:"port", envconfig:"MONGO_PORT" env-default:"27017"`
	} `yaml:"database"`
	Crypto struct {
		JWTSecret string `envconfig:"JWT_SECRET"`
	}
}
