package infrastructure

import "os"

type Env struct {
	Environment       string
	ServerPort        string
	DBUser            string
	DBPassword        string
	DBHost            string
	DBPort            string
	DBName            string
	JWTSecret         string
	ClientID          string
	ClientSecret      string
	GoogleRedirectURI string
	SentryDSN         string
}

func NewEnv() Env {
	env := Env{}
	env.LoadEnv()
	return env
}

func (env *Env) LoadEnv() {
	env.Environment = os.Getenv("ENVIRONMENT")
	env.ServerPort = os.Getenv("SERVER_PORT")
	env.DBUser = os.Getenv("DB_USER")
	env.DBPassword = os.Getenv("DB_PASSWORD")
	env.DBHost = os.Getenv("DB_HOST")
	env.DBPort = os.Getenv("DB_PORT")
	env.DBName = os.Getenv("DB_NAME")
	env.JWTSecret = os.Getenv("JWT_SECRET")
	env.ClientID = os.Getenv("CLIENT_ID")
	env.ClientSecret = os.Getenv("CLIENT_SECRET")
	env.GoogleRedirectURI = os.Getenv("GOOGLE_REDIRECT_URI")
	env.SentryDSN = os.Getenv("SENTRY_DSN")
}
