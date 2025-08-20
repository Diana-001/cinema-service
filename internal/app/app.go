package app

import (
	"cinema-service/internal/factories"
	"cinema-service/internal/models"
	"cinema-service/internal/repositories"
	"cinema-service/internal/routes"
	"cinema-service/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var _ App = (*MyApp)(nil)

type (
	App interface {
		RunHttp()
	}

	MyApp struct {
		Router  *gin.Engine
		DB      repositories.Repository
		Logger  *logger.Logger
		Handler *factories.HandlersFactory
	}
)

// NewApp - Первоначальная инициализация приложения
func NewApp() (App, error) {
	app := &MyApp{}

	err := app.prepareConfigsAndComponents()
	if err != nil {
		log.Fatal("Ошибка при подготовки компонентов для проекта", err)
	}

	return app, nil
}

// RunHttp - Запуск сервера
func (app *MyApp) RunHttp() {
	app.Router = gin.Default()
	routes.SetRoutes(app.Router, app.Handler)
	app.Router.Run()
}

func (app *MyApp) prepareConfigsAndComponents() error {
	db := app.initDB()
	l := logger.NewLoggerImpl(&logger.LoggerCfg{Lvl: "debug"})

	repoFactory := factories.NewSqlRepositoryFactory(db, l)
	app.Handler = factories.NewHandlersFactory(repoFactory, l)

	return nil
}

func (app *MyApp) initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("cinema.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка при подключений sqlite", err)
	}

	if err = db.AutoMigrate(&models.User{}, &models.Movie{}, &models.Hall{}, &models.Session{}, &models.RefreshToken{}); err != nil {
		log.Fatal(err)
	}

	return db
}
