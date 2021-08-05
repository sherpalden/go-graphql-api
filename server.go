package main

import (
	"go-graphql-api/module"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	fx.New(module.Module).Run()
}
