package main

import (
	"os"

	"github.com/ljsea6/go-hello-world/app/router"
	"github.com/ljsea6/go-hello-world/config"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	app.Run(":" + port)
}
