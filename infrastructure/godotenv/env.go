package godotenv

import (
	"github.com/joho/godotenv"
	"os"
)

type Env struct {
	DATABASE_HOST string
	ServerAddr    string
}

func NewEnv() *Env {
	return &Env{}
}

func (e *Env) Load() {
	godotenv.Load(".env") // using .env file is not mandatory
	e.DATABASE_HOST = os.Getenv("DATABASE_HOST")
	e.ServerAddr = os.Getenv("ServerAddr")
}
