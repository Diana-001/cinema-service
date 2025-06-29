package app

import (
	"cinema-service/internal/factories"
	"cinema-service/internal/models"
	"cinema-service/internal/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type MyApp struct {
	Router  *gin.Engine
	DB      *gorm.DB
	Handler *factories.HandlersFactory
}

// NewApp Первоначальная инициализация приложения
func NewApp() (*MyApp, error) {
	app := &MyApp{}

	err := app.prepareConfigsAndComponents()
	if err != nil {
		log.Fatal("Ошибка при подготовки компонентов для проекта", err)
	}

	return app, nil
}

// Запуск сервера
func (app *MyApp) RunHttp() {
	app.Router = gin.Default()
	routes.SetRoutes(app.Router, app.Handler)
	app.Router.Run()
}

func (app *MyApp) prepareConfigsAndComponents() error {
	app.initDB()

	repoFactory := factories.NewSqlRepositoryFactory(app.DB)
	app.Handler = factories.NewHandlersFactory(repoFactory)

	return nil
}

func (app *MyApp) initDB() {
	db, err := gorm.Open(sqlite.Open("cinema.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка при подключений sqlite", err)
	}

	app.DB = db

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Movie{})
}
