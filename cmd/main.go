package main

import (
	app2 "cinema-service/internal/app"
	"log"
)

// @title Cinema Service API
// @version 1.0
// @description REST API для кино-сервиса на Go + Gin + SQLite
// @host localhost:8080
// @BasePath /cinema-service
func main() {
	app, err := app2.NewApp()
	if err != nil {
		log.Fatal("Ошибка при инициализаций сервиса", err)
	}

	app.RunHttp()
}
